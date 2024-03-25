package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type ReadResult struct {
	Error error
	Bytes []byte
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	inputFilePath := filepath.Join(currentDir, "source-all.txt")
	outputFilePath := filepath.Join(currentDir, "mtr/ex5", "dst-all.txt")

	rChannel := readBytesFromFile(ctx, inputFilePath)

	err = writeBytesToFile(ctx, outputFilePath, rChannel)
	if err != nil {
		cancel()
	}

	fmt.Printf("Input file: %s\nOutput file: %s\n", inputFilePath, outputFilePath)

	// other ...
}

func readBytesFromFile(ctx context.Context, inputFilePath string) chan ReadResult {
	ch := make(chan ReadResult)

	go func() {
		defer close(ch)
		var n int
		file, err := os.Open(inputFilePath)
		if err != nil {
			ch <- ReadResult{Error: err}
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
			select {
			case <-ctx.Done():
				ch <- ReadResult{Error: ctx.Err()}
				return
			default:
			}

			n, err = file.Read(buffer)
			if n == 0 {
				break
			}

			if err != nil {
				ch <- ReadResult{Error: err}
				return
			}

			ch <- ReadResult{Bytes: buffer}
		}
	}()

	return ch
}

func writeBytesToFile(ctx context.Context, outputFilePath string, rChannel <-chan ReadResult) error {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}

	defer func(outputFile *os.File) {
		err = outputFile.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(outputFile)

	for ch := range rChannel {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if ch.Error != nil {
			return err
		}

		_, err = outputFile.Write(ch.Bytes)
		if err != nil {
			return err
		}
	}

	return nil
}
