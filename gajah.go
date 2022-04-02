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
)

func cek(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	salah := cp.Copy("src", "build")
	cek(salah)

	pengecualian := []string{}

	// ambil ignore.txt
	if _, err := os.Stat("ignore.txt"); err == nil {
		ignore, err := os.ReadFile("ignore.txt")
		cek(err)

		isiIgnore := string(ignore)
		pengecualian = strings.Split(isiIgnore, "\n")
	}
	// for n := range pengecualian {
	// 	print(pengecualian[n])
	// }

	pengecualianHtml := []string{}
	for n := range pengecualian {
		aturan_regex := regexp.MustCompile("([a-z0-9]).php")
		jadiHtml := aturan_regex.ReplaceAllString(pengecualian[n], "$1.html")
		pengecualianHtml = append(pengecualianHtml, jadiHtml)
	}
	// for n := range pengecualianHtml {
	// 	print(pengecualianHtml[n])
	// }

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
		if strings.Contains(semuaFile[n], ".js") && !strings.Contains(semuaFile[n], ".min.js") {
			fileJs = append(fileJs, semuaFile[n])
		}
		if strings.Contains(semuaFile[n], ".css") {
			fileCss = append(fileCss, semuaFile[n])
		}
	}

	// println("CSS:")
	// for n := range fileCss {
	// 	println(fileCss[n])
	// }

	if len(os.Args) == 3 && os.Args[2] == "minify" {
		for n := range fileCss {
			isi, err := os.ReadFile(fileCss[n])
			cek(err)
			isiString := string(isi)

			// println(isiString)

			m := minify.New()
			m.AddFunc("text/css", css.Minify)

			isiString, err = m.String("text/css", isiString)
			cek(err)

			// println(isiString)

			ioutil.WriteFile(fileCss[n], []byte(isiString), 0755)
		}

		// println("JS:")
		for n := range fileJs {
			isi, err := os.ReadFile(fileJs[n])
			cek(err)
			isiString := string(isi)

			// println(isiString)

			m := minify.New()
			m.AddFunc("text/javascript", js.Minify)

			isiString, err = m.String("text/javascript", isiString)
			cek(err)

			// println(isiString)

			ioutil.WriteFile(fileJs[n], []byte(isiString), 0755)
		}
	}

	proses := 0
	for n := range filePhp {
		filenya := strings.Replace(filePhp[n], "build/", "", -1)
		data, _ := http.Get("http://localhost:" + port + "/" + filenya)
		isi, _ := ioutil.ReadAll(data.Body)
		isinya := string(isi)

		// kasih "pengecualian" dan "pengecualianHtml" (array) di sini
		aturan_regex := regexp.MustCompile("([a-z0-9]).php")
		isinya = aturan_regex.ReplaceAllString(isinya, "$1.html")

		for n := range pengecualian {
			isinya = strings.Replace(isinya, pengecualianHtml[n], pengecualian[n], -1)
		}

		// isinya = strings.Replace(isinya, "scrappy.html.herokuapp.com", "scrappy-php.herokuapp.com", -1)
		// isinya = strings.Replace(isinya, "v.gd%2Fcreate.html", "v.gd%2Fcreate.php", -1)
		// isinya = strings.Replace(isinya, "is.gd%2Fcreate.html", "is.gd%2Fcreate.php", -1)
		// isinya = strings.Replace(isinya, "vurl.com%2Fapi.html", "vurl.com%2Fapi.php", -1)
		// isinya = strings.Replace(isinya, "cutt.us/api.html?url", "cutt.us/api.php?url", -1)
		// isinya = strings.Replace(isinya, "tinyurl.com/api-create.html", "tinyurl.com/api-create.php", -1)

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

			isinya = strings.Replace(isinya, "<script>a,", "<script type=module>", -1)
			isinya = strings.Replace(isinya, "<script>a;", "<script type=module>", -1)
		}

		namaFile := strings.Replace(filePhp[n], ".php", ".html", -1)
		ioutil.WriteFile(namaFile, []byte(isinya), 0755)
		// di sini, hapus file php
		salah := os.Remove(filePhp[n])
		cek(salah)
		proses = proses + 1
	}
	if proses == len(filePhp) {
		println("Gajah selesai")
	}
}
