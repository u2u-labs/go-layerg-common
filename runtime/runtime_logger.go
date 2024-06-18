package runtime

/*
Logger exposes a logging framework to use in modules. It exposes level-specific logging functions and a set of common functions for compatibility.
*/
type Logger interface {
	/*
		Log a message with optional arguments at DEBUG level. Arguments are handled in the manner of fmt.Printf.
	*/
	Debug(format string, v ...interface{})
	/*
		Log a message with optional arguments at INFO level. Arguments are handled in the manner of fmt.Printf.
	*/
	Info(format string, v ...interface{})
	/*
		Log a message with optional arguments at WARN level. Arguments are handled in the manner of fmt.Printf.
	*/
	Warn(format string, v ...interface{})
	/*
		Log a message with optional arguments at ERROR level. Arguments are handled in the manner of fmt.Printf.
	*/
	Error(format string, v ...interface{})
	/*
		Return a logger with the specified field set so that they are included in subsequent logging calls.
	*/
	WithField(key string, v interface{}) Logger
	/*
		Return a logger with the specified fields set so that they are included in subsequent logging calls.
	*/
	WithFields(fields map[string]interface{}) Logger
	/*
		Returns the fields set in this logger.
	*/
	Fields() map[string]interface{}
}
