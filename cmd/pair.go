package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var pairCmd = &cobra.Command{
	Use:   "pair",
	Short: "Pair your Android device wirelessly using ADB",
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		port, _ := cmd.Flags().GetString("port")

		fmt.Println("Starting wireless pairing...")
		pairCmd := exec.Command("adb", "pair", fmt.Sprintf("%s:%s", ip, port))
		pairCmd.Stdin = os.Stdin
		pairCmd.Stdout = cmd.OutOrStdout()
		pairCmd.Stderr = cmd.ErrOrStderr()

		if err := pairCmd.Run(); err != nil {
			fmt.Println("Failed to pair device:", err)
		} else {
			fmt.Println("Successfully paired with device!")
		}
	},
}

func init() {
	pairCmd.Flags().String("ip", "", "Device IP address")
	pairCmd.Flags().String("port", "5555", "Device pairing port")
	pairCmd.MarkFlagRequired("ip")
}
