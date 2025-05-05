package error_handler

import (
	"fmt"
)

// HandleLogError is a function that handles errors and logs them with a message.
// It takes an error and a message as parameters.
func LogError(err error, message string) {
    if err != nil {
        fmt.Printf("Error: %s %v\n", message, err)
    }
}

var errorMessages = []map[string]string{
    {"error": "ErrNotFound", "message": "The requested resource was not found."},
    {"error": "ErrUnauthorized", "message": "You are not authorized to access this resource."},
    {"error": "ErrInvalidInput", "message": "The input provided is invalid."},
    {"error": "ErrInternalServer", "message": "An internal server error occurred."},
    {"error": "ErrTimeout", "message": "The request timed out."},
}

func ApiErrorResponse(err error, message string) map[string]interface{} {
    if err != nil {
        return map[string]interface{}{
            "error":   true,
            "message": fmt.Sprintf("API Error: %s %v", message, err),
        }
    }
    return map[string]interface{}{
        "error":   false,
        "message": message,
    }
}
