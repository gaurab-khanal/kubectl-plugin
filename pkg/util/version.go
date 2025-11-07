package util

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev" // overridden by goreleaser during build process

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the version of kubectl plugin",
	Long:  `Shows the version information of the KubeStellar kubectl plugin.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kubectl plugin %s\n", Version)
	},
}
