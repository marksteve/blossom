package cmd

import (
	"context"
	"log"
	"strings"

	"github.com/ljvmiranda921/blossom/pkg/auth"
	"github.com/ljvmiranda921/blossom/pkg/metrics"
	"github.com/shurcooL/githubv4"
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

		// Check if owner, repo has been captured
		if len(targetSlice) != 2 {
			log.Fatal("Ensure that target is: <owner>/<repository>")
		}

		owner, repo := targetSlice[0], targetSlice[1]

		ctx := context.Background()
		client := auth.GetClient(ctx)

		request := map[string]interface{}{
			"owner": githubv4.String(owner),
			"name":  githubv4.String(repo),
		}

		m, rl := metrics.Get(ctx, client, request)
		log.Println(m)
		log.Println(rl)

	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
	reportCmd.PersistentFlags().BoolP("discovery", "d", false, "Get discovery metrics")
}
