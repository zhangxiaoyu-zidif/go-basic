package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount     int
	SpaceCount  int
	NumberCount int
	Other       int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string failed: %v", err)
		return
	}

	fmt.Printf("read str successful, ret: %s\n", str)
	file, err := os.Open("D://a.log")
	if err != nil {
		fmt.Println("read file failed.")
	}
	defer file.Close()

	readerFromFile := bufio.NewReader(file)

	var count CharCount
	for {
		strFromFile, err := readerFromFile.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read string failed, err: ", err)
		}

		//fmt.Printf("%s\n", strFromFile)
		runArr := []rune(strFromFile)
		for _, v := range runArr {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumberCount++
			default:
				count.Other++
			}
		}
	}
	fmt.Println(count)
}
