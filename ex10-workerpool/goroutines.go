package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func Worker(id int, wg *sync.WaitGroup, sec <-chan float64) {
	defer wg.Done()
	var flag bool
	for j := range sec {
		if flag == false {
			fmt.Printf("worker:%v spawning\n", id)
		}

		flag = true
		fmt.Printf("worker:%v sleep:%.1f\n", id, j)
		time.Sleep(time.Duration(int(j*1000)) * time.Millisecond)
	}

	if flag == true {
		fmt.Printf("worker:%v stopping\n", id)
	}
}

func Run(poolSize int) {
	var wg sync.WaitGroup
	jobs := make(chan float64, 100)
	read := bufio.NewScanner(os.Stdin)

	for read.Scan() {
		f := string(read.Bytes())
		s, _ := strconv.ParseFloat(f, 64)
		jobs <- s
	}

	for i := 1; i < poolSize+1; i++ {
		wg.Add(1)
		go Worker(i, &wg, jobs)
		time.Sleep(1 * time.Millisecond)
	}

	close(jobs)

	wg.Wait()
}
