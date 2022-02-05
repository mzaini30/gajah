package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	// "log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Ambil port
	port := os.Args[1]

	// Ambil file: file_php
	var semua_file = []string{}
	filepath.Walk(".", func(x string, _ os.FileInfo, _ error) error {
		semua_file = append(semua_file, x)
		return nil
	})
	var file_php = []string{}
	for n := range semua_file {
		if strings.Contains(semua_file[n], ".php") {
			file_php = append(file_php, semua_file[n])
		}
	}

	proses := 0
	for n := range file_php {
		data, _ := http.Get("http://localhost:" + port + "/" + file_php[n])
		isi, _ := ioutil.ReadAll(data.Body)
		isinya := string(isi)

		aturan_regex := regexp.MustCompile("([a-z0-9]).php")
		isinya = aturan_regex.ReplaceAllString(isinya, "$1.html")

		if len(os.Args) == 3 && os.Args[2] == "minify" {
			isinya = strings.Replace(isinya, "<script type=\"module\">", "<script>a;", -1)
			isinya = strings.Replace(isinya, "<script type='module'>", "<script>a;", -1)

			m := minify.New()
			m.AddFunc("text/css", css.Minify)
			m.AddFunc("text/html", html.Minify)
			m.AddFunc("image/svg+xml", svg.Minify)
			m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
			m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
			m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
			isinya, _ = m.String("text/html", isinya)

			isinya = strings.Replace(isinya, "<script>a,", "<script type=\"module\">", -1)
			isinya = strings.Replace(isinya, "<script>a;", "<script type=\"module\">", -1)
		}

		nama_file := strings.Replace(file_php[n], ".php", ".html", -1)
		ioutil.WriteFile(nama_file, []byte(isinya), 0755)
		proses = proses + 1
	}
	if proses == len(file_php) {
		fmt.Println("Gajah selesai")
	}
}
