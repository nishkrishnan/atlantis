// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"

	http "net/http"
)

func AnyHTTPResponseWriter() http.ResponseWriter {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(http.ResponseWriter))(nil)).Elem()))
	var nullValue http.ResponseWriter
	return nullValue
}

func EqHTTPResponseWriter(value http.ResponseWriter) http.ResponseWriter {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue http.ResponseWriter
	return nullValue
}

func NotEqHTTPResponseWriter(value http.ResponseWriter) http.ResponseWriter {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue http.ResponseWriter
	return nullValue
}

func HTTPResponseWriterThat(matcher pegomock.ArgumentMatcher) http.ResponseWriter {
	pegomock.RegisterMatcher(matcher)
	var nullValue http.ResponseWriter
	return nullValue
}
