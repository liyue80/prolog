package prolog

import (
	"fmt"
	"os"
	"time"
)

// Log ..
var Log Logger

type TraceLevel uint8

const (
	LevelDebug TraceLevel = iota
	LevelInfo
	LevelWarning
	LevelCritical
	LevelNone = 0x7F
)

// Logger ..
type Logger struct {
	Addr      string
	SockLevel TraceLevel

	FilePath  string
	FileLevel TraceLevel

	ConsoleLevel TraceLevel
}

// Debug ...
func Debug(format string, v ...interface{}) {
	s := formatLogger(LevelDebug, format, v...)

	if Log.Addr != "" && Log.SockLevel <= LevelDebug {
		err := SocketWrite(Log.Addr, []byte(fmt.Sprintln("[Debug]", v)))
		if err != nil {
			fmt.Println("write socket failed")
		}
	}

	if Log.FilePath != "" && Log.FileLevel <= LevelDebug {
		_, err := FileWrite(Log.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, s)
		if err != nil {
			fmt.Println("write file failed")
		}
	}

	if Log.ConsoleLevel <= LevelDebug {
		fmt.Println(s)
	}
}

// Info ..
func Info(format string, v ...interface{}) {
	s := formatLogger(LevelInfo, format, v...)

	if Log.Addr != "" && Log.SockLevel <= LevelInfo {
		err := SocketWrite(Log.Addr, []byte(fmt.Sprintln("[Info]", v)))
		if err != nil {
			fmt.Println("write socket failed")
		}
	}

	if Log.FilePath != "" && Log.FileLevel <= LevelInfo {
		_, err := FileWrite(Log.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, s)
		if err != nil {
			fmt.Println("write file failed")
		}
	}

	if Log.ConsoleLevel <= LevelInfo {
		fmt.Println(s)
	}
}

// Warning ..
func Warning(format string, v ...interface{}) {
	s := formatLogger(LevelWarning, format, v...)

	if Log.Addr != "" && Log.SockLevel <= LevelWarning {
		err := SocketWrite(Log.Addr, []byte(fmt.Sprintln("[Warning]", v)))
		if err != nil {
			fmt.Println("write socket failed")
		}
	}

	if Log.FilePath != "" && Log.FileLevel <= LevelWarning {
		_, err := FileWrite(Log.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, s)
		if err != nil {
			fmt.Println("write file failed")
		}
	}

	if Log.ConsoleLevel <= LevelWarning {
		fmt.Println(s)
	}
}

// Critical ..
func Critical(format string, v ...interface{}) {
	s := formatLogger(LevelCritical, format, v...)

	if Log.Addr != "" && Log.SockLevel <= LevelCritical {

	}

	if Log.FilePath != "" && Log.FileLevel <= LevelCritical {
		_, err := FileWrite(Log.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, s)
		if err != nil {
			fmt.Println("write log failed")
		}
	}

	if Log.ConsoleLevel <= LevelCritical {
		fmt.Println(s)
	}
}

func formatLogger(level TraceLevel, format string, v ...interface{}) string {
	prefixTime := time.Now().Format("2006-01-02 15:04:05")
	prefixLevel := ""

	switch level {
	case LevelDebug:
		prefixLevel = " [DEBUG] "
	case LevelInfo:
		prefixLevel = " [INFOR] "
	case LevelWarning:
		prefixLevel = " [WARNI] "
	case LevelCritical:
		prefixLevel = " [CRITI] "
	}

	return prefixTime + prefixLevel + fmt.Sprintf(format, v...)
}

func init() {
	Log.Addr = ""
	Log.SockLevel = LevelNone

	Log.FilePath = ""
	Log.FileLevel = LevelNone

	Log.ConsoleLevel = LevelInfo
}

func SetLogFileName(fileName string, level TraceLevel) {
	Log.FilePath = fileName
	Log.FileLevel = level
}

func SetLogSocket(addr string, level TraceLevel) {
	Log.Addr = addr
	Log.SockLevel = level
}

func SetConsoleLevel(level TraceLevel) {
	Log.ConsoleLevel = level
}
