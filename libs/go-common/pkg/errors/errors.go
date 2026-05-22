package errors

import (
	"fmt"
)

// TMFError represents a TM Forum compliant error structure
type TMFError struct {
	Code    string \`json:"code"\`
	Reason  string \`json:"reason"\`
	Message string \`json:"message,omitempty"\`
	Status  int    \`json:"-"\`
}

func (e *TMFError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Reason)
}

func NewBadRequestError(reason string) *TMFError {
	return &TMFError{Code: "400", Reason: reason, Status: 400}
}

func NewNotFoundError(reason string) *TMFError {
	return &TMFError{Code: "404", Reason: reason, Status: 404}
}

func NewInternalError(reason string) *TMFError {
	return &TMFError{Code: "500", Reason: reason, Status: 500}
}
