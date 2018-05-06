package main

import (
	"bufio"
	"os"
)

func main() {
	outputFile, outputErr := os.OpenFile("D://a.txt", os.O_CREATE|os.O_RDWR, 0644)
	if outputErr != nil {
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "Hello world."
	outputWriter.WriteString(outputString)

	//由内存写入磁盘。必须加上
	outputWriter.Flush()
}
