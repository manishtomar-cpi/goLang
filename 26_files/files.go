package main

import (
	"fmt"
	"os"
)

/*
- We can use OS package for files
*/
func main() {
	f, err := os.Open("example.txt") // we can use OS package for files

	// error handling
	if err != nil {
		// log the error or
		panic(err)
	}

	defer f.Close()

	//applying stat method -> file info
	fileInfo, err := f.Stat()
	if err != nil {
		// log the error or
		panic(err)
	}

	//FILE INFO
	fmt.Println("+++++++++FILE INFO+++++++++")
	fmt.Println("file name:", fileInfo.Name())         // reading file name
	fmt.Println("is Folder:", fileInfo.IsDir())        // is folder
	fmt.Println("file size:", fileInfo.Size())         // in byte
	fmt.Println("file mofify at:", fileInfo.ModTime()) // when modify

	//REading from file
	fmt.Println("+++++++++Reading from file++++++++++++")
	//First store the file in the buffer -> buffer is the array of byte
	buf := make([]byte, fileInfo.Size()) //making buffer
	d, err := f.Read(buf)                // storing file data inside buffer, d basically number of bytes read

	if err != nil {
		panic(err)
	}

	// println("data:", d, buf) //data: 12 [12/12]0x1400002be28 -> we need to convert the buffer in string
	for i := 0; i < len(buf); i++ {
		println("data", d, string(buf[i])) // convert to string each byte
	}

	//reading file simply
	fmt.Println("+++++Reading Data Directly+++++++++")

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data)) //hello golang

	//create a file
	fil, err := os.Create("example2.txt")
	if err != nil {
		panic(err)
	}
	defer fil.Close()
	fil.WriteString("hello go lang") // writing inside the file
}
