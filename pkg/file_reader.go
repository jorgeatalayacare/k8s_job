package pkg

import (
	"fmt"
	"io"
	"os"
)

func ReadPrintFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		errMsg := fmt.Errorf("Error opening a file: %v", err)
		fmt.Println(errMsg)
		return err
	}
	defer file.Close()
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			errMsg := fmt.Errorf("Error reading file: %v", err)
			fmt.Println(errMsg)
			return err
		}
		if n == 0 {
			break
		}
		if _, err := os.Stdout.Write(buffer[:n]); err != nil {
			errMsg := fmt.Errorf("Error writing to stdout: %v", err)
			fmt.Println(errMsg)
			return err
		}
	}
	return nil
}
