package main

import (
	// "fmt"
	"os"
	"path/filepath"
	// "log"
	"strings"
	"net/http"
	"io/ioutil"
)

func main(){
	// Ambil port
	port := os.Args[1]

	// Ambil file: file_php
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

	for n := range file_php {
		data, _ := http.Get("http://localhost:", port, "/", file_php[n])
		nama_file := strings.Replace(file_php[n], ".php", ".html", -1)
		ioutil.WriteFile(nama_file, [] byte (data), 0755)
	}
}
