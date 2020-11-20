package request

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "use",
	Short: "short",
	Long:  "long",

	Run: func(cmd *cobra.Command, args []string) { fmt.Println("pew pew") },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(addListCmd)
}
