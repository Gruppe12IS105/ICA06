// Copyright 2017 Gruppe 12 IS-105. All rights reserved.
package main

import(
	"net/http"
	"fmt"
)

func main(){
	fs := http.FileServer(http.Dir("static")) //Setter verdien "fs" til å være en filserver som bruker mappen static
	http.Handle("/", fs) //Binder "fs" til "http://servername/"
	http.HandleFunc("/speech",inputhandler) //Binder "Http://servername/speech" til funksjonen inputhandler
	http.ListenAndServe(":8005", nil) //Starter serveren på port 8005
}

func inputhandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm() //Parser forms fra html siden. slik at den kan brukes senere
	fmt.Println(r.Form) //Printer index.html siden
	speechText := r.Form.Get("speechText") //Henter verdien "speechText" fra HTML formet
	text1 := speechText
	http.Redirect(w, r, "http://is105.gnorli.no:8080/speech?text=" + text1, 307) //Redirecter til espeak URLen sammen med teksten som ble skrevet inn tidligere
}



