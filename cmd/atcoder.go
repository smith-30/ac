// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// atcoderCmd represents the atcoder command
var atcoderCmd = &cobra.Command{
	Use:   "atcoder",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("start atcoder test %v", verInfo())

		// for logging panic
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic ... %#v\n", err)
			}
		}()
	},
}

func init() {
	rootCmd.AddCommand(atcoderCmd)

	// rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "example.toml", "config file name")
}
