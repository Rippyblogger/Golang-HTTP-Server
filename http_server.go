package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func joke(w http.ResponseWriter, r *http.Request) {

	jokes := []string{"What do mother-in-laws and rain clouds have in common?  When they clear away, it’s a beautiful day", "Who knows best what people are missing? The thief", "What does the hammer say to the thumb? Nice to meet you again.", "Which birds do not lay eggs?  The males"}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", jokes[rand.Intn(len(jokes))])

}

func madlib(w http.ResponseWriter, r *http.Request) {

	answers := make(map[string][]string)

	answers["name"] = []string{"Tom", "Dick", "Harry", "Sylvia", "Maddie", "Lizzy"}
	answers["age"] = []string{"1", "2", "3", "4", "5", "6"}
	answers["occupation"] = []string{"plumber", "teacher", "engineer", "lawyer", "doctor", "business analyst"}
	answers["computing device"] = []string{"iPad", "Macbook", "Iphone", "Kindle", "Android tablet", "Television"}
	answers["body part"] = []string{"wrist", "arm", "ankle", "thigh", "head", "ears"}
	answers["mood word"] = []string{"sad", "happy", "elated", "grumpy", "calm", "agitated"}
	answers["action word"] = []string{"playing a song", "vibrating", "playing a joke"}

	name := answers["name"][rand.Intn(len(answers["name"]))]
	age := answers["age"][rand.Intn(len(answers["age"]))]
	occupation := answers["occupation"][rand.Intn(len(answers["occupation"]))]
	computing := answers["computing device"][rand.Intn(len(answers["computing device"]))]
	body := answers["body part"][rand.Intn(len(answers["body part"]))]
	mood := answers["mood word"][rand.Intn(len(answers["mood word"]))]
	action := answers["action word"][rand.Intn(len(answers["action word"]))]

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s is a a %s-year old %s who has been struggling with a lot of job-related stress. He/she decided to try a new application to relieve stress, which runs on a/an %s to help improve his/her mood.\n\nThe application senses his/her mood through a device he/she wears on his/her %s.\n\nWhen the device senses that he/she is %s, it respond by %s.", name, age, occupation, computing, body, mood, action)
}

func main() {

	http.HandleFunc("/joke", joke)
	http.HandleFunc("/madlib", madlib)

	fmt.Println("The server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
