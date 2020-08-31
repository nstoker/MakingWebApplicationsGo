package main

import (
	"database/sql"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nstoker/MakingWebApplicationsGo/src/controller"
	"github.com/nstoker/MakingWebApplicationsGo/src/middleware"
	"github.com/nstoker/MakingWebApplicationsGo/src/model"

	_ "github.com/lib/pq"
)

func main() {
	dsn, port, salt := readEnvironmentVariables()
	model.SetPasswordSalt(salt)
	templates := populateTemplates()
	db := connectToDatabase(dsn)
	defer db.Close()

	controller.Startup(templates)

	log.Printf("Starting up on :%s", port)
	http.ListenAndServe(
		":"+port,
		&middleware.TimeoutMiddleware{
			new(middleware.GzipMiddleware),
		},
	)
}

func readEnvironmentVariables() (string, string, string) {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("$PORT missing")
	}
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalln("$DATABASE_URL missing")
	}
	salt := os.Getenv("PASSWORD_SALT")
	if salt == "" {
		log.Fatalln("$PASSWORD_SALT missing")
	}
	if port == "" || dsn == "" {
		log.Panicf("Check $PORT '%s' and $DATABASE_URL '%s'", port, dsn)
	}

	return dsn, port, salt
}

func connectToDatabase(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
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
