package main

import(
	//"bufio" //缓存IO
	"fmt"
	//"io"
	//"io/ioutil" //io 工具包
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func createFile(fileName string){

	if checkFileIsExist(fileName) { //如果文件存在
		fmt.Println("文件存在")
		f, err := os.OpenFile(fileName, os.O_APPEND, 0777) //打开文件
		defer f.Close()
		if err != nil {
			fmt.Println("无法打开")
		}
		
	} else {
		fmt.Println("文件不存在")
		f, err := os.Create(fileName) //创建文件
		defer f.Close()
		if err != nil {
			fmt.Println("文件创建失败")
		}
		fmt.Println("文件创建成功")	
	}
}