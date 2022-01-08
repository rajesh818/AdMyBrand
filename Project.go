package main

import (
	"database/sql" // provides a generic interface around SQL databases

	_ "github.com/go-sql-driver/mysql" //MySQL-Driver

	"encoding/json" //Package json implements encoding and decoding of JSON

	"github.com/gorilla/mux" //implements a request router and dispatcher for matching incoming requests to their respective handler

	"log"

	"net/http"
)
var database *sql.DB
var err error

type UserData struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Dob string `json:"dob"`
	Address string `json:"address"`
	Description string `json:"description"`
	CreatedAt string `json:"createdat"`
}

func GetUserInformation(w http.ResponseWriter, r *http.Request){

	Information, err := database.Query("select * from userinformation;")
	if err != nil {
		log.Fatal(err)
	}
	defer Information.Close()
	var response []UserData
	for Information.Next(){
		var data UserData
		err := Information.Scan(&data.Id,&data.Name,&data.Dob,&data.Address,&data.Description,&data.CreatedAt)
		if err!= nil{
			log.Fatal(err)
		}
		response = append(response, data)
	}
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	database, err = sql.Open("mysql","root:Reddy@123@tcp(127.0.0.1:3306)/admybrand")
	if err!= nil{
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/get",GetUserInformation).Methods("GET")
	http.ListenAndServe(":8000",router)
}