package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Pyramid struct {
	MaxNum int `json:"max_num"`
	Height int `json:"height"`
}

var c *http.Client

func main() {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	c = &client
	post()
}

func post() {
	p := Pyramid{
		MaxNum: 5,
		Height: 5,
	}
	jsonData, err := json.Marshal(p)
	checkErr(err)

	req, err := http.NewRequest("POST", "http://localhost:9000/api/post", bytes.NewBuffer(jsonData))
	checkErr(err)

	res, err := c.Do(req)
	checkErr(err)
	defer res.Body.Close()

	respData, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	var pyr Pyramid
	err = json.Unmarshal(respData, &pyr)
	checkErr(err)

	index := 0
	for i := 0; i < pyr.Height; i++ {
		for j := 0; j <= i; j++ {
			index++
			if index > pyr.MaxNum {
				index = 1
			}
			print(index)
		}
		println()
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
