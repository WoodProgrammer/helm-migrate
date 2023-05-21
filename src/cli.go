package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "helm restore",
	Short: "helm restore - is helm plugin allows you to backup and restore helm releases between clusters",
	Long:  `helm restore rocking yeaaa :)) `,
	Run: func(cmd *cobra.Command, args []string) {
		InfoLogger.Println(args[0])
	},
}

var mode = &cobra.Command{
	Use:     "mode",
	Short:   "Type of the operation in cluster",
	Aliases: []string{"mode"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		targetCluster, _ := cmd.Flags().GetString("targetcluster")
		sourceCluster, _ := cmd.Flags().GetString("sourcecluster")
		targetNs, _ := cmd.Flags().GetString("ns")

		WarningLogger.Println("Source cluster is :: ", sourceCluster)
		WarningLogger.Println("Target cluster is ::", targetCluster)
		WarningLogger.Println("Source namespace is ::", targetNs)

		if args[0] == "backup" {
			WarningLogger.Println("Running only backup mode.. Extracting files under this directory...")

			// TODO backup only mode
		} else if args[0] == "restore" {
			WarningLogger.Println("This option provides both backup and restore functionality...")
			sourceClusterclientset := configHandler(sourceCluster)
			targetClusterclientset := configHandler(targetCluster)
			backup := getBackup(targetNs, sourceClusterclientset)

			restoreBackup(targetNs, targetClusterclientset, backup)

		}
	},
}

func Execute() {

	rootCmd.AddCommand(mode)
	mode.PersistentFlags().String("ns", "", "The target namespace to fetch helm release and restore")
	mode.PersistentFlags().String("targetcluster", "", "Source of the backup of helm releases")
	mode.PersistentFlags().String("sourcecluster", "", "Target cluster address of helm restore operation")

	if err := rootCmd.Execute(); err != nil {
		ErrorLogger.Println(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
