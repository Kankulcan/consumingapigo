package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//Struct to get the json data.
type Advice struct {
	Advices Message `json:"slip"`
}

// Struct to get the json data.
type Message struct {
	Id     int    `json:"id"`
	Advice string `json:"advice"`
}

const apiURL = "https://api.adviceslip.com/advice"

func main() {

	repeatPrintMessage(apiURL, 5, 20)

}

//getMessage - Get message from API - return Message
func getMessage(apiURL string) Message {

	var jsonRequest Advice
	var message Message

	apiResponse := getAPIResponse(apiURL)
	apiResponseBody := getAPIResponseBody(&apiResponse)

	if json.Valid(apiResponseBody) {

		err := json.Unmarshal(apiResponseBody, &jsonRequest)
		if err != nil {
			return errorMessage("can't read from Json File")
		}

		message = jsonRequest.Advices
		return message

	} else {

		return errorMessage("this is not a valid Json File.")
	}

}

//getAPIResponse - Get response from API url.
func getAPIResponse(apiURL string) http.Response {
	apiResponse, err := http.Get(apiURL)
	if err != nil {
		log.Panic(err)
	}
	return *apiResponse

}

//getAPIResponseBody - get body from response - return in []byte
func getAPIResponseBody(apiResponse *http.Response) []byte {
	apiResponseBody, err := ioutil.ReadAll(apiResponse.Body)
	if err != nil {
		log.Println(err)
	}

	apiResponse.Body.Close()
	return apiResponseBody

}

func errorMessage(err string) Message {
	var message Message
	message = Message{Id: -1, Advice: err}

	return message
}

//printMessage - Print the API Message
func printMessage(message Message) {

	if message.Id > 0 {

		fmt.Println("====== Advice =======")
		fmt.Printf("%v  ", message.Id)
		fmt.Println(message.Advice)
		fmt.Println("======== End =========")

	} else {

		fmt.Println("====== warning error! ========")
		fmt.Printf("%v  ", message.Id)
		fmt.Println(message.Advice)
		fmt.Println("======== End of error =========")
	}

}

func repeatPrintMessage(apiURL string, timeSecond int, repeatCount int) {

	for i := 0; i < repeatCount; i++ {

		var apiMessage Message
		apiMessage = getMessage(apiURL)
		printMessage(apiMessage)
		time.Sleep(time.Second * time.Duration(timeSecond))
	}
}
