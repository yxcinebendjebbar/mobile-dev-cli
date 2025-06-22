package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mobile-dev",
	Short: "Mobile Dev CLI - wirelessly connect ADB and mirror your device with scrcpy",
}

func init() {
	RootCmd.AddCommand(pairCmd)
	RootCmd.AddCommand(mirrorCmd)
}
