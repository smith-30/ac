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
		fmt.Printf("start atcoder test %v\n", verInfo())

		// test case 取得
		// cs, err := atcoder_client.GetTestCase()
		// if err != nil {
		// 	fmt.Printf("%v", err)
		// 	return
		// }

		// for idx, item := range cs {
		// 	out, err := exec.Command("ls", "-la").Output()
		// 	if string(out) == item.Exp {
		// 		fmt.Print
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(atcoderCmd)

	// rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "example.toml", "config file name")
}
