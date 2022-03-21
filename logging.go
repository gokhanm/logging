package logging

import (
	"io"
	"io/ioutil"
	"os"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
)

var (
	logger *log.Logger
	field  map[string]interface{}
)

// WriterHook is a hook that writes logs of specified LogLevels to specified Writer
type WriterHook struct {
	Writer    io.Writer
	LogLevels []log.Level
}

// init is the initiliaze logrus logger
func init() {
	logger = log.New()
}

// Initialize is using default set text formatter if formatter is nil.
func Initialize(formatter log.Formatter) *log.Logger {
	if formatter == nil {
		formatter = &log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
			FullTimestamp:   true,
		}
	}
	logger.SetFormatter(formatter)

	return logger
}

//SetDefaultFields, you can have fields always attached to log statements in application
func SetDefaultFields(fields map[string]interface{}) {
	field = fields
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (hook *WriterHook) Fire(entry *log.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *WriterHook) Levels() []log.Level {
	return hook.LogLevels
}

// SplitLogs adds hooks to send logs to different destinations depending on level
func SplitLogs() {
	logger.SetOutput(ioutil.Discard) // Send all logs to nowhere by default

	logger.AddHook(&WriterHook{ // Send logs with level higher than warning to stderr
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
	logger.AddHook(&WriterHook{ // Send info and debug logs to stdout
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
			log.InfoLevel,
			log.DebugLevel,
		},
	})
}

// AddAsyncGraylogHook also send log data to graylog
func AddAsyncGraylogHook(host, port string, extra map[string]interface{}) {
	addr := host + ":" + port
	hook := graylog.NewAsyncGraylogHook(addr, extra)
	hook.Level = log.WarnLevel
	logger.AddHook(hook)
}

// Fatalln is equivalent to Logln() followed by a call to os.Exit(1)
func Fatalln(args ...interface{}) {
	logger.WithFields(field).Fatalln(args...)
}

// Fatalf is equivalent to Logf() followed by a call to os.Exit(1)
func Fatalf(format string, args ...interface{}) {
	logger.WithFields(field).Fatalf(format, args...)
}

// Println calls Output to print to the standard logger
func Println(args ...interface{}) {
	logger.WithFields(field).Println(args...)
}

// Printf calls Output to print to the standard logger
func Printf(format string, args ...interface{}) {
	logger.WithFields(field).Printf(format, args...)
}

// Errorf calls Output to print to the standard logger with error level
func Errorf(format string, args ...interface{}) {
	logger.WithFields(field).Errorf(format, args...)
}

// Errorln calls Output to print to the standard logger with error level
func Errorln(format string, args ...interface{}) {
	logger.WithFields(field).Errorln(args...)
}

// Warnf calls Output to print to the standard logger with warm level
func Warnf(format string, args ...interface{}) {
	logger.WithFields(field).Warnf(format, args...)
}

// Warnln calls Output to print to the standard logger with warm level
func Warnln(format string, args ...interface{}) {
	logger.WithFields(field).Warnln(args...)
}

// Println calls Output to print to the standard logger
func Debugln(args ...interface{}) {
	logger.WithFields(field).Debugln(args...)
}

// Printf calls Output to print to the standard logger
func Debugf(format string, args ...interface{}) {
	logger.WithFields(field).Debugf(format, args...)
}
