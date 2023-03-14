package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v50/github"
)

var (
	name        = flag.String("name", "TestRepo", "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "This is a test repo", "Description of created repo.")
	private     = flag.Bool("private", false, "Will created repo be private.")
	autoInit    = flag.Bool("auto-init", false, "Pass true to create an initial commit with empty README.")
)

func main() {
	flag.Parse()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}
	ctx := context.Background()
	client := github.NewTokenClient(ctx, token)

	r := &github.Repository{Name: name, Private: private, Description: description, AutoInit: autoInit}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
