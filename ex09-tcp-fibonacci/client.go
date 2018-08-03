package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type write struct {
	Number string `json: "number"`
}

type read struct {
	Time   string
	Result string
}

//const StopCharacter = "\r\n\r\n"

func Client(ip string, port int, num *write) {
	location := read{}

	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}

	mes, _ := json.Marshal(num)
	conn.Write(mes)
	//conn.Write([]byte(StopCharacter))

	buff := make([]byte, 1000000000)
	n, _ := conn.Read(buff)

	if err := json.Unmarshal(buff[:n], &location); err != nil {
		fmt.Println("Decoding error reading from server: ", err)
	}

	fmt.Println(location.Time, location.Result)
}

func main() {
	var (
		ip   = "127.0.0.1"
		port = 3333
	)

	read := bufio.NewScanner(os.Stdin)

	for read.Scan() {
		f := &write{string(read.Bytes())}
		Client(ip, port, f)
	}
}
