package api

// ApiErrorResponse is a function that generates an API error response.
// It takes an error and a message as parameters.
type ErrorCode string

const (
    ErrNotFound        ErrorCode = "ErrNotFound"
    ErrUnauthorized    ErrorCode = "ErrUnauthorized"
    ErrInvalidInput    ErrorCode = "ErrInvalidInput"
    ErrInternalServer  ErrorCode = "ErrInternalServer"
    ErrTimeout         ErrorCode = "ErrTimeout"
)

var errorMessages = map[ErrorCode]string{
    ErrNotFound:       "The requested resource was not found.",
    ErrUnauthorized:   "You are not authorized to access this resource.",
    ErrInvalidInput:   "The input provided is invalid.",
    ErrInternalServer: "An internal server error occurred.",
    ErrTimeout:        "The request timed out.",
}

func ApiErrorResponseWithCode(code ErrorCode) map[string]interface{} {
    message, exists := errorMessages[code]
    if !exists {
        return map[string]interface{}{
            "error":   true,
            "message": "An unknown error occurred.",
        }
    }
    return map[string]interface{}{
        "error":   true,
        "message": message,
    }
}


