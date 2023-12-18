package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "bosun",
		Short: "Bosun is a cli tool for creating Helm Release templates",
	}
	configFlags(rootCmd)
	rootCmd.AddCommand(generateCmd())
	return rootCmd
}

func configFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().BoolP("bootstrap", "b", false, "Generate the template in your bootstrap directory")
}
