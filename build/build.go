package main

import (
	"bytes"
	"errors"
	cp "github.com/otiai10/copy"
	"go-static-website/routes"
	"go-static-website/text/en"
	"maragu.dev/gomponents"
	"os"
	"path/filepath"
)

func main() {
	path := "dist"
	must(os.RemoveAll(path))
	must(os.Mkdir(path, 0755))
	must(writePage("/index", en.IndexPage()))
	allRoutes := routes.GetRoutes()
	for route, content := range allRoutes {
		must(writePage(route, content))
	}
	must(cp.Copy("static", "dist/static"))
	must(os.RemoveAll("dist/static/scripts/test"))
}

func writePage(path string, page gomponents.Node) error {
	// create a writer
	builder := bytes.Buffer{}
	must(page.Render(&builder))
	pathOfFile := "dist" + path + ".html"
	err := os.Mkdir(filepath.Dir(pathOfFile), 0755)
	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			panic(err)
		}
	}
	file, err := os.Create(pathOfFile)
	if err != nil {
		return err
	}
	_, err = file.Write(builder.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
