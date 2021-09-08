package demo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func DemoIO() {

	fmt.Println("*************** Demo I/O ***************")

	demoWriteFile()

	demoReadFile()

}

func demoWriteFile() {

	fmt.Println("--- Try Write File ---")

	file, err := os.Create("tmp/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hi... there")
	file.Close()

}

func demoReadFile() {

	fmt.Println("--- Try Read File ---")

	stream, err := ioutil.ReadFile("tmp/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)

}
