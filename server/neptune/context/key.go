package context

import "context"

type Key string

func (c Key) String() string {
	return string(c)
}

const (
	InstallationIDKey = Key("gh-installation-id")
	RequestIDKey      = Key("gh-request-id")
	RepositoryKey     = Key("repository")
	SHAKey            = Key("sha")
	PullNumKey        = Key("pull-num")
	ProjectKey        = Key("project")
	Err               = Key("err")
	DeploymentIDKey   = Key("deployment-id")
)

var Keys = []Key{RequestIDKey, RepositoryKey, PullNumKey, ProjectKey, SHAKey, InstallationIDKey, DeploymentIDKey}

type KVStore interface {
	Value(key interface{}) interface{}
}

// Extracts relevant fields from context for structured logging.
func ExtractFields(ctx KVStore) map[string]interface{} {
	args := make(map[string]interface{})

	for _, k := range Keys {
		if v, ok := ctx.Value(k).(string); ok {
			args[k.String()] = v
		}
	}

	return args
}

func ExtractFieldsAsList(ctx KVStore) []interface{} {
	var args []interface{}

	for _, k := range Keys {
		if v, ok := ctx.Value(k).(string); ok {
			args = append(args, k)
			args = append(args, v)
		}
	}

	return args
}

// Copies fields from a context to a new context created from a given base.
func CopyFields(base context.Context, from context.Context) context.Context {
	for _, k := range Keys {
		base = context.WithValue(base, k, from.Value(k))
	}
	return base
}