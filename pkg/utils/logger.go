package utils

import (
    "log"
    "os"
)

var (
    infoLogger  *log.Logger
    errorLogger *log.Logger
    debugLogger *log.Logger
)

// Initialize the loggers with different log levels.
func init() {
    // Define the format for each log type
    infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs general information.
func Info(msg string, v ...interface{}) {
    infoLogger.Printf(msg, v...)
}

// Error logs error messages.
func Error(msg string, v ...interface{}) {
    errorLogger.Printf(msg, v...)
}

// Debug logs detailed information for debugging purposes.
func Debug(msg string, v ...interface{}) {
    debugLogger.Printf(msg, v...)
}

// SetLogLevel allows users to change log levels (e.g., enabling or disabling debug logs).
func SetLogLevel(level string) {
    switch level {
    case "info":
        debugLogger.SetOutput(os.Stdout)  // Enable info and higher levels
    case "error":
        debugLogger.SetOutput(io.Discard) // Disable debug logs
    case "debug":
        debugLogger.SetOutput(os.Stdout)  // Enable all logs
    default:
        debugLogger.SetOutput(os.Stdout)  // Default to all logs enabled
    }
}
