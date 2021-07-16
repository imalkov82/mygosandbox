package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func lineCounter(r io.Reader) (int, error) {
	log.Printf("lineCounter ->")
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		// if count%10000 == 0 {
		// 	log.Printf("count : %v", count)
		// }
		switch {
		case err == io.EOF:
			log.Printf("<- lineCounter: err == io.EOF")
			return count, nil

		case err != nil:
			log.Printf("<- lineCounter: err == io.EOF")
			return count, err
		}
	}
}

func main() {
	file, _ := os.Open("/Users/imalkov/Documents/EPS/tmp/bwseeds.csv")
	c, err := lineCounter(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("lines=%d\n", c)
}
