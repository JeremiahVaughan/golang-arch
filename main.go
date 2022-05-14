package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string `json:"first"`
}

func main() {
	//p1 := person{
	//	First: "Jenny",
	//}
	//
	//p2 := person{
	//	First: "James",
	//}
	//
	//xp := []person{p1, p2}
	//
	//bs, err := json.Marshal(xp)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println("PRINT JSON", string(bs))
	//
	//xp2 := []person{}
	//err = json.Unmarshal(bs, &xp2)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println("Back into Go data structure", xp2)

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic()
	}

}

func foo(writer http.ResponseWriter, request *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	err := json.NewEncoder(writer).Encode(p1)
	if err != nil {
		log.Println("Encode bad data", err)
		return
	}
}

func bar(writer http.ResponseWriter, request *http.Request) {
	var p1 person
	err := json.NewDecoder(request.Body).Decode(&p1)
	if err != nil {
		log.Println("Decode bad data", err)
		return
	}
	log.Println("Person: ", p1)
}
