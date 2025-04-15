package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"sync"
)

// Setup the logger to simple STDOUT logging for now
func init() {
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime)
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	criticalLogger = log.New(os.Stdout, "CRITICAL: ", log.Ldate|log.Ltime)
	fatalLogger = log.New(os.Stdout, "FATAL: ", log.Ldate|log.Ltime)
}

// Logger instances
var (
	// Logging directories and files
	LOG_DIR  string
	LOG_FILE string
	// Logging levels
	LOG_LEVEL  int = 1
	LOG_LEVELS     = map[string]int{
		"DEBUG":    0,
		"INFO":     1,
		"ERROR":    2,
		"CRITICAL": 3,
		"FATAL":    4,
	}
	// Mutexes
	mu_LOG_DIR   = &sync.RWMutex{}
	mu_LOG_FILE  = &sync.RWMutex{}
	mu_LOG_LEVEL = &sync.RWMutex{}
	// Loggers
	debugLogger    *log.Logger
	infoLogger     *log.Logger
	errorLogger    *log.Logger
	criticalLogger *log.Logger
	fatalLogger    *log.Logger
)

type LoggerSetupArgs struct {
	// LogDir is the directory where the log file will be created.
	LogDir string
	// LogName is the name of the log file.
	LogName string
	// LogLevel is the logging level. It can be DEBUG, INFO, ERROR, CRITICAL, or FATAL.
	LogLevel string
}

func SetupLogger(args LoggerSetupArgs) {
	// Lock the mutexes
	mu_LOG_DIR.Lock()
	mu_LOG_FILE.Lock()
	defer mu_LOG_DIR.Unlock()
	defer mu_LOG_FILE.Unlock()
	// Set the logging directories and files
	LOG_DIR = args.LogDir
	LOG_FILE = path.Join(args.LogDir, args.LogName)

	// Set the logging level
	mu_LOG_LEVEL.Lock()
	defer mu_LOG_LEVEL.Unlock()
	if _, ok := LOG_LEVELS[strings.ToUpper(args.LogLevel)]; !ok {
		fmt.Println("Invalid log level:", args.LogLevel, "Defaulting to INFO")
		LOG_LEVEL = LOG_LEVELS["INFO"]
	} else {
		LOG_LEVEL = LOG_LEVELS[strings.ToUpper(args.LogLevel)]
	}

	// Check if the directory for the log directory exists
	err := os.MkdirAll(LOG_DIR, os.ModePerm)
	if err != nil {
		fmt.Println("Failed to create log directory at", LOG_DIR, ":", err)
		os.Exit(1)
	}

	file, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		fmt.Println("Fatal error: Failed to open log file:", err)
		os.Exit(1)
	}

	multi := io.MultiWriter(file, os.Stdout)

	// Initialize loggers
	debugLogger = log.New(multi, "DEBUG: ", log.Ldate|log.Ltime)
	infoLogger = log.New(multi, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(multi, "ERROR: ", log.Ldate|log.Ltime)
	criticalLogger = log.New(multi, "CRITICAL: ", log.Ldate|log.Ltime)
	fatalLogger = log.New(multi, "FATAL: ", log.Ldate|log.Ltime)
}

// Debug logs a debug message.
func Debug(v ...interface{}) {
	if LOG_LEVEL <= 0 {
		debugLogger.Println(v...)
	}
}

// Info logs an info message.
func Info(v ...interface{}) {
	if LOG_LEVEL <= 1 {
		infoLogger.Println(v...)
	}
}

// Error logs an error message.
func Error(v ...interface{}) {
	if LOG_LEVEL <= 2 {
		errorLogger.Println(v...)
	}
}

// Critical logs a critical message but does not exit the application.
func Critical(v ...interface{}) {
	if LOG_LEVEL <= 3 {
		criticalLogger.Println(v...)
	}
}

// Fatal logs a fatal message and exits the application.
func Fatal(v ...interface{}) {
	if LOG_LEVEL <= 4 {
		fatalLogger.Println(v...)
	}
	os.Exit(1)
}
