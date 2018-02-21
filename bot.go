package main

import (
	"time"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func tweet(){

	//Suas Keys
	consumerKey := ""
	consumerSecret := ""
	accessToken := ""
	accessSecret := ""
	
	baseTweetText := " ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn"

	//oauth config
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	
	httpClient := config.Client(oauth1.NoContext, token)
	
	client := twitter.NewClient(httpClient)

	searchTweetParams := &twitter.SearchTweetParams{
		Query : "#cthulhu exclude:retweets",
		ResultType: "recent",
		Count : 1,				
	}
	
	searchResult, _, _ := client.Search.Tweets(searchTweetParams)
	   
	for _, tweet := range searchResult.Statuses {		

		fmt.Printf("\nUser: %+v\n", tweet.User.ScreenName)
		fmt.Printf("Tweet Text: %+v\n", tweet.Text)

		client.Statuses.Update("@"+tweet.User.ScreenName + baseTweetText, &twitter.StatusUpdateParams {
			InReplyToStatusID : tweet.ID,			
		})
		
	}	
	fmt.Printf("\n#\n")
}

func main() {   
	tweet()
    ticker := time.NewTicker(time.Minute * 15)
    go func() {
        for _ = range ticker.C {            
            tweet()
        }
    }()
    time.Sleep(time.Hour * 1)    
}
