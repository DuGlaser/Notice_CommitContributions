package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func getTime() (time.Time, time.Time) {
	const layout = "2006-01-02T15:04:05"
	t := time.Now()
	y := time.Now().AddDate(0, 0, -1)
	return t, y
}

func main() {
	today, yesterday := getTime()
	loadErr := godotenv.Load()
	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var q struct {
		User struct {
			ContrbutionsCollection struct {
				TotalCommitContributions githubv4.Int
			} `graphql:"contributionsCollection(from: $Yesterday,to: $Today)"`
		} `graphql:"user(login: \"DuGlaser\")"`
	}
	variable := map[string]interface{}{
		"Today":     githubv4.DateTime{today},
		"Yesterday": githubv4.DateTime{yesterday},
	}

	err := client.Query(context.Background(), &q, variable)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(q.User.ContrbutionsCollection.TotalCommitContributions)
}
