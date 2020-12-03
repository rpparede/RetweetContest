package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"./twitter"
)

func main() {
	key, secret, err := getKeys()
	if err != nil {
		panic(err)
	}
	client, err := twitter.New(key, secret)
	if err != nil {
		panic(err)
	}

	usernames, err := client.Retweeters("1333932197506650112")
	if err != nil {
		panic(err)
	}
	//tclient, err := tweeterClient(keys.Key, keys.Secret)
	//usernames, err := retweeters(tclient, "1333932197506650112")

	fmt.Println(usernames)
	fmt.Println(len(usernames))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(len(usernames))
	fmt.Printf("Winner is: %s", usernames[n])
}

func getKeys() (key string, secret string, err error) {
	var keys struct {
		Key    string `json:"consumer_key"`
		Secret string `json:"consumer_secret"`
	}

	f, err := os.Open("keys.json")
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	dec.Decode(&keys)
	return keys.Key, keys.Secret, nil
}
