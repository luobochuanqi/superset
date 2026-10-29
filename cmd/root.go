package cmd

import "github.com/spf13/cobra"

// 有点像 new 一个实例
var rootCmd = cobra.Command{}

func Excute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(wordCmd)
}
