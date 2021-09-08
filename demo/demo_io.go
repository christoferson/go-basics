package demo

import (
	"fmt"
	"log"
	"os"
)

func DemoIO() {

	fmt.Println("*************** Demo I/O ***************")

	demoWriteFile()

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
