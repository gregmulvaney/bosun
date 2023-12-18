package cmd

import (
	"fmt"
	"github.com/gregmulvaney/bosun/internal/tui/generate"
	"github.com/spf13/cobra"
	"os"
)

func generateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "generate",
		Aliases: []string{"g"},
		Short:   "Generate a helm release",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := cmd.Flags().GetBool("bootstrap")
			if err != nil {
				panic(err)
			}
			if _, err = os.Stat("kubernetes"); os.IsNotExist(err) {
				fmt.Println("Error: No kubernetes directory found")
				return
			}
			p := generate.Prompt(true)
			if _, err := p.Run(); err != nil {
				panic(err)
			}
		},
	}
}
