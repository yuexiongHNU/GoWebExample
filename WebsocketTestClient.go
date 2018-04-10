package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"net/url"
	"github.com/gorilla/websocket"
	"time"
)

// define parameters
var addr2 = flag.String("addr", "localhost:8080", "http service address")

func main()  {
	// parse parameters
	flag.Parse()
	// set format of logs, remove the time heading
	log.SetFlags(0)

	// new channel to receive the os interrupt
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme:"ws", Host:*addr2, Path:"/echo"}
	log.Printf("connecting to %s", u.String())

	// new ws client connection
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("ws dial error:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	// call anonymous function in complicated way
	go func() {
		defer close(done)
		for {
			t, message, err := c.ReadMessage()
			log.Println("message t:", t)
			if err != nil {
				log.Println("read message error:",err)
				return
			}
			log.Printf("receive: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}