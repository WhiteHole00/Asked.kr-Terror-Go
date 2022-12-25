package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var id string
	var content string

	var count int

	fmt.Print("ENTER ID >> ")
	fmt.Scanln(&id)

	fmt.Print("ENTER CONTENT >> ")
	fmt.Scanln(&content)

	fmt.Print("ENTER COUNT >> ")
	fmt.Scanln(&count)

	for i := 0; i < count; i++ {
		var BaseURL string = ("https://asked.kr/" + id)
		resp_, err := http.Get(BaseURL)

		if err != nil {
			log.Fatal(err)
		}

		defer resp_.Body.Close()

		data := url.Values{
			"id":           {id},
			"content":      {content},
			"makarong_bat": {"-1"},
			"show_user":    {"0"}}

		resp, e := http.PostForm("https://asked.kr/query.php?query=100", data)

		if e != nil {
			log.Fatal(e)
		}

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

	}

	fmt.Println("질문 완료! | " + id + " | " + content)
}
