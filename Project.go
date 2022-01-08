package main

import (
	"database/sql" // provides a generic interface around SQL databases

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql" //MySQL-Driver

	"encoding/json" //Package json implements encoding and decoding of JSON

	"github.com/gorilla/mux" //implements a request router and dispatcher for matching incoming requests to their respective handler

	"log"

	"strconv"

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

type Response_Message struct {
	Message string `json:"message"`
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

func GetUserInformationById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	ID,err := strconv.Atoi(params["id"])
	if err!=nil{
		log.Fatal(err)
	}
	Information, err := database.Query("select * from userinformation where id = ?;",ID)
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

func AddUserData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user_data UserData
	var response Response_Message
	err = json.Unmarshal(body,&user_data)
	if err!=nil{
		response.Message = "Failed to add user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	}
	_, err = database.Exec("insert into userinformation(id,name,dob,address,description) values(?,?,?,?,?);",user_data.Id,user_data.Name,user_data.Dob,user_data.Address,user_data.Description)
	if err != nil {
		response.Message = "Failed to add user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	} else{
		response.Message = "Successfully added user information"
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteAllUsersData(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	var response Response_Message
	_, err = database.Exec("delete from userinformation")
	if err != nil {
		response.Message = "Failed to delete user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	} else{
		response.Message = "Successfully deleted user data"
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteUserDataById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	w.Header().Set("content-type","application/json")
	var response Response_Message
	ID,err := strconv.Atoi(params["id"])
	if err!=nil{
		response.Message = "Failed to delete user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	}
	_, err = database.Exec("delete from userinformation where id = ?",ID)
	if err!= nil{
		response.Message = "Failed to delete user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	} else{
		response.Message = "Successfully deleted user data"
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateUserInformation(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	w.Header().Set("content-type","application/json")
	var response Response_Message
	ID,err := strconv.Atoi(params["id"])
	if err!=nil{
		response.Message = "Failed to update user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Message = "Failed to update user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	}
	var user_data UserData
	err = json.Unmarshal(body,&user_data)
	if err!=nil{
		response.Message = "Failed to update user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	}
	_, err = database.Exec("update userinformation set id = ?,name = ?,dob = ?,address = ?,description = ? where id = ?",user_data.Id,user_data.Name,user_data.Dob,user_data.Address,user_data.Description,ID)
	if err!= nil{
		response.Message = "Failed to update user data"
		json.NewEncoder(w).Encode(response)
		log.Fatal(err)
	} else{
		response.Message = "Successfully updated user data"
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	database, err = sql.Open("mysql","root:Reddy@123@tcp(127.0.0.1:3306)/admybrand")
	if err!= nil{
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/get",GetUserInformation).Methods("GET")
	router.HandleFunc("/get/{id}",GetUserInformationById).Methods("GET")
	router.HandleFunc("/create",AddUserData).Methods("POST")
	router.HandleFunc("/delete",DeleteAllUsersData).Methods("DELETE")
	router.HandleFunc("/delete/{id}",DeleteUserDataById).Methods("DELETE")
	router.HandleFunc("/update/{id}",UpdateUserInformation).Methods("PUT")
	http.ListenAndServe(":8000",router)
}