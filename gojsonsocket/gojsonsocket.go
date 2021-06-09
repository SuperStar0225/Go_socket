package gojsonsocket

import (
	"net"
	"regexp"
	"strconv"
	"time"
	// "fmt"
)

type Response struct {
	Len  int         `json:"len"`
	Data interface{} `json:"data"`
}

func Connect(addr string) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, 15*time.Second)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func HandleMessage(conn net.Conn) (chan Response, error) {
	ch := make(chan Response)

	buf := make([]byte, 1024)

	go func() {
		_, err := conn.Read(buf)
		if err != nil {
			ch <- Response{0, err}
		}

		reg := regexp.MustCompile(`(\S*)#(\S*)`)
		str := reg.FindStringSubmatch(string(buf))

		dataLen, err := strconv.Atoi(str[1])
		if err != nil {
			ch <- Response{0, err}
		}

		res := Response{dataLen, str[2]}
		ch <- res
	}()

	return ch, nil
}