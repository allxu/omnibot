package main

import (
	"fmt"
	"log"
	"net/http"
)

func getURL(s string){
	
	"sina" :=
}

func redirect(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
	} else {
		key := keys[0]
		log.Println("Url Param is " + string(key))
	}

	http.Redirect(w, r, "http://www.sina.com", 301)
}

func main() {
	http.HandleFunc("/redirect/", redirect)
	fmt.Print("start......")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
