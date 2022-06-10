package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic("failed launch server")
	}
	fmt.Println("server running")
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic("create conn failed")
		}
		go func(conn net.Conn) {
			fmt.Println("conn accept")
			buf := make([]byte, 1024)
			for {
				// Read the incoming connection into the buffer.
				reqLen, err := conn.Read(buf)
				if err != nil {
					if err.Error() == "EOF" {
						fmt.Println("Disconned ")
						break
					} else {
						fmt.Println("Error reading:", err.Error())
						break
					}
				}
				str := string(buf[:reqLen])
				i, err := strconv.Atoi(str)
				if err != nil {
					fmt.Println(len(str), reqLen, str, []byte(str))
					panic(str)
				}
				ipaddr := "10.0.0.2"
				if i > 100 {
					ipaddr = "10.0.0.4"
				}
				// new conn to B
				conn2, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ipaddr, 8080))
				if err != nil {
					return
				}
				defer conn2.Close()

				conn2.Write([]byte(strconv.Itoa(i + 1)))

				// listen for reply
				bs := make([]byte, 1024)
				len, err := conn2.Read(bs)
				if err != nil {
					return
				}

				// Send a response back
				fmt.Println(bs[:len])
				conn.Write(bs[:len])
			}
			// Close the connection when you're done with it.
			conn.Close()
		}(conn)
	}
}
