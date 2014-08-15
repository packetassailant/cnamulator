package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Opencnams struct {
	Results []Opencnam `json:"results"`
}

type Opencnam struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}

func cnamReq(phonenum, sid, token string) (Opencnam, error) {
	oc := Opencnam{}
	resp, err := http.Get("https://api.opencnam.com/v2/phone/+" + phonenum + "?format=json&account_sid=" + sid + "&auth_token=" + token)
	if err != nil {
		return oc, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if len(data) == 0 {
		oc.Number = phonenum
		oc.Name = "undefined"
		return oc, err
	}
	if err != nil {
		return oc, err
	}
	err = json.Unmarshal(data, &oc)
	if err != nil {
		return oc, err
	}
	oc.Number = strings.TrimLeft(oc.Number, "+")
	return oc, nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	phone := flag.String("phone", "", "a single phone number")
	phoneFile := flag.String("file", "", "a list of phone numbers")
	sid := flag.String("sid", "", "the opencnam api sid")
	token := flag.String("token", "", "the opencnam api auth token")
	flag.Parse()

	requests := []string{}
	cnams := &Opencnams{}

	if (*phone != "") && (*phoneFile != "") {
		log.Fatal("-phone and -file are mutually exclusive")
	}
	if *sid == "" {
		log.Fatal("an opencnam sid is required")
	}
	if *token == "" {
		log.Fatal("an opencnam auth token is required")
	}
	if *phone != "" {
		requests = append(requests, *phone)
	}
	if *phoneFile != "" {
		lines, err := readLines(*phoneFile)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		requests = append(requests, lines...)
	}
	for _, r := range requests {
		result, err := cnamReq(r, *sid, *token)
		if err != nil {
			log.Fatal(err)
		}
		cnams.Results = append(cnams.Results, result)
	}
	j, err := json.MarshalIndent(cnams, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
