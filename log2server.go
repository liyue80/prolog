package prolog

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
	"time"
)

// SocketWrite ..
func SocketWrite(serverAddr string, b []byte) (err error) {
	client := &http.Client{}
	url := serverAddr
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, []byte(time.Now().Format("2006-01-02 15:04:05")))
	binary.Write(buf, binary.LittleEndian, b)

	request, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	//stdout := os.Stdout
	//_, err = io.Copy(stdout, response.Body)

	status := response.StatusCode
	fmt.Println(status)

	return
}
