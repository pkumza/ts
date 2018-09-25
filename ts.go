package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	var cmdStreamlog = &cobra.Command{
		Use:   "e",
		Short: "encode",
		Long:  `ts e 1537879483 -> 2018/9/25 20:44:43`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("Usage: ts e 1537879483")
			}
			tsInt64, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				fmt.Println("Usage: ts e 1537879483, your input is not int")
			}
			tm := time.Unix(tsInt64, 0)
			fmt.Println(tm.Format("2006-01-02 15:04:05"))
		},
	}
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ts",
		Long:  `All software has versions. This is ts's`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("ts v0.1 -- HEAD")
		},
	}
	var rootCmd = &cobra.Command{Use: "ts"}
	rootCmd.AddCommand(cmdStreamlog)
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()

}
