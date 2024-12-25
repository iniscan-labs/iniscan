package cmd

import (
	"fmt"
	"github.com/iniscan-labs/iniscan/cmd/migrate"
	"github.com/iniscan-labs/iniscan/cmd/server"
	"github.com/iniscan-labs/iniscan/pkg"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "iniscan",
	Short:        "iniscan",
	SilenceUsage: true,
	Long:         "iniscan is a opensource project for scan ethereum chain",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return fmt.Errorf(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `welcome to use ` + pkg.Green(`iniscan `+pkg.Version) + ` you can use ` + pkg.Red(`-h`) + ` to see the command`
	usageStr += "\nalso can see https://scan-docs.iniscan.com for more detail"
	fmt.Println(usageStr)
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
