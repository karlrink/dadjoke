package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

func getJoke() (string, error) {
	url := "https://icanhazdadjoke.com/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var jokeResp JokeResponse
	err = json.Unmarshal(body, &jokeResp)
	if err != nil {
		return "", err
	}

	return jokeResp.Joke, nil
}

func main() {
	joke, err := getJoke()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(joke)
}

