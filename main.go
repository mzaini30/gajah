package main

import (
	"os"
	"os/exec"
	"fmt"
)

func jalankan_server() {
	port := os.Args[1]
	fmt.Println("Server berjalan di localhost:" + port)
	exec.Command("bash", "-c", "cd src && php -S localhost:"+port).Run()
}

func main() {
	jalankan_server()
}
