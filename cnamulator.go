package main

import "flag"
import "fmt"
import "strings"
import "net/http"
import "io/ioutil"
import "bufio"
import "os"
import "log"
import "encoding/json"

type Opencnam struct {
	Number  string `json:"number"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Name    string `json:"name"`
	Uri     string `json:"uri"`
}

func cnamReq(phonenum, sid, token string) {
	resp, err := http.Get("https://api.opencnam.com/v2/phone/+" + phonenum + "?format=json&account_sid=" + sid + "&auth_token=" + token)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		defer resp.Body.Close()
		src_json, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
		}
		oc := Opencnam{}
		e := json.Unmarshal(src_json, &oc)
		if e != nil {
			panic(e)
		}
		if src_json != nil {
			mapJsonOne := map[string]interface{}{"Phone": strings.TrimLeft(oc.Number, "+"), "Name": oc.Name}
			mapJsonTwo, _ := json.Marshal(mapJsonOne)
			fmt.Println(string(mapJsonTwo))
		}
	}
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
	phonePtr := flag.String("phone", "", "a single phone number")
	phoneFilePtr := flag.String("file", "", "a list of phone numbers")
	sidPtr := flag.String("sid", "", "the opencnam api sid")
	tokenPtr := flag.String("token", "", "the opencnam api auth token")
	flag.Parse()
	phone := *phonePtr
	phonefile := *phoneFilePtr
	sid := *sidPtr
	token := *tokenPtr

	if (*phonePtr != "") && (*phoneFilePtr != "") {
		fmt.Println("-phone and -file are mutually exclusive")
		os.Exit(1)
	}
	if *sidPtr == "" {
		fmt.Println("an opencnam sid is required")
		os.Exit(1)
	}
	if *tokenPtr == "" {
		fmt.Println("an opencnam auth token is required")
		os.Exit(1)
	}
	if *phonePtr != "" {
		cnamReq(phone, sid, token)
	}
	if *phoneFilePtr != "" {
		lines, err := readLines(phonefile)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		} else {
			for _, line := range lines {
				cnamReq(line, sid, token)
			}
		}
	}
}
