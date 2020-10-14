package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Response struct will pick out the elements of JSON I want to display
type Response struct {
	Components []struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	Status struct {
		Description string `json:"description"`
	}
}

func main() {
	response, err := http.Get("https://kctbh9vrtdwd.statuspage.io/api/v2/summary.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject)

}
