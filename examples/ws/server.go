package main

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("error:", err)
			break
		}

		// 处理接收到的消息
		var buf bytes.Buffer
		buf.WriteString("Received message: ")
		buf.Write(message)
		c.send <- buf.Bytes()
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()

	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case message := <-c.send:
			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("error:", err)
				break
			}
		case <-tick.C:
			err := c.conn.WriteJSON(map[string]any{"time": time.Now().Unix()})
			if err != nil {
				log.Println("error:", err)
			}
		}
	}
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error:", err)
		return
	}

	client := &Client{conn: conn}
	client.send = make(chan []byte, 256)

	go client.readPump()
	go client.writePump()
}

func main() {
	http.HandleFunc("/ws", ServeWs)
	log.Println("WebSocket server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
