package controller

import (
	"client/interfaces"
	"client/service"
	"encoding/json"
	"io/ioutil"
	"log"
)

func Login(loginData interfaces.Login) {
	json_data, err := json.Marshal(loginData)
	if err != nil {
		log.Printf("somthing went wrong\n", err)
	}
	result, error := service.Login(json_data)
	if error != nil || result.StatusCode == 401 {
		log.Printf("Unable to Login please check email and Password")
	} else {
		user := interfaces.Register{}
		json.NewDecoder(result.Body).Decode(&user)
		Chat(user)
	}

}

func Register(registerData interfaces.Register) {
	json_data, err := json.Marshal(registerData)
	if err != nil {
		log.Printf("somthing went wrong\n", err)
		return
	}
	result, error := service.Register(json_data)
	if error != nil || result.StatusCode == 500 {
		resBody, errors := ioutil.ReadAll(result.Body)
		if errors != nil {
			log.Printf("Unable to Register", errors)
		}
		log.Printf("Unable to Register", string(resBody), error)
		result.Body.Close()
		return
	}
	log.Printf("Hi! %s successfully Registerd\n", registerData.UserName)
	log.Printf("Please use cmd ./client login <email> <password> to login\n")
	return
}
