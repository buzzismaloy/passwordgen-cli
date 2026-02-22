package cmd

import (
	"fmt"
	"os"
)

var RootCmd = &Command{
	Use:   "passwordgen",
	Short: "CLI password genertator",
	Long:  "CLI wrapper around passwordgen-lib library",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
