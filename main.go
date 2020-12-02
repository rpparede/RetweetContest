package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

func main() {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}
	f, err := os.Open("keys.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	fmt.Printf("%+v\n", keys)

	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(keys.Key, keys.Secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset-UTF-8")
	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var token oauth2.Token
	dec = json.NewDecoder(res.Body)
	err = dec.Decode(&token)
	if err != nil {
		panic(err)
	}

	var conf oauth2.Config
	tclient := conf.Client(context.Background(), &token)
	res2, err := tclient.Get("https://api.twitter.com/1.1/statuses/retweets/1333940412709543936.json?count=100")
	if err != nil {
		panic(err)
	}
	defer res2.Body.Close()
	//io.Copy(os.Stdout, res2.Body)

	var retweets []Retweet
	dec = json.NewDecoder(res2.Body)
	err = dec.Decode(&retweets)
	if err != nil {
		panic(err)
	}
	var usernames []string
	for _, retweet := range retweets {
		usernames = append(usernames, retweet.User.ScreenName)
	}
	fmt.Println(usernames)
	fmt.Println(len(usernames))

}

type Retweet struct {
	User struct {
		ScreenName string `json:"screen_name"`
	} `json:"user"`
}
