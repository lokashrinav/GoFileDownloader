package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"sync"
	"path"
	"time"
)

var filing = []string {
	"https://raw.githubusercontent.com/lokashrinav/Jake_Resume/main/Jake_Resume1.pdf",
	"https://raw.githubusercontent.com/lokashrinav/Jake_Resume/main/Jake_Resume2.pdf",
	"https://raw.githubusercontent.com/lokashrinav/Jake_Resume/main/Jake_Resume3.pdf",
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup

	for _, file := range filing {
		wg.Add(1)
		go download(file, &wg)
	}

	wg.Wait()
	duration := time.Since(startTime)
	fmt.Printf("Time taken: %v\n", duration)
}

func download(file string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(file)
	if(err != nil) {
		fmt.Println("Error with Getting File")
		return
	}
	defer resp.Body.Close() 

	filename := path.Base(file)

	newFile, err2 := os.Create(filename)
	if(err2 != nil) {
		fmt.Println("Error with Getting File")
		return
	}
	defer newFile.Close()

	_, err3 := io.Copy(newFile, resp.Body)
	if(err3 != nil) {
		fmt.Println("Error with Getting File")
		return
	}

}





