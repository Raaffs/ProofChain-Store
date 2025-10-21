package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJson(w http.ResponseWriter,status int, response any){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	if err:=json.NewEncoder(w).Encode(response);err!=nil{
		http.Error(w, fmt.Sprintf("Error encoding JSON: %s", err), http.StatusInternalServerError)
		return
	}
}

func ReadJson(w http.ResponseWriter, r *http.Request,data any)(error){
	return json.NewDecoder(r.Body).Decode(data)
}