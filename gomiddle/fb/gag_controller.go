package gomiddle

import (
	"net/http"
)

type GagEntity struct{
	ServerZoneId string
	GameId string
	ServerId string
	Guid string
	GagTime string
	GagStart string
	GagEnd string
}

func GagHandler() {
	http.HandleFunc("/fbserver/gag/getAllGagAccount", GetAllGagAccount)
	http.HandleFunc("/fbserver/gag/addGagAccount", AddGagAccount)
	http.HandleFunc("/fbserver/gag/updateGagAccount", UpdateGagAccount)
	http.HandleFunc("/fbserver/gag/delGagAccountById", DelGagAccountById)
}


func GetAllGagAccount(w http.ResponseWriter, r *http.Request){
	
}

func AddGagAccount(w http.ResponseWriter, r *http.Request){
	
}

func UpdateGagAccount(w http.ResponseWriter, r *http.Request){
	
}

func DelGagAccountById(w http.ResponseWriter, r *http.Request){
	
}