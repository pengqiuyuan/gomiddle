package gomiddle

import (
	"net/http"
)

type SealEntity struct{
	ServerZoneId string
	GameId string
	ServerId string
	Guid string
	SagTime string
	SagStart string
	SagEnd string
}

func SealHandler() {
	http.HandleFunc("/fbserver/seal/getAllSealAccount", GetAllSealAccount)
	http.HandleFunc("/fbserver/seal/addSealAccount", AddSealAccount)
	http.HandleFunc("/fbserver/seal/updateSealAccount", UpdateSealAccount)
	http.HandleFunc("/fbserver/seal/delSealAccount", DelSealAccount)
}


func GetAllSealAccount(w http.ResponseWriter, r *http.Request){
	
}

func AddSealAccount(w http.ResponseWriter, r *http.Request){
	
}

func UpdateSealAccount(w http.ResponseWriter, r *http.Request){
	
}

func DelSealAccount(w http.ResponseWriter, r *http.Request){
	
}