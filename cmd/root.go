package cmd

import "github.com/spf13/cobra"

var rootCMD = &cobra.Command{}

func Execute() error {
	return rootCMD.Execute()
}

func init(){
	rootCMD.AddCommand(wordCmd)
	rootCMD.AddCommand(jsonCmd)
	rootCMD.AddCommand(timeCmd)
}
