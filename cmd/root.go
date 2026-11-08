package cmd

import (
	"fmt"
	"os"

	"github.com/JacobsM14/img-relation-cli/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "img-relation",
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
		width, height, err := internal.GetImageDimensions(args[0])
		if err != nil {
			return err
		}
		fmt.Printf("Width: %d, Height: %d, Aspect Ratio: %.2f\n", width, height, float64(width)/float64(height))
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
