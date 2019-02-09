package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Joke struct {
	Fortune string
}

func main() {
	resp, err := http.Get("http://yerkee.com/api/fortune")
	if err != nil {
		log.Fatal("you suck")
	}
	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	s := Joke{}
	json.Unmarshal(robots, &s)
	fmt.Println(s.Fortune)
}
