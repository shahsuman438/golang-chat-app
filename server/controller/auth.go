package controller

import (
	"encoding/json"
	"net/http"
	"server/db"
	"server/services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var data db.User
	err := json.NewDecoder(r.Body).Decode(&data)
	var loginData = data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result, err := services.GetUser(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if loginData.Password != result.Password {
		http.Error(w, "Password Doesnot Matched", http.StatusUnauthorized)
		return
	} else {
		jsonResp, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResp)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var data db.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, error := services.CreateUser(&data)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
	}
	jsonResp, errs := json.Marshal(result)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusNotFound)
		return
	}
	w.Write(jsonResp)
}
