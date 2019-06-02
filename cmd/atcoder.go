package cmd

import (
	"fmt"
	"strings"

	"github.com/smith-30/acc/domain"
	"github.com/smith-30/acc/infra/client"
	"github.com/spf13/cobra"
)

var (
	url = ""
)

// atcoderCmd represents the atcoder command
var atcoderCmd = &cobra.Command{
	Use:   strings.ToLower(domain.Atcoder.String()),
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("start atcoder test %v\n", verInfo())

		c := client.NewClient(domain.Atcoder)
		if c == nil {
			fmt.Printf("can't create client")
			return
		}

		// test case 取得
		cs, err := c.GetTestCase(url)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		fmt.Printf("%#v\n", cs)

		// for idx, item := range cs {
		// 	out, err := exec.Command("ls", "-la").Output()
		// 	if string(out) == item.Exp {
		// 	}
		// }
	},
}

func init() {
	atcoderCmd.Flags().StringVarP(&url, "url", "u", "", "test case acquisition destination (required)")
	atcoderCmd.MarkFlagRequired("url")

	rootCmd.AddCommand(atcoderCmd)
}
