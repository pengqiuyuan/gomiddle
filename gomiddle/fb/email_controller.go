package gomiddle

import (
	"net/http"
)

type Annex struct{
	ItemId string
	ItemNuM int
}

type EmailEntity struct{
	ServerZoneId string 
	GameId string
	ServerId string
	PlatForm string   
	Sender string 
	Title string  
	Contents string   
	Annexs []Annex
}

func EmailHandler() {
	http.HandleFunc("/fbserver/seal/getAllEmails", GetAllEmails)
	http.HandleFunc("/fbserver/seal/addEmail", AddEmail)
	http.HandleFunc("/fbserver/seal/updateEmail", UpdateEmail)
	http.HandleFunc("/fbserver/seal/delEmailById", DelEmailById)
	http.HandleFunc("/fbserver/seal/getEmailById", GetEmailById)
}


func GetAllEmails(w http.ResponseWriter, r *http.Request){
	
}

func AddEmail(w http.ResponseWriter, r *http.Request){
	
}

func UpdateEmail(w http.ResponseWriter, r *http.Request){
	
}

func DelEmailById(w http.ResponseWriter, r *http.Request){
	
}

func GetEmailById(w http.ResponseWriter, r *http.Request){
	
}