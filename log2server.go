package prolog

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net/http"
)

// SocketWrite ..
func SocketWrite(serverAddr string, s string) (err error) {
	client := &http.Client{}
	url := serverAddr
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, []byte(s))

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
