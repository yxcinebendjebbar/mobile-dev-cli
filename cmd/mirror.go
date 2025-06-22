package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Start screen sharing using scrcpy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting scrcpy...")

		scrcpyCmd := exec.Command("scrcpy", "--video-codec=h265", "-m1920", "--max-fps=60", "--no-audio", "-K")
		scrcpyCmd.Stdout = cmd.OutOrStdout()
		scrcpyCmd.Stderr = cmd.ErrOrStderr()

		if err := scrcpyCmd.Run(); err != nil {
			fmt.Println("Failed to start scrcpy:", err)
		}
	},
}
