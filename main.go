package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	localRootFlag      bool
	persistentRootFlag bool
	times              int
	newFile            string
	rootCmd            = &cobra.Command{
		Use:   "Example",
		Short: "An example cobra program",
		Long: `This is a simple example of a cobra program.
It will have several subcommands and flags.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from the root command")
		},
	}
	echoCmd = &cobra.Command{
		Use:   "echo [strings to echo]",
		Short: "prints given strings to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}
	timesCmd = &cobra.Command{
		Use:   "times [strings to echo]",
		Short: "prints given strings to stdout multiple times",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if times == 0 {
				return errors.New("times cannot be 0")
			}
			for i := 0; i < times; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
			return nil
		},
	}
	createFileCmd = &cobra.Command{
		Use:   "newFile [new file name]",
		Short: "creates a text file with the name provided",
		Long: `This command creates a .txt file with the name of the input
provided from the command line`,
		Args: cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if newFile == "" {
				return errors.New("a file name must be provided")
			}

			var path string = "newFile.txt"
			var content string = "This is a test text file"

			file, err := os.Create(path)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println("file created", file)
			}

			addContent, err := file.WriteString(content + "\n")
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Sprintln("Content %d added to file %f", string(addContent), file.Name())
				file.Sync()
			}

			return nil
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistentRootFlag, "persistFlag", "p", false, "a persistent root flag")
	rootCmd.PersistentFlags().StringVarP(&newFile, "newFile", "n", " ", "new file name")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(timesCmd)
	rootCmd.AddCommand(createFileCmd)

}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
