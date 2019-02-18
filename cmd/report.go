package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/ljvmiranda921/blossom/pkg/auth"
	"github.com/ljvmiranda921/blossom/pkg/metrics"
	"github.com/spf13/cobra"
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report <username>/<repository>",
	Short: "Get open-source software metrics in stdout",
	Long: `The report command is the main entrypoint for obtaining
open-source software metrics related to your project. You can pass several
flags to check various aspects of the project`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		targetSlice := strings.Split(target, "/")

		ctx := context.Background()
		client := auth.GetClient(ctx)

		repo, _, _ := client.Repositories.Get(ctx, targetSlice[0], targetSlice[1])
		dm := metrics.GetDiscoveryMetrics(repo)
		log.Print(dm)
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.PersistentFlags().BoolP("discovery", "d", false, "Get discovery metrics")
}
