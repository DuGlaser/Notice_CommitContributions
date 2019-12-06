package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
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
		Viewer struct {
			ContrbutionsCollection struct {
				TotalCommitContributions githubv4.Int
			} `graphql:"contributionsCollection(from: \"2019-12-02T00:00:00\", to: \"2019-12-03T00:00:00\")"`
		}
	}
	err := client.Query(context.Background(), &q, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(q.Viewer.ContrbutionsCollection.TotalCommitContributions)
}
