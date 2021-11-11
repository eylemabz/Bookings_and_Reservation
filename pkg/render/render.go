package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, html string) {

	_, err := RenderTemplateTest(w)

	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + html)

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

var functions = template.FuncMap{}

func RenderTemplateTest(w http.ResponseWriter, hmtl string) (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./teplates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("page is current ", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
		}
		if err != nil {
			return myCache, err
		}
		myCache[name] = ts
	}
	return myCache, nil
}
