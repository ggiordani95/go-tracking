package main

import (
	"fmt"

	error_handler "github.com/ggiordani95/iot-tracking/utils/handlers/error_handler"
)

func main() {
	fmt.Println("Hello, IoT Tracking!")
	  errHandler := error_handler.ErrorHandler{
        Condition: "Some condition",
        Error:     fmt.Errorf("An example error"),
    }

    errHandler.HandleError(errHandler)
}