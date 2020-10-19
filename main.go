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

func SetUpRoutes() {
	// Create the route handler listening on '/'
	http.HandleFunc("/", Home)
	fmt.Println("Starting server on port 8080")

	// Start the sever
	http.ListenAndServe(":8080", nil)
}

func main() {
	SetUpRoutes()
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Calling GitHub API for Data")
	response, err := http.Get("https://kctbh9vrtdwd.statuspage.io/api/v2/summary.json")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// Pretty-print the json
	responseBodyBytes := new(bytes.Buffer)
	json.NewEncoder(responseBodyBytes).Encode(responseObject)
	byteArray := []byte(responseBodyBytes.Bytes())
	// byteBuffer := &bytes.Buffer{}
	// if err := json.Indent(byteBuffer, byteArray, "", "  "); err != nil {
	// 	panic(err)
	// }

	// Use for console printing / debugging
	// fmt.Println(byteBuffer.String())

	// Write the response to the byte array - Sprintf formats and returns a string without printing it anywhere
	// w.Write([]byte(fmt.Sprintf(byteBuffer.String())))
	w.Write(byteArray)
}
