package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "sortex",
	Long: "Sort your files with EXIF data by date on directories. *The numbers down below in the description of the arguments, mean that are compatible with each other - ones with ones, twos with twos, nought only with itself.",
	Run:  sort,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
