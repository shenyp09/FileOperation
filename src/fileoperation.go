package main

import "fmt"
import "os"

func main() {
	fileName := "test.dat"
	dstFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()
	s := "hello world"
	dstFile.WriteString(s + "\n")

}
