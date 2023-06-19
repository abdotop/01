package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	//Id           int      `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
	//Members      []string `json:"members"`
	//Date         int      `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
	Locations    string `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations    string `json:"relations"`
}

func main() {
	api, err := http.Get("https://groupietrackers.herokuapp.com/api/artist")
	if err != nil {
		//panic(err)
		return
	}
	defer api.Body.Close()
	apiBody, err := ioutil.ReadAll(api.Body)
	if err != nil {
		//panic(err)
		return
	}
	var ApiLink Api
	err = json.Unmarshal(apiBody, &ApiLink)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(ApiLink)
}
