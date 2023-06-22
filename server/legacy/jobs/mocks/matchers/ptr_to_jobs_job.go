// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	jobs "github.com/runatlantis/atlantis/server/legacy/jobs"
)

func AnyPtrToJobsJob() *jobs.Job {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*jobs.Job))(nil)).Elem()))
	var nullValue *jobs.Job
	return nullValue
}

func EqPtrToJobsJob(value *jobs.Job) *jobs.Job {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *jobs.Job
	return nullValue
}

func NotEqPtrToJobsJob(value *jobs.Job) *jobs.Job {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue *jobs.Job
	return nullValue
}

func PtrToJobsJobThat(matcher pegomock.ArgumentMatcher) *jobs.Job {
	pegomock.RegisterMatcher(matcher)
	var nullValue *jobs.Job
	return nullValue
}