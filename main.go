package main

import (
	"context"
	"log"

	"github.com/thinkingmachines/blossom/pkg/auth"
)

func main() {
	ctx := context.Background()
	client := auth.GetClient(ctx)

	// Test if working...
	repos, _, _ := client.Repositories.List(ctx, "", nil)
	log.Print(repos)
}
