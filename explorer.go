package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	type Course struct {
		Category string `json:"category"`
		Title    string `json:"title_en"`
		Price    string `json:"price"`
	}
	type Response struct {
		Courses []Course `json:"courses"`
	}
	today := time.Now()
	matchDate := today.Format("2006/01/02")

	url := fmt.Sprintf("https://www.sodexo.fi/ruokalistat/output/daily_json/31/%s/en", matchDate)
	resp, _ := http.Get(url)

	var response Response
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	for _, course := range response.Courses {
		fmt.Printf("%s: %s %s€\n", course.Category, course.Title, course.Price)
	}
}
