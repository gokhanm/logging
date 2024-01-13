# logging
Golang logging library with logrus and graylog hook support

```go
import log "github.com/gokhanm/logging"

func main() {
    // you can set your own logrus formatter. text, json ..
    formatter = &log.TextFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
        ForceColors:     true,
        FullTimestamp:   true,
    }
    logger := log.Initialize(formatter)
    // Initialize returns logrus Logger pointer. 
    // with this you can use logrus functions if your want
    logger.SetLevel(logrus.DebugLevel)
}
```

Multi-custom logrus fields and async graylog hook example
```go
import log "github.com/gokhanm/logging"

func main() {

    // or you can use the default text logrus formatter in the lib.
    log.Initialize(nil)

    // you can have fields always attached to log statements by SetDefaultFields in an application
    fields := map[string]interface{}{"name": "gokhan", "app": "appName"}
    log.SetDefaultFields(fields)

    // also if you use graylog log management system
    // you can send your log data to the graylog as async 
    log.AddAsyncGraylogHook("127.0.0.1", "1000", nil)
}
```
