package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)	
	if err := parsedTemplate.Execute(w, nil);err != nil {
		fmt.Println("error parsing template", err)	
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string){
	var tmpl *template.Template
	var err error

	// template이 캐시안에 있는지 확인
	if _, inMap:= tc[t]; !inMap {
		// 캐시를 만들어야함
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// 캐시가 존재함 -> 불러옴
		log.Println("using cached template")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl
	return nil
}