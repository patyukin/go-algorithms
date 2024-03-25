package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errorCB := func(err error) {
		fmt.Println("Error:", err.Error())
	}

	resultCB := func(pg float64) {
		fmt.Println("Page:", pg)
	}

	UploadMany(ctx, []string{"source1.txt", "source2.txt"}, []string{"destination1.txt", "destination2.txt"}, errorCB, resultCB)
}

func UploadMany(ctx context.Context, src, dst []string, errorCallback func(error), resultCallback func(float64)) {
	var wg sync.WaitGroup

	for i, s := range src {
		wg.Add(1)
		go func(ctx context.Context, src, dst string) {
			defer wg.Done()

			err := Upload(ctx, src, dst, errorCallback, resultCallback)

			fmt.Printf("Upload %s to %s: %v\n", src, dst, p)
		}(ctx, s, dst[i])
	}

	wg.Wait()
}

func readBatchFromFile(file *os.File) []byte {
	batchSize := 1024
	batch := make([]byte, batchSize)
	n, err := file.Read(batch)
	if err != nil {
		return nil
	}

	return batch[:n]
}

func Upload(ctx context.Context, src, dst string, errorCallback func(error), resultCallback func(float64)) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		currentDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Error getting current directory: %w", err)
		}

		filePath := filepath.Join(currentDir, "mtr/ex4", src)
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("Error opening file: %w", err)
		}

		go func() {
			defer func(file *os.File) {
				err = file.Close()
				if err != nil {
					fmt.Printf("Error closing file: %v\n", err)
				}
			}(file)
		}()

		r := readBatchFromFile(file)
		if len(r) == 0 {
			errorCallback(errors.New("empty batch"))
			return nil
		}

		pg := uploadBatch(r)

		dstFilePath := filepath.Join(currentDir, "mtr/ex4", dst)
		dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			errorCallback(err)
			return nil
		}

		go func() {
			defer func(dstFile *os.File) {
				err = dstFile.Close()
				if err != nil {
					fmt.Printf("Error closing file: %v\n", err)
				}
			}(dstFile)
		}()

		_, err = dstFile.Write(r)
		if err != nil {
			errorCallback(err)
			return nil
		}

		resultCallback(pg)
	}
}

func uploadBatch(batch []byte) float64 {
	return float64(len(batch)) / 1024.0 * 100.0
}
