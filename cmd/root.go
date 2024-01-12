package cmd

import (
	"github.com/spf13/cobra"
)

var (
	flag_outputWidth  int
	flag_outputHeight int
)

var rootCmd = &cobra.Command{
	Use:   "splitter",
	Short: "Splitter is a CLI tool for image splitting",
	Long:  `A longer description...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&flag_outputHeight, "height", "H", 1500, "Sets the height (in pixels) of the resultant image. Default value will result in the output image height being the input image height.")
	rootCmd.PersistentFlags().IntVarP(&flag_outputWidth, "width", "W", 2100, "Sets the width (in pixels) of the resultant image. Default value will result in the output image width being double of the input image width.")
}
