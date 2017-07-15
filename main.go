package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/machinebox/sdk-go/textbox"
)

var consumerKey = os.Getenv("TWITTER_CONSUMER")
var consumerSecret = os.Getenv("TWITTER_SECRET")
var consumerToken = os.Getenv("TWITTER_TOKEN")
var consumerTokenSecret = os.Getenv("TWITTER_TOKEN_SECRET")

func main() {
	tweets := getTweets()
	getSentiment(tweets)
}

func getSentiment(tweets []twitter.Tweet) {
	client := textbox.New("http://textbox.demo.gs")
	for _, tweet := range tweets {
		analysis, err := client.Check(strings.NewReader(tweet.Text))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(tweet.Text)
		fmt.Println()

		fmt.Println(analysis.Keywords)

		sentimentTotal := 0.0
		for _, sentence := range analysis.Sentences {
			sentimentTotal += sentence.Sentiment
		}

		// higher is more positive, lower is more negative
		fmt.Println("Sentiment:", sentimentTotal/float64(len(analysis.Sentences)))

		fmt.Println()
		fmt.Println()
	}
}

func getTweets() []twitter.Tweet {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(consumerToken, consumerTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, _, err := client.Timelines.HomeTimeline(
		&twitter.HomeTimelineParams{
			Count: 20,
		},
	)
	if err != nil {
		fmt.Println(err)
	}

	return tweets
}
