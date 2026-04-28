package ports

// Logger is an interface that defines the logging methods used in the application
type Logger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	IPrintf(level int, format string, v ...interface{})
}
