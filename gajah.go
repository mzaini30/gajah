package main

import (
	cp "github.com/otiai10/copy"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"fmt"
)

func cek(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	salah := cp.Copy("src", "build")
	cek(salah)

	// Ambil port
	port := os.Args[1]

	// // Ambil file: filePhp
	var semuaFile = []string{}
	filepath.Walk("build", func(x string, _ os.FileInfo, _ error) error {
		semuaFile = append(semuaFile, x)
		return nil
	})

	var filePhp = []string{}
	var fileJs = []string{}
	var fileCss = []string{}
	for n := range semuaFile {
		if strings.Contains(semuaFile[n], ".php") {
			filePhp = append(filePhp, semuaFile[n])
		}
		if strings.Contains(semuaFile[n], ".js") {
			fileJs = append(fileJs, semuaFile[n])
		}
		if strings.Contains(semuaFile[n], ".css") {
			fileCss = append(fileCss, semuaFile[n])
		}
	}

	// fmt.Println("CSS:")
	// for n := range fileCss {
	// 	fmt.Println(fileCss[n])
	// }

	for n := range fileCss {
		isi, err := os.ReadFile(fileCss[n])
		cek(err)
		isiString := string(isi)

		// fmt.Println(isiString)

		m := minify.New()
		m.AddFunc("text/css", css.Minify)

		isiString, err = m.String("text/css", isiString)
		cek(err)

		// fmt.Println(isiString)

		ioutil.WriteFile(fileCss[n], []byte(isiString), 0755)
	}

	// fmt.Println("JS:")
	for n := range fileJs {
		isi, err := os.ReadFile(fileJs[n])
		cek(err)
		isiString := string(isi)

		// fmt.Println(isiString)

		m := minify.New()
		m.AddFunc("text/javascript", js.Minify)

		isiString, err = m.String("text/javascript", isiString)
		cek(err)

		// fmt.Println(isiString)

		ioutil.WriteFile(fileJs[n], []byte(isiString), 0755)
	}

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
		salah := os.Remove(filePhp[n])
		cek(salah)
		proses = proses + 1
	}
	if proses == len(filePhp) {
		fmt.Println("Gajah selesai")
	}
}
