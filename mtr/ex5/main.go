package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	inputFilePath := filepath.Join(currentDir, "source-all.txt")
	outputFilePath := filepath.Join(currentDir, "mtr/ex5", "dst-all.txt")

	byteChannel := make(chan byte)
	errorChannel := make(chan error)
	doneChannel := make(chan struct{})

	go func() {
		readBytesFromFile(inputFilePath, byteChannel, errorChannel, doneChannel)
	}()

	go func() {
		writeBytesToFile(outputFilePath, byteChannel, errorChannel, doneChannel)
	}()

	fmt.Printf("Input file: %s\nOutput file: %s\n", inputFilePath, outputFilePath)

	for {
		select {
		case err = <-errorChannel:
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		case <-doneChannel:
			fmt.Println("File read and write complete")
			return
		}
	}
}

func readBytesFromFile(inputFilePath string, byteChannel chan<- byte, errorChannel chan<- error, doneChannel chan<- struct{}) {
	var n int
	file, err := os.Open(inputFilePath)
	if err != nil {
		errorChannel <- err
		return
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	buffer := make([]byte, 1)

	for {
		n, err = file.Read(buffer)
		if n == 0 {
			break
		}

		if err != nil {
			errorChannel <- err
			return
		}

		byteChannel <- buffer[0]
	}

	close(byteChannel)
	doneChannel <- struct{}{}
}

func writeBytesToFile(outputFilePath string, byteChannel <-chan byte, errorChannel chan<- error, doneChannel chan<- struct{}) {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		errorChannel <- err
		return
	}

	defer func(outputFile *os.File) {
		err = outputFile.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(outputFile)

	for b := range byteChannel {
		_, err = outputFile.Write([]byte{b})
		if err != nil {
			errorChannel <- err
			return
		}
	}

	doneChannel <- struct{}{}
}
