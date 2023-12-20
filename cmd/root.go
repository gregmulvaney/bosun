package cmd

import (
	"github.com/gregmulvaney/bosun/internal/tui/browser"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "bosun",
		Short: "Bosun is a cli tool for creating Helm Release templates",
		Run: func(cmd *cobra.Command, args []string) {
			p := browser.Browser()
			if _, err := p.Run(); err != nil {
				panic(err)
			}
		},
	}
	configFlags(rootCmd)
	rootCmd.AddCommand(generateCmd())
	return rootCmd
}

func configFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolP("bootstrap", "b", false, "Generate the template in your bootstrap directory")
}
