package logger

import (
    "github.com/cihub/seelog"
    "gopkg.in/macaron.v1"
    "fmt"
    "os"
)

// 日志库

type Level int8

var logger seelog.LoggerInterface

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

func InitLogger()  {
    config := getLogConfig()
    l, err := seelog.LoggerFromConfigAsString(config)
    if err != nil {
        panic(err)
    }
    logger = l
}

func Debug(v ...interface{}) {
	write(DEBUG, v)
}

func Debugf(format string, v... interface{})  {
    writef(DEBUG, format, v)
}

func Info(v ...interface{}) {
	write(INFO, v)
}

func Infof(format string, v... interface{})  {
    writef(INFO, format, v)
}

func Warn(v ...interface{}) {
	write(WARN, v)
}

func Warnf(format string, v... interface{})  {
    writef(WARN, format, v)
}

func Error(v ...interface{}) {
	write(ERROR, v)
}

func Errorf(format string, v... interface{})  {
    writef(ERROR, format, v)
}

func Fatal(v ...interface{}) {
	write(FATAL, v)
}

func Fatalf(format string, v... interface{})  {
    writef(FATAL, format, v)
}

func write(level Level, v... interface{}) {
    defer logger.Flush()

    switch level {
        case DEBUG:
            logger.Debug(v)
        case INFO:
            logger.Info(v)
        case WARN:
            logger.Warn(v)
        case FATAL:
            logger.Critical(v)
            os.Exit(1)
        case ERROR:
            logger.Error(v)
	}
}

func writef(level Level, format string, v... interface{})  {
    defer logger.Flush()

    switch level {
        case DEBUG:
            logger.Debugf(v)
        case INFO:
            logger.Infof(v)
        case WARN:
            logger.Warnf(v)
        case FATAL:
            logger.Criticalf(v)
            os.Exit(1)
        case ERROR:
            logger.Errorf(v)
    }
}

func getLogConfig() string {
	config := `
    <seelog>
        <outputs formatid="main">
            %s
            <filter levels="info,critical,error,warn">
                <file path="log/cron.log" />
            </filter>
        </outputs>
        <formats>
            <format id="main" format="%%Date/%%Time [%%LEV] %%Msg%%n"/>
        </formats>
    </seelog>`

    consoleConfig := ""
    if macaron.Env == macaron.DEV {
        consoleConfig =
        `
            <filter levels="info,debug,critical,warn,error">
                <console />
            </filter>
         `
    }
    config = fmt.Sprintf(config, consoleConfig)

    return config
}
