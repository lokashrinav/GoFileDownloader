package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"path"
	"time"
)

var filing = []string{
	"https://raw.githubusercontent.com/lokashrinav/Jake_Resume/main/Jake_Resume1.pdf",
	"https://raw.githubusercontent.com/lokashrinav/Jake_Resume/main/Jake_Resume2.pdf",
	"https://raw.githubusercontent.com/lokashrinav/Jake_Resume/main/Jake_Resume3.pdf",
}

func main() {
	startTime := time.Now()

	for _, file := range filing {
		err := download(file)
		if err != nil {
			fmt.Printf("Failed to download %s: %v\n", file, err)
		}
	}

	duration := time.Since(startTime)
	fmt.Printf("Time taken: %v\n", duration)
}

func download(file string) error {
	resp, err := http.Get(file)
	if err != nil {
		return fmt.Errorf("error with getting file: %w", err)
	}
	defer resp.Body.Close()

	filename := path.Base(file)

	newFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	fmt.Printf("Downloaded: %s\n", filename)
	return nil
}
