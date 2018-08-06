package main

//import "log"

import "io"
//import "io/ioutil"
//import "strconv"
import "net/http"
//import "encoding/json"

func httpHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello world!")
}

func main(){
	dbg(true, "Starting \n")
	
	readConfig("tootd.conf")

	http.HandleFunc("/", httpHandler)
	http.ListenAndServe("localhost:" + config["Port"], nil)
	
	dbg(true, "Finished \n")
}
