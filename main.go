package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/rpparede/RetweetContest/twitter"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {
		//panic(err)
		fmt.Println("No env file found")
	}

}
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc

	myRouter.HandleFunc("/twitterContest/{id}", returnSingleWinner)
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func returnSingleWinner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tweetID := vars["id"]

	winner, err := getRetweetContestWinner(tweetID)
	if err != nil {
		winner = ""
	}

	var retweets twitter.Retweet
	retweets.User.ScreenName = winner
	jsonRes, err := json.Marshal(retweets)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func main() {
	fmt.Println("RT Contest API")
	handleRequests()
}

func getRetweetContestWinner(tweetID string) (username string, err error) {
	key, secret, err := getKeys()
	if err != nil {
		return "", err
	}
	client, err := twitter.New(key, secret)
	if err != nil {
		return "", err
	}

	usernames, err := client.Retweeters(tweetID) //"1333932197506650112")
	if err != nil || usernames == nil {
		return "", err
	}

	fmt.Println(usernames)
	fmt.Println(len(usernames))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(len(usernames))
	fmt.Printf("Winner is: %s", usernames[n])
	return usernames[n], nil
}
func getKeys() (key string, secret string, err error) {
	return os.Getenv("KEY"), os.Getenv("SECRET"), nil
}
