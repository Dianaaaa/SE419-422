package main

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Geturl(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	
	// para := user["shortlink"]
	// shortlink, ok := para.(string)
	// if !ok {
	// 	fmt.Println("conver error")
	// }
	// fmt.Println(shortlink)

	vars := r.URL.Query();
	var shortlink string
	shortlink = string(vars["shortlink"][0])
	fmt.Println(shortlink)

	var path string
	path = Get_path(shortlink)

	if path == "" {
		response := []byte("no such shortlink")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		response := []byte(path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}	
}

func Generate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &user)
	
	para := user["path"]
	path, ok := para.(string)
	if ok {
		fmt.Println("conver error")
	}
	fmt.Println(path)
	var ip string
	ip = r.RemoteAddr
	fmt.Println(ip)

	var shortlink string
	shortlink = Gen_short_url(ip, path)

	response := []byte(shortlink)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Heartbeat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	fmt.Fprint(w, "heart beat\n")
}