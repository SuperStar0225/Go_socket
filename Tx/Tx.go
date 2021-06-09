package main

import (
	jsonsocket "github.com/Txiaozhe/go-json-socket"
	"log"
)

type Person struct {
    Name   string
    Emails []string
}

func main() {
	listener, err := jsonsocket.Listen("127.0.0.1:3000")
	if err != nil {
		log.Println("listen error: ", err)
	}
	person := Person{Name: "Jan",
        Emails: []string{"ja@newmarch.name", "jan.newmarch@gmail.com"},
    }
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error: ", err)
		}

		ch, err := jsonsocket.HandleMessage(conn)
		if err != nil {
			log.Println("handel message error: ", err)
		}

		res := <-ch
		log.Println(res)

		ch1, err := jsonsocket.SendMessage(conn, person)
		if err != nil {
			log.Println("response error:", err)
		}
		log.Println("send msg length: ", <-ch1)
	}
}