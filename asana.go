package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func createReq(url string) *http.Request {
	accessToken := os.Getenv("ACCESS_TOKEN")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error init request")
	}
	auth := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Set("Authorization", auth)

	return req
}

func request(req *http.Request) string {
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	respStr := string(byteArray)
	return respStr
}

func main() {
	envLoad()
	url := "https://app.asana.com/api/1.0/users/me"
	req := createReq(url)
	resp := request(req)

	fmt.Println(resp)
}
