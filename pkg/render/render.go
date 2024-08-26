package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tcnam/NewsFeed/pkg/config"
)

var app *config.AppConfig

func NewConfig(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc := app.TemplateCache

	t, ok := tc[tmpl]
	if !ok {
		log.Printf("Couldn't find %s in templateCache", tmpl)
	}
	// write template to buffer
	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		log.Println("error writing template to buffer", err)
		return
	}
	_, err = buf.WriteTo(w)
	// var parsedTemplate, _ = template.ParseFiles("./template/"+tmpl, "./template/base.html")
	// err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing the template", err)
		return
	}
}

func InitTemplateCache() (map[string]*template.Template, error) {
	var myCache map[string]*template.Template = make(map[string]*template.Template)

	var pages, err = filepath.Glob(filepath.Join(".", "template", "*.html"))
	// log.Fatalf("%v", pages)
	if err != nil {
		log.Printf("Error while reading .html files: %s", err)
		return myCache, err
	}
	for _, page := range pages {
		var name string = filepath.Base(page)
		var ts, err_ts = template.New(name).ParseFiles(page)
		if err_ts != nil {
			log.Printf("Error while reading page.html files: %s", err_ts)
			return myCache, err_ts
		}
		// log.Printf("Index: %d, full page: %s, page: %s", index, page, name)

		// check if base template is existed
		var matches, err_match = filepath.Glob(filepath.Join(".", "template", "base.html"))
		if err_match != nil {
			log.Printf("Error while reading base.html files: %s", err_match)
			return myCache, err_match
		}
		if len(matches) > 0 {
			ts, err_ts = ts.ParseGlob(filepath.Join(".", "template", "base.html"))
			if err_ts != nil {
				return myCache, err_ts
			}
		}
		myCache[name] = ts
	}
	return myCache, nil

}
