package cmd

import (
	"fmt"
	"go-skeleton/internal/app"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCmd)
	seedCmd.Flags().Bool("all", false, "Seed all")

}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed database",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, _ := cmd.Flags().GetBool("all")
		_, err := app.New()
		if err != nil {
			panic(err)
		}

		fmt.Println(all)

		return nil
	},
}
