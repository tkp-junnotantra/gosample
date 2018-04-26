package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Pyramid struct {
	MaxNum int `json:"max_num"`
	Height int `json:"height"`
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	jsonData, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	var p Pyramid
	err = json.Unmarshal(jsonData, &p)
	checkErr(err)

	fmt.Printf("%+v", p)

	newP := Pyramid{
		MaxNum: p.MaxNum + 2,
		Height: p.Height + 3,
	}

	data, err := json.Marshal(newP)
	checkErr(err)

	w.Write(data)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
