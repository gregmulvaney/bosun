package main

import "github.com/gregmulvaney/bosun/cmd"

func main() {
	rootCmd := cmd.Root()
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
