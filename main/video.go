package main

import (
	"encoding/json"
	"garagem/data"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "aplication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	paramJSON := r.URL.Query().Get("Id")
	foundedVideo := data.Populate(paramJSON)

	j, err := json.Marshal(foundedVideo)
	if err != nil {
		panic(err)
	}
	w.Write(j)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}

/*
func testehandle(w http.ResponseWriter, r *http.Request){
	//Implementar json.Marshal para passar uma lista
	var videos []Video
	   	videos[0] = video
	   	videos[1] = video2
	   	videos[2] = video3
	   	videos[3] = video4
	   	for vid := range videos {
	   		b, err := json.Marshal(vid)
	   		if err != nil {
	   			panic(err)
	   		}
	   		fmt.Println("Primerio Json: ", string(b))
	   	}
	l, err := json.Marshal(data.VideoCode)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Context-Type", "aplication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(l)
	//fmt.Println("Primerio Json: ", string(l))
	m, err := json.Marshal(data.VideoCodeInovation)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Context-Type", "aplication/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(m)
	fmt.Println("Primerio Json: ", string(m))
}
*/
