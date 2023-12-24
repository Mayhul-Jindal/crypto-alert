package main

func main(){

}













// import (
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"golang.org/x/net/websocket"
// )

// type Server struct {
// 	conn map[*websocket.Conn]bool
// }

// func NewServer() *Server {
// 	return &Server{
// 		conn: make(map[*websocket.Conn]bool),
// 	}
// }

// func (s *Server) handleWs(ws *websocket.Conn) {
// 	fmt.Println("New conn:", ws.LocalAddr())

// 	s.conn[ws] = true
// 	defer ws.Close()
// 	defer delete(s.conn, ws)

// 	s.readLoop(ws)
// }

// func (s *Server) readLoop(ws *websocket.Conn) {
// 	buf := make([]byte, 1024)
// 	addr := ws.LocalAddr()
// 	for {
// 		n, err := ws.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				fmt.Println("EOF")
// 				return
// 			}

// 			fmt.Println("read error", err)
// 			continue
// 		}

// 		fmt.Println(addr, "-->", string(buf[:n]))
// 		ws.Write([]byte("hello from server"))
// 	}
// }

// func main() {
// 	server := NewServer()
// 	http.Handle("/ws", websocket.Handler(server.handleWs))
// 	http.ListenAndServe(":3000", nil)
// }
