package prolog

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"
)

// FileWrite ..
func FileWrite(fileName string, flag int, s string) (n int, err error) {
	dir := path.Dir(fileName)
	os.Mkdir(dir, 0777)

	createDate := fmt.Sprintf("%04d%02d%02d", time.Now().Year(), time.Now().Month(), time.Now().Day())
	f, err := os.OpenFile(fileName+createDate+".log", flag, 0)
	if err != nil {
		return
	}
	defer f.Close()

	wr := bufio.NewWriter(f)
	if _, err = wr.Write([]byte(s)); err == nil {
		wr.Flush()
	}

	return
}
