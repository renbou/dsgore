package cmd

import (
	"os"

	"github.com/renbou/dsgore/pkg/slayer"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	directory string

	rootCmd = &cobra.Command{
		Use:   "dsgore",
		Short: "Release your hatred upon the .DS_Store files",
		Long: `Turn existing .DS_Store files into a bloody mess
and prevent them from ending up in your commits.`,
		Run: func(cmd *cobra.Command, args []string) {
			errChan := make(chan error, 100)
			// Release the beast (and maybe die if the beast is not strong enough...)
			go slayer.Slay(directory, errChan)
			// Log all errors after we started
			for err := range errChan {
				log.Error().Msg(err.Error())
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&directory, "directory", "d", ".", "Directory where hell will break loose once the slayer is released.")
	setupLogger()
}

func Execute() error {
	return rootCmd.Execute()
}

func setupLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	})
}
