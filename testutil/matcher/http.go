package matcher

import (
	"errors"
	"fmt"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"net/http"
	"reflect"
)

func Name() string {
	return "matcher"
}

func ShouldHttpOK(actual interface{}, expected ...interface{}) string {
	code, ok := actual.(int)
	if !ok {
		return fmt.Sprintf("actual not RespondRecorder type %v", reflect.TypeOf(actual))
	}

	expectedCode := http.StatusOK
	if code == expectedCode {
		return ""
	}

	return fmt.Sprintf("http code: %d not %d", code, expectedCode)
}

func ShouldHttpCreated(actual interface{}, expected ...interface{}) string {
	code, ok := actual.(int)
	if !ok {
		return fmt.Sprintf("actual not RespondRecorder type %v", reflect.TypeOf(actual))
	}

	expectedCode := http.StatusCreated
	if code == expectedCode {
		return ""
	}

	return fmt.Sprintf("http code: %d not %d", code, expectedCode)
}

func ShouldHttpNoContent(actual interface{}, expected ...interface{}) string {
	code, ok := actual.(int)
	if !ok {
		return fmt.Sprintf("actual not RespondRecorder type %v", reflect.TypeOf(actual))
	}

	expectedCode := http.StatusNoContent
	if code == expectedCode {
		return ""
	}

	return fmt.Sprintf("http code: %d not %d", code, expectedCode)
}

func ShouldHttpBadRequest(actual interface{}, expected ...interface{}) string {
	code, ok := actual.(int)
	if !ok {
		return fmt.Sprintf("actual not RespondRecorder type %v", reflect.TypeOf(actual))
	}

	expectedCode := http.StatusBadRequest
	if code == expectedCode {
		return ""
	}

	return fmt.Sprintf("http code: %d not %d", code, expectedCode)
}

func ShouldHttpNotFound(actual interface{}, expected ...interface{}) string {
	code, ok := actual.(int)
	if !ok {
		return fmt.Sprintf("actual not RespondRecorder type %v", reflect.TypeOf(actual))
	}

	expectedCode := http.StatusNotFound
	if code == expectedCode {
		return ""
	}

	return fmt.Sprintf("http code: %d not %d", code, expectedCode)
}

func ShouldHttpInternalServerError(actual interface{}, expected ...interface{}) string {
	code, ok := actual.(int)
	if !ok {
		return fmt.Sprintf("actual not RespondRecorder type %v", reflect.TypeOf(actual))
	}

	expectedCode := http.StatusInternalServerError
	if code == expectedCode {
		return ""
	}

	return fmt.Sprintf("http code: %d not %d", code, expectedCode)
}

type beHttpCode struct {
	code int
}

func BeHttpStatusOK() gomega.OmegaMatcher {
	return &beHttpCode{code: http.StatusOK}
}

func BeHttpCreated() gomega.OmegaMatcher {
	return &beHttpCode{code: http.StatusCreated}
}
func BeHttpNoContent() gomega.OmegaMatcher {
	return &beHttpCode{code: http.StatusNoContent}
}
func BeHttpBadRequest() gomega.OmegaMatcher {
	return &beHttpCode{code: http.StatusBadRequest}
}
func BeHttpNotFound() gomega.OmegaMatcher {
	return &beHttpCode{code: http.StatusNotFound}
}
func BeHttpInternalServerError() gomega.OmegaMatcher {
	return &beHttpCode{code: http.StatusInternalServerError}
}

func (b *beHttpCode) Match(actual interface{}) (success bool, err error) {
	if code, ok := actual.(int); ok {
		return b.code == code, nil
	}
	err = errors.New(format.Message(actual, "HTTP Code", b.code))
	return
}

func (b *beHttpCode) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "bo be http code", b.code)
}

func (b *beHttpCode) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "not bo be http code", b.code)
}
