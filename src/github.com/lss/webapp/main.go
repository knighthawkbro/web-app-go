package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"web-app-go/src/github.com/lss/webapp/controller"
	"web-app-go/src/github.com/lss/webapp/middleware"
	"web-app-go/src/github.com/lss/webapp/model"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	templates := populateTemplates()
	//db := connectToDatabase()
	//defer db.Close()
	controller.Startup(templates)
	http.ListenAndServe(":8080", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

// Ignore this Using Microsoft SQL Server instead of PostgreSQL
func connectToDatabase() *sql.DB {
	db, err := sql.Open("mssql", "server=localhost;database=lss;user id=sa;password=<Removed SA password for reasons>;")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
