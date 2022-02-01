package main

import (
	// "fmt"
	"os"
	"path/filepath"
	// "log"
	"strings"
)

func main(){
	// Ambil port
	// port := os.Args[1]

	// Ambil file
	var semua_file = [] string {}
	filepath.Walk(".", func(x string, _ os.FileInfo, _ error) error {
		semua_file = append(semua_file, x)
		return nil
	})
	var file_php = [] string {}
	for n := range semua_file {
		if strings.Contains(semua_file[n], ".php"){
			file_php = append(file_php, semua_file[n])
		}
	}
}
