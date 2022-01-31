package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func jalankan_server(wg *sync.WaitGroup) {
	defer wg.Done()
	port := os.Args[1]
	fmt.Println("Server berjalan di localhost:" + port)
	exec.Command("bash", "-c", "cd src && php -S localhost:"+port).Run()
}

func dapatkan_data(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Ambil data")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go jalankan_server(&wg)
	go dapatkan_data(&wg)
}
