package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

type File struct {
	Id 			int
	Name 		string
	Location	string
	Metadata	string
}

func TestEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Test")
}

func main()  {
	db, err := sql.Open("mysql", "root:bitbucket@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
		log.Fatal(err)
	}

	fmt.Println("Started Server Testing")
	r := mux.NewRouter()
	r.HandleFunc("/test/", TestEndpoint)
	http.ListenAndServe(":8090", r)
	if err != nil {
		panic(fmt.Errorf("Unable to listen: %s", err))
	}

	defer db.Close()

}