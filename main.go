/*
 * Copyright (c) Elliot Peele <elliot@bentlogic.net>
 *
 * This program is distributed under the terms of the MIT License as found
 * in a file called LICENSE. If it is not present, the license
 * is always available at http://www.opensource.org/licenses/mit-license.php.
 *
 * This program is distributed in the hope that it will be useful, but
 * without any warrenty; without even the implied warranty of merchantability
 * or fitness for a particular purpose. See the MIT License for full details.
 */

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
)

const API_MESSAGE = "https://api.pushover.net/1/messages.json"

type ApiResponse struct {
	Status  int      `json:"status"`
	Request string   `json:"request"`
	User    string   `json:"user,omitempty"`
	Errors  []string `json:"errors,omitempty"`
}

var (
	token   = flag.String("token", "", "Pushover API Token")
	user    = flag.String("user", "", "Pushover User Token")
	title   = flag.String("title", "", "Message Title")
	message = flag.String("message", "", "Messge Content")
)

func init() {
	flag.Parse()
}

func main() {
	if *token == "" || *user == "" || *title == "" || *message == "" {
		log.Fatal("Must set token, user, title, and message.")
	}

	values := url.Values{}
	values.Set("token", *token)
	values.Set("user", *user)
	values.Set("title", *title)
	values.Set("message", *message)

	resp, err := http.PostForm(API_MESSAGE, values)

	if err != nil {
		log.Fatalf("Error making request: %s", err)
	}

	var out ApiResponse
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&out); err != nil {
		log.Fatalf("Error parsing response: %s", err)
	}

	if out.Status == 1 {
		log.Print("sent successfully")
	} else {
		log.Fatalf("error sending message: %s", out.Errors[0])
	}
}
