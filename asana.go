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

func get(url string) string {
	accessToken := os.Getenv("ACCESS_TOKEN")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error init request")
	}
	auth := fmt.Sprintf("Bearer %s", accessToken)
	req.Header.Set("Authorization", auth)
	resp := request(req)

	return resp
}

func request(req *http.Request) string {
	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	respStr := string(byteArray)
	return respStr
}

// Debug
func main() {
	envLoad()
	var url, resp string
	url = "https://app.asana.com/api/1.0/users/me"
	resp = get(url)
	fmt.Println(resp)

	url = "https://app.asana.com/api/1.0/tasks?project=616529690850115&opt_fields=name,completed,due_on"
	resp = get(url)
	fmt.Println(resp)

	// myurl := "https://app.asana.com/api/1.0/users/me"
	// req := createReq(myurl)
	// resp := request(req)
	// fmt.Println(resp)

	// myurl := "https://app.asana.com/api/1.0/tasks?limit=10"
	// data := url.Values{}
	// data.Set("project", "14641")
	// r, _ := http.NewRequest("POST", myurl, bytes.NewBufferString(data.Encode()))

	// accessToken := os.Getenv("ACCESS_TOKEN")
	// auth := fmt.Sprintf("Bearer %s", accessToken)
	// r.Header.Set("Authorization", auth)
	// resp := request(r)

}
