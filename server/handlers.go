package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var users = make(map[string]*User)
var pairs = make(map[*User]*User)
var lobby = make([]*User, 0)

func Receive(w http.ResponseWriter, r *http.Request) {

	wholeurl := r.URL.String()
	body := r.URL.Query()["Body"]
	phone := r.URL.Query()["From"]

	fmt.Printf("wholeurl:\n%s\n\nPhone: %s\nBody: %s,\n\n", wholeurl, phone, body)
}

func sendSMS(phonenumber, message string) {

	apiusr := os.Getenv("TWILIO_APIUSR")
	apikey := os.Getenv("TWILIO_APIKEY")

	u := "https://api.twilio.com/2010-04-01/Accounts/AC7dbbd979132aeb252095fa79059a5de4/Messages.json"

	hc := http.Client{}
	form := url.Values{}
	form.Add("To", phonenumber)
	form.Add("From", "+13208398785")
	form.Add("Body", message)

	req, err := http.NewRequest("POST", u, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(apiusr, apikey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//fmt.Printf("the request was: \n%v\n\n",req)

	resp, err := hc.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 201 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(body)
	}
}
