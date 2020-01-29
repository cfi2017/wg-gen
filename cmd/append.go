/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/cfi2017/wg-gen/pkg"
	"github.com/spf13/cobra"
)

var (
	publicKey string
	dryRun    bool
)

// appendCmd represents the append command
var appendCmd = &cobra.Command{
	Use:   "append",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var peer pkg.Peer
		var err error
		if publicKey != "" {
			peer, err = pkg.GeneratePeerWithPublicKey(publicKey)
		} else {
			peer, err = pkg.GeneratePeer()
		}
		if err != nil {
			panic(err)
		}
		// output public key for later use
		fmt.Println(peer.PublicKey.String())
		if dryRun {
			fmt.Println("--------- CONFIG ---------")
			str, err := peer.String()
			if err != nil {
				panic(err)
			}
			fmt.Print(str)
			return
		}
		err = pkg.AppendToConfigFile(pkg.ConfigFile, peer)
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(appendCmd)
	appendCmd.Flags().StringVarP(&publicKey, "public-key", "p", "", "use public key instead of generating")
	appendCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "dry run")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}