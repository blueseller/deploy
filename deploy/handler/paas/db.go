package paas

import (
	"fmt"
	"net/http"
)

func MongoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mongo Handler")
}

func PgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("pg Handler")
}

func MysqlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("mysql Handler")
}
