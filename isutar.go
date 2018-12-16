package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	_ "net/http/pprof"
)

var (
	baseUrl *url.URL
	db      *sql.DB
	re      *render.Render
)

func initializeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := db.Exec("TRUNCATE star")
	panicIf(err)
	re.JSON(w, http.StatusOK, map[string]string{"result": "ok"})
}

func starsPostHandler(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")

	var tmp int
	if err := db.QueryRow(`SELECT id FROM isuda.entry WHERE keyword = ?`, keyword).Scan(&tmp); err == sql.ErrNoRows {
		notFound(w)
		return
	}

	user := r.FormValue("user")
	_, err := db.Exec(`INSERT INTO star (keyword, user_name, created_at) VALUES (?, ?, NOW())`, keyword, user)
	panicIf(err)

	re.JSON(w, http.StatusOK, map[string]string{"result": "ok"})
}

func main() {

	go func() {
		log.Println(http.ListenAndServe(":6061", nil))
	}()

	host := os.Getenv("ISUTAR_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	portstr := os.Getenv("ISUTAR_DB_PORT")
	if portstr == "" {
		portstr = "3306"
	}
	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable ISUTAR_DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("ISUTAR_DB_USER")
	if user == "" {
		user = "root"
	}
	password := os.Getenv("ISUTAR_DB_PASSWORD")
	dbname := os.Getenv("ISUTAR_DB_NAME")
	if dbname == "" {
		dbname = "isutar"
	}

	db, err = sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=true",
		user, password, host, port, dbname,
	))
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	db.Exec("SET SESSION sql_mode='TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY'")
	db.Exec("SET NAMES utf8mb4")

	re = render.New(render.Options{Directory: "dummy"})

	r := mux.NewRouter()
	r.HandleFunc("/initialize", myHandler(initializeHandler))
	s := r.PathPrefix("/stars").Subrouter()
	// s.Methods("GET").HandlerFunc(myHandler(starsHandler))
	s.Methods("POST").HandlerFunc(myHandler(starsPostHandler))

	// log.Fatal(http.ListenAndServe(":5001", handlers.LoggingHandler(os.Stdout, r)))
	log.Fatal(http.ListenAndServe(":5001", r))
}
