package error_handler

import (
	"fmt"
)

type ErrorHandler struct {
    // Error is the error to be handled.
    Condition any
    Error     error
}

func (e *ErrorHandler) HandleError(err ErrorHandler) {
    // Log the error
    fmt.Println("Error:", e.Error)
}