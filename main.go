package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Data struct {
	FirstName     string
	LastName      string
	Email         string
	Age           int
	MonthlySalary []MonthlySalary
}

type MonthlySalary struct {
	Basic int `json:"basic"`
	HRA   int `json:"hra"`
	TA    int `json:"ta"`
}

func main() {
	for {
		response, err := http.Get("http://localhost:90/allusers")
		if err != nil {
			log.Println("ERROR! Failed to retrieve all users.")
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println("ERROR! Failed to close the body.")
			}
		}(response.Body)

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println("ERROR! Failed to read response of the body.")
		}

		var jsonData []Data
		err = json.Unmarshal(body, &jsonData)
		if err != nil {
			log.Println("ERROR! Failed to unmarshal the response body.")
		}

		output, err := json.Marshal(jsonData)
		if err != nil {
			log.Println("ERROR! Failed to marshal the data.")
		}

		err = ioutil.WriteFile("output.json", output, 0644)
		if err != nil {
			log.Println("ERROR! Failed to write to output file.")
		}

		log.Println("Successfully updated output file.")

		time.Sleep(time.Minute)
		log.Println("One minute cycle has finished! Updating output file.")
	}
}
