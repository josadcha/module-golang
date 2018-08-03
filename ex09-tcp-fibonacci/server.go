package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

var hash = make(map[int64]*big.Int)

func Search(n int64) (bool, *big.Int) {
	val, ok := hash[n]
	if ok {
		return true, val
	}
	return false, nil
}

func fibonacci(limit int64) *Write {
	start := time.Now()
	if limit < 0 {
		panic("Negative arguments not implemented")
	}

	var res *big.Int
	var b bool
	b, res = Search(limit)
	if !b {
		res, _ = fib(limit)
		hash[limit] = res
	}
	result := &Write{time.Since(start).String(), res.String()}
	return result
}

func fib(n int64) (*big.Int, *big.Int) {
	if n == 0 {
		return big.NewInt(0), big.NewInt(1)
	}
	a, b := fib(n / 2)
	c := Mul(a, Sub(Mul(b, big.NewInt(2)), a))
	d := Add(Mul(a, a), Mul(b, b))
	if n%2 == 0 {
		return c, d
	} else {
		return d, Add(c, d)
	}
}

func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}
func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}
func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

type number struct {
	Number string `json: "number"`
}

type Write struct {
	Time   string
	Result string
}

func Server(port int) {

	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		fmt.Printf("Listen port %d failed,%s\n", port, err)
		os.Exit(1)
	}
	fmt.Printf("Begin listen port: %d\n", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {

	defer conn.Close()

	var (
		buff = make([]byte, 1024)
		r    = bufio.NewReader(conn)
		w    = bufio.NewWriter(conn)
		num  = number{}
		data = int64(0)
	)

ILOOP:
	for {
		n, err := r.Read(buff)

		if err := json.Unmarshal(buff[:n], &num); err != nil {
			fmt.Printf("Decoding error reading from cl: \n", err)
		}

		data, _ = strconv.ParseInt(string(num.Number), 10, 64)

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			break ILOOP
		default:
			fmt.Printf("Receive data failed:%s\n", err)
			return
		}

	}
	res := fibonacci(data)
	mes, _ := json.Marshal(res)
	w.Write(mes)
	w.Flush()
}

func main() {
	port := 3333
	Server(port)
	hash = nil
}
