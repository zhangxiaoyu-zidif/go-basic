package main

import(
	"bufio"
	"os"
	"fmt"
	"io"
)

func readfile(fileName string) {
    fi, err := os.Open(fileName)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    defer fi.Close()

    br := bufio.NewReader(fi)
    for {
        a, _, c := br.ReadLine()
        if c == io.EOF {
            break
        }
        fmt.Println(string(a))
    }
}