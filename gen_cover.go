package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var (
	r    = regexp.MustCompile(`[a-z0-9-]`)
	host = getEnv("HOST", "localhost")
	port = getEnv("PORT", "8091")
	addr = getEnv("ADDR", fmt.Sprintf("%s:%s", host, port))
)

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("listening on TCP", addr)

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		log.Println("Error reading:", err.Error())
	}

	defer conn.Close()

	group := r.FindAllString(strings.ReplaceAll(string(buf), " ", "-"), -1)
	room := strings.Join(group, "")
	size := len(room)

	if size == 0 || size > 20 {
		return
	}

	img, err := gen(room)
	if err != nil {
		log.Println(err)
		return
	}

	buffer_size := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer_size, uint32(len(img.Bytes())))

	if _, err := conn.Write(buffer_size); err != nil {
		log.Println(err)
	}

	if _, err := conn.Write(img.Bytes()); err != nil {
		log.Println(err)
	}
}
