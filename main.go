// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START cloudrun_helloworld_service]
// [START run_helloworld_service]

// Sample run-helloworld is a minimal Cloud Run service.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type State struct {
	code     string
	secret   string
	clientID string
	client   *http.Client
}

func main() {
	s := &State{}
	s.secret = os.Getenv("SECRET")
	s.secret = os.Getenv("CLIENTID")
	log.Print("starting server...")
	http.HandleFunc("/", s.handler)
	http.HandleFunc("/connected", s.connected)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type AccessTokenRequest struct {
	Client_id     string `json"client_id"`
	Client_secret string `json"client_secret"`
	Code          string `json"code"`
	Redirect_uri  string `json"redirect_uri"`
	GrandType     string `json"grant_type"`
}

type AccessTokenResponse struct {
	Access_token  string "access_token"
	Token_type    string "token_type"
	Expires_in    int    "expires_in"
	Refresh_token string "refresh_token"
}

func (a *State) connected(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}

func (a *State) handler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.String())

	c := r.URL.Query().Get("code")

	if c != "" {
		a.code = c
	}
	log.Print(a.code)

	// body := strings.NewReader("client_id=" + "iQKsY6VwmrZPnDmJZtnXEZZBT0OTXaGp_GOGZ9CYexg" + "&client_secret=" + "YwvBTlkFU37_2nYz8Y7TD44cmPP7B3w9iWcqUJkHPyM" + "&code=" + "LMVCHYzzf0WF4rV3R_hBgU7QV50-cykS43bRU27hnmI" + "&grant_type=authorization_code&redirect_uri=https://www.briggsquote.com/")
	body := strings.NewReader("client_id=" + a.clientID + "&client_secret=" + a.secret + "&code=" + a.code + "&grant_type=authorization_code&redirect_uri=https://www.briggsquote.com/")

	// body := strings.NewReader("client_id=" + a.clientID + "&client_secret=" + a.secret + "&code=" + "T7eikSRruoxe_mGOZ8KdjTnXwcKTU9rOU7TkflGR4cY" + "&redirect_uri=https://www.briggsquote.com/")
	req, err := http.NewRequest("POST", "https://app.companycam.com/oauth/token", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	// reqBody, err := json.Marshal(map[string]string{
	// 	"grant_type":    "authorization_code",
	// 	"client_id":     "iQKsY6VwmrZPnDmJZtnXEZZBT0OTXaGp_GOGZ9CYexg",
	// 	"client_secret": "YwvBTlkFU37_2nYz8Y7TD44cmPP7B3w9iWcqUJkHPyM",
	// 	"code":          "T7eikSRruoxe_mGOZ8KdjTnXwcKTU9rOU7TkflGR4cY",
	// 	"redirect_uri":  "https://www.briggsquote.com/",
	// })
	// if err != nil {
	// 	log.Fatal("marshaling request for retrieving an access token: %w", err)
	// }

	// // request, err := http.NewRequest(http.MethodPost, "https://app.companycam.com/oauth/token?grant_type=authorization_code&client_id="+a.clientID+"&client_secret="+a.secret+"&code="+"T7eikSRruoxe_mGOZ8KdjTnXwcKTU9rOU7TkflGR4cY"+"&redirect_uri=https://www.briggsquote.com/", nil)
	// // if err != nil {
	// // 	log.Fatal("creating request for auth tokwn: %w", err)
	// // }

	// request, err := http.NewRequest(http.MethodPost, "https://app.companycam.com/oauth/token/", bytes.NewBuffer(reqBody))
	// if err != nil {
	// 	log.Fatal("creating request for auth tokwn: %w", err)
	// }

	fmt.Println(req)
	log.Printf(resp.Status)

	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal("reading response body from requesting the auth token: %w", err)
	// }
	// log.Println(string(body))

	// if res.StatusCode != http.StatusOK {
	// 	log.Fatalf("got HTTP code %d while retrieving access token", res.StatusCode)

	// }

	// atr := &AccessTokenResponse{}
	// if err := json.Unmarshal(body, atr); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(atr.Access_token)

	fmt.Fprintf(w, "200")
}

// [END run_helloworld_service]
// [END cloudrun_helloworld_service]
