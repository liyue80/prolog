package prolog

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Log ..
var Log Logger

const (
	outNone = iota
	outConsole
	outFile
	outSocket
)

// Logger ..
type Logger struct {
	Addr        string
	FilePath    string
	DebugOut    []byte
	InfoOut     []byte
	WarningOut  []byte
	CriticalOut []byte
}

// Debug ...
func Debug(v ...interface{}) {
	for _, p := range Log.DebugOut {
		switch p {
		case outNone:
		case outConsole:
			fmt.Printf("%v\n", v)
		case outFile:
			_, err := FileWrite(Log.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, []byte(fmt.Sprintln("[Debug]", v)))
			if err != nil {
				fmt.Println("write file failed")
			}
		case outSocket:
			err := SocketWrite(Log.Addr, []byte(fmt.Sprintln("[Debug]", v)))
			if err != nil {
				fmt.Println("write socket failed")
			}
		}
	}
}

// Info ..
func Info(v ...interface{}) {
	for _, p := range Log.DebugOut {
		switch p {
		case outNone:
		case outConsole:
			fmt.Printf("%v\n", v)
		case outFile:
			_, err := FileWrite(Log.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, []byte(fmt.Sprintln("[Info]", v)))
			if err != nil {
				fmt.Println("write file failed")
			}
		case outSocket:
			err := SocketWrite(Log.Addr, []byte(fmt.Sprintln("[Info]", v)))
			if err != nil {
				fmt.Println("write socket failed")
			}
		}
	}
}

// Warning ..
func Warning(v ...interface{}) {
	for _, p := range Log.DebugOut {
		switch p {
		case outNone:
		case outConsole:
			fmt.Printf("%v\n", v)
		case outFile:
			_, err := FileWrite(Log.FilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, []byte(fmt.Sprintln("[Warning]", v)))
			if err != nil {
				fmt.Println("write file failed")
			}
		case outSocket:
			err := SocketWrite(Log.Addr, []byte(fmt.Sprintln("[Warning]", v)))
			if err != nil {
				fmt.Println("write socket failed")
			}
		}
	}
}

// Critical ..
func Critical(v ...interface{}) {
	for _, p := range Log.CriticalOut {
		switch p {
		case outNone:
		case outConsole:
			fmt.Printf("%v\n", v)
		case outFile:
			filename := Log.FilePath
			_, err := FileWrite(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, []byte(fmt.Sprintln("[Critical]", v)))
			if err != nil {
				fmt.Println("write log failed")
			}
		case outSocket:
		}
	}
}

func init() {
	if err := readConfig(); err != nil {
		Log.Addr = "127.0.0.1"
		Log.FilePath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		Log.DebugOut = []byte{outConsole}
		Log.InfoOut = []byte{outConsole}
		Log.WarningOut = []byte{outFile, outSocket}
		Log.CriticalOut = []byte{outConsole, outFile, outSocket}
	}
}

func readConfig() error {
	type LogConfig struct {
		ServerAddr     string `xml:"server"`
		FileName       string `xml:"localfilepath"`
		DebugOutput    string `xml:"debugoutput"`
		Infooutput     string `xml:"infooutput"`
		Warningoutput  string `xml:"warningoutput"`
		Criticaloutput string `xml:"criticaloutput"`
	}

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	fileName := path + `\conf\config.xml`
	config := LogConfig{}

	var content []byte
	if content, err = ioutil.ReadFile(fileName); err != nil {
		return err
	}

	if err = xml.Unmarshal(content, &config); err != nil {
		return err
	}

	Log.Addr = config.ServerAddr
	Log.FilePath = config.FileName
	for _, s := range strings.Join(strings.Split(config.DebugOutput, ","), "") {
		b, _ := strconv.Atoi(string(s))
		Log.DebugOut = append(Log.DebugOut, uint8(b))
	}
	for _, s := range strings.Join(strings.Split(config.Infooutput, ","), "") {
		b, _ := strconv.Atoi(string(s))
		Log.InfoOut = append(Log.InfoOut, uint8(b))
	}
	for _, s := range strings.Join(strings.Split(config.Warningoutput, ","), "") {
		b, _ := strconv.Atoi(string(s))
		Log.WarningOut = append(Log.WarningOut, uint8(b))
	}
	for _, s := range strings.Join(strings.Split(config.Criticaloutput, ","), "") {
		b, _ := strconv.Atoi(string(s))
		Log.CriticalOut = append(Log.CriticalOut, uint8(b))
	}

	return nil
}
