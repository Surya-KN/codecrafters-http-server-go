package main

import (
	"fmt"
	"strconv"
	"strings"

	// Uncomment this block to pass the first stage
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
	fmt.Println("Failed to bind to port 4221")
	os.Exit(1)
	}
	//
  conn, err := l.Accept()
	if err != nil {
	fmt.Println("Error accepting connection: ", err.Error())
	os.Exit(1)

	}
  buf := make([]byte,5000)
  _ , err = conn.Read(buf)

  if err != nil {
	fmt.Println("Error reading:", err.Error())
	os.Exit(1)
  }
//   fmt.Println(string(buf))

  path := strings.Split(string(buf), "\r\n")
  pathfirst := strings.Split(path[0], " ")
  output := strings.Split(pathfirst[1], "/")
  useragent := []string{}
  for i := 0; i < len(path); i++ {
	if strings.Contains(path[i], "User-Agent") {
	  useragent = strings.Split(path[i], " ")
	}
}
//   fmt.Println(useragent)

  if pathfirst[1] == "/" {
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
  } else if output[1] == "echo" {
	// fmt.Println(path[1][6:])
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Content-Type: text/plain\r\n"))
	conn.Write([]byte("Content-Length: " + strconv.Itoa(len(pathfirst[1][6:])) +"\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write([]byte(pathfirst[1][6:]))
  } else if output[1] == "user-agent" {
	fmt.Println(useragent[1])
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Content-Type: text/plain\r\n"))
	conn.Write([]byte("Content-Length: " + strconv.Itoa(len(useragent[1]) ) +"\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write([]byte(useragent[1]))
  } else {
	conn.Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\n"))
  }
  
  conn.Close()
}
