package main

import (
	"bytes"
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
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
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

	// Pretty-print the json
	responseBodyBytes := new(bytes.Buffer)
	json.NewEncoder(responseBodyBytes).Encode(responseObject)
	byteArray := []byte(responseBodyBytes.Bytes())
	byteBuffer := &bytes.Buffer{}
	if err := json.Indent(byteBuffer, byteArray, "", "  "); err != nil {
		panic(err)
	}

	fmt.Println(byteBuffer.String())

}
