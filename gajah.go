package main

import (
	"fmt"
	"os"
	"path/filepath"
	"log"
)

func main(){
	// Ambil port
	port := os.Args[1]
	fmt.Println(port)

	// Ambil file
	err := filepath.Walk(".",
	    func(path string, info os.FileInfo, err error) error {
	    if err != nil {
	        return err
	    }
	    fmt.Println(path, info.Size())
	    return nil
	})
	if err != nil {
	    log.Println(err)
	}
}
