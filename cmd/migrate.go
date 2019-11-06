package cmd

import (
	"fmt"
	"go-skeleton/app"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Bool("up", false, "migrate database up")
	migrateCmd.Flags().Bool("down", false, "migrate database down")
	migrateCmd.Flags().Int("step", 0, "determine how many step migration will takes")
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	RunE: func(cmd *cobra.Command, args []string) error {
		up, _ := cmd.Flags().GetBool("up")
		down, _ := cmd.Flags().GetBool("down")
		step, _ := cmd.Flags().GetInt("step")
		ex, _ := os.Executable()
		exPath := filepath.Dir(ex)

		m, err := migrate.New("file://"+exPath+"/migration", app.GetDbURL())
		if err != nil {
			fmt.Println(err)
		}

		if step != 0 {
			if err := m.Steps(step); err != nil {
				fmt.Println(err)
			}
		} else if up {
			if err := m.Up(); err != nil {
				fmt.Println(err)
			}
		} else if down {
			if err := m.Down(); err != nil {
				fmt.Println(err)
			}
		}

		return nil
	},
}
