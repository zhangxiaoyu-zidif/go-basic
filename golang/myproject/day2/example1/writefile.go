package main

import(
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeFile(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		fmt.Println("could not open a file!")
	}
	//var writeString string
	//fmt.Scanf("%s", &writeString)
	//fmt.Println(writeString)

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	_, err = io.WriteString(f, input)
	if err != nil {
		fmt.Println("write failed")
	}
}