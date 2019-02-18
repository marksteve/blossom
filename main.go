package main

import (
	"context"
	"log"

	"github.com/ljvmiranda921/blossom/pkg/auth"
	"github.com/ljvmiranda921/blossom/pkg/metrics"
)

func main() {
	ctx := context.Background()
	client := auth.GetClient(ctx)

	// Test if working...
	repo, _, _ := client.Repositories.Get(ctx, "ljvmiranda921", "blossom")
	dm := metrics.GetDiscoveryMetrics(repo)
	log.Print(dm)
}
