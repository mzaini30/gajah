package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"io/ioutil"
	cp "github.com/otiai10/copy"
	"net/http"
	"strings"
)

func main() {
	salah := cp.Copy("src", "build")
	if salah != nil {
		fmt.Println(salah)
	}

	// Ambil port
	port := os.Args[1]

	// // Ambil file: filePhp
	var semuaFile = []string{}
	filepath.Walk("build", func(x string, _ os.FileInfo, _ error) error {
		semuaFile = append(semuaFile, x)
		return nil
	})
	// fmt.Println(semuaFile)

	var filePhp = []string{}
	for n := range semuaFile {
		if strings.Contains(semuaFile[n], ".php") {
			filePhp = append(filePhp, semuaFile[n])
		}
	}
	// fmt.Println(filePhp)

	proses := 0
	for n := range filePhp {
		filenya := strings.Replace(filePhp[n], "build/", "", -1)
		data, _ := http.Get("http://localhost:" + port + "/" + filenya)
		isi, _ := ioutil.ReadAll(data.Body)
		isinya := string(isi)

		aturan_regex := regexp.MustCompile("([a-z0-9]).php")
		isinya = aturan_regex.ReplaceAllString(isinya, "$1.html")

		if len(os.Args) == 3 && os.Args[2] == "minify" {
			isinya = strings.Replace(isinya, "<script type=\"module\">", "<script>a;", -1)
			isinya = strings.Replace(isinya, "<script type='module'>", "<script>a;", -1)

			isinya = strings.Replace(isinya, "<script type=\"magic\">", "<script>b;", -1)
			isinya = strings.Replace(isinya, "<script type='magic'>", "<script>b;", -1)

			isinya = strings.Replace(isinya, "<script type=\"magic\" data-type=\"module\">", "<script>c;", -1)
			isinya = strings.Replace(isinya, "<script type='magic' data-type='module'>", "<script>c;", -1)

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

			isinya = strings.Replace(isinya, "<script>b,", "<script type=\"magic\">", -1)
			isinya = strings.Replace(isinya, "<script>b;", "<script type=\"magic\">", -1)

			isinya = strings.Replace(isinya, "<script>c,", "<script type=\"magic\" data-type=\"module\">", -1)
			isinya = strings.Replace(isinya, "<script>c;", "<script type=\"magic\" data-type=\"module\">", -1)
		}

		namaFile := strings.Replace(filePhp[n], ".php", ".html", -1)
		ioutil.WriteFile(namaFile, []byte(isinya), 0755)
		// di sini, hapus file php
		// fmt.Println(filePhp[n])
		salah := os.Remove(filePhp[n])
		if salah != nil {
			fmt.Println(salah)
		}
		proses = proses + 1
	}
	if proses == len(filePhp) {
		fmt.Println("Gajah selesai")
	}
}
