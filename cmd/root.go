package cmd

import (
	"fmt"
	"os"

	"github.com/JacobsM14/imgdim/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "imgdim",
	Short:         "Show image dimensions and aspect ratio",
	SilenceUsage:  true,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return []string{"png", "jpg", "jpeg", "gif"}, cobra.ShellCompDirectiveFilterFileExt
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		info, err := internal.GetImageInfo(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Width: %d\nHeight: %d\nAspect Ratio: %.2f\nOrientation: %s\nSimplified: %s\nCommon Ratio: %s\n", info.Width, info.Height, info.Ratio, info.Orientation, info.Simplified, info.Common)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
