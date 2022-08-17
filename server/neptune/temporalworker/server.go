package temporalworker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.temporal.io/sdk/client"

	"github.com/gorilla/mux"
	"github.com/palantir/go-githubapp/githubapp"
	"github.com/pkg/errors"
	"github.com/runatlantis/atlantis/server/controllers"
	"github.com/runatlantis/atlantis/server/controllers/templates"
	"github.com/runatlantis/atlantis/server/core/config/valid"
	"github.com/runatlantis/atlantis/server/logging"
	neptune_http "github.com/runatlantis/atlantis/server/neptune/http"
	"github.com/runatlantis/atlantis/server/neptune/temporal"
	"github.com/runatlantis/atlantis/server/neptune/workflows"
	"github.com/uber-go/tally/v4"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
	"go.temporal.io/sdk/worker"
)

const (
	AtlantisNamespace        = "atlantis"
	DeployTaskqueue          = "deploy"
	ProjectJobsViewRouteName = "project-jobs-detail"
)

// Config is TemporalWorker specific user config
type Config struct {
	AtlantisURL     *url.URL
	AtlantisVersion string
	CtxLogger       logging.Logger
	SslCertFile     string
	SslKeyFile      string
	Scope           tally.Scope
	Closer          io.Closer
	TemporalCfg     valid.Temporal
	Port            int
	App             githubapp.Config
}

type Server struct {
	Logger          logging.Logger
	HttpServerProxy *neptune_http.ServerProxy
	Port            int
	StatsScope      tally.Scope
	StatsCloser     io.Closer
	TemporalClient  client.Client
	Activities      *workflows.Activities
}

func NewServer(config *Config) (*Server, error) {
	jobsController := &controllers.JobsController{
		AtlantisVersion:     config.AtlantisVersion,
		AtlantisURL:         config.AtlantisURL,
		KeyGenerator:        controllers.JobIDKeyGenerator{},
		StatsScope:          config.Scope,
		Logger:              config.CtxLogger,
		ProjectJobsTemplate: templates.ProjectJobsTemplate,
	}

	// temporal client + worker initialization
	temporalClient, err := temporal.NewClient(config.Scope, config.CtxLogger, config.TemporalCfg)
	if err != nil {
		return nil, errors.Wrap(err, "initializing temporal client")
	}
	// router initialization
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	router.HandleFunc("/jobs/{job-id}", jobsController.GetProjectJobs).Methods("GET").Name(ProjectJobsViewRouteName)
	router.HandleFunc("/jobs/{job-id}/ws", jobsController.GetProjectJobsWS).Methods("GET")
	n := negroni.New(&negroni.Recovery{
		Logger:     log.New(os.Stdout, "", log.LstdFlags),
		PrintStack: false,
		StackAll:   false,
		StackSize:  1024 * 8,
	})
	n.UseHandler(router)
	httpServerProxy := &neptune_http.ServerProxy{
		SSLCertFile: config.SslCertFile,
		SSLKeyFile:  config.SslKeyFile,
		Server:      &http.Server{Addr: fmt.Sprintf(":%d", config.Port), Handler: n},
		Logger:      config.CtxLogger,
	}

	activities, err := workflows.NewActivities(config.App, config.Scope.SubScope("app"))

	if err != nil {
		return nil, errors.Wrap(err, "initializing activities")
	}
	server := Server{
		Logger:          config.CtxLogger,
		HttpServerProxy: httpServerProxy,
		Port:            config.Port,
		StatsScope:      config.Scope,
		StatsCloser:     config.Closer,
		TemporalClient:  temporalClient,
		Activities:      activities,
	}
	return &server, nil
}

func (s Server) Start() error {
	var wg sync.WaitGroup
	wg.Add(1)
	defer s.Logger.Close()
	defer s.TemporalClient.Close()

	go func() {
		defer wg.Done()
		w := worker.New(s.TemporalClient, workflows.DeployTaskQueue, worker.Options{
			// ensures that sessions are preserved on the same worker
			EnableSessionWorker: true,
		})

		w.RegisterActivity(s.Activities)
		w.RegisterWorkflow(workflows.Deploy)

		if err := w.Run(worker.InterruptCh()); err != nil {
			log.Fatalln("unable to start worker", err)
		}
	}()

	// Ensure server gracefully drains connections when stopped.
	stop := make(chan os.Signal, 1)
	// Stop on SIGINTs and SIGTERMs.
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	s.Logger.Info(fmt.Sprintf("Atlantis started - listening on port %v", s.Port))

	go func() {
		err := s.HttpServerProxy.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			s.Logger.Error(err.Error())
		}
	}()

	<-stop

	// flush stats before shutdown
	if err := s.StatsCloser.Close(); err != nil {
		s.Logger.Error(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.HttpServerProxy.Shutdown(ctx); err != nil {
		return cli.NewExitError(fmt.Sprintf("while shutting down: %s", err), 1)
	}
	wg.Wait()
	return nil
}

// Healthz returns the health check response. It always returns a 200 currently.
func Healthz(w http.ResponseWriter, _ *http.Request) {
	data, err := json.MarshalIndent(&struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating status json response: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data) // nolint: errcheck
}
