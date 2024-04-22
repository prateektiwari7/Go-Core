package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*

{
  "userId": 1,
  "id": 2,
  "title": "quis ut nam facilis et officia qui",
  "completed": false
}

*/

type Post struct {
	Userid    int    `json : "userId"`
	Id        int    `json : "id"`
	Title     string `json : "title"`
	Completed bool   `json : "completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

		data := Post{}

		Decoder := json.NewDecoder(response.Body)
		// helps to disallowed if any struct body mismatched
		Decoder.DisallowUnknownFields()

		if err := Decoder.Decode(&data); err != nil {
			log.Fatal("Decode err", err)
		}

		fmt.Printf("Data is %+v", data)

	}

}
