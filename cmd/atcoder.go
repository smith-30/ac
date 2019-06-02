package cmd

import (
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/smith-30/acc/color"
	"github.com/smith-30/acc/domain"
	"github.com/smith-30/acc/infra/client"
	"github.com/spf13/cobra"
)

var (
	url     = ""
	execCmd = ""
)

// atcoderCmd represents the atcoder command
var atcoderCmd = &cobra.Command{
	Use:   strings.ToLower(domain.Atcoder.String()),
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("start atcoder test %v\n\n", verInfo())

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

		cmds := strings.Split(execCmd, " ")

		for idx, item := range cs {
			cmd := exec.Command(cmds[0], cmds[1:]...)
			stdin, _ := cmd.StdinPipe()
			io.WriteString(stdin, item.Content)
			stdin.Close()
			out, err := cmd.CombinedOutput()

			fmt.Printf("Case [%d] exp: %s", idx, item.Exp)
			if err != nil {
				fmt.Println(color.Redf("\texecute error: %s because %s", err, string(out)))
				continue
			}
			if string(out) == item.Exp {
				fmt.Println("\t" + color.Green("ok!"))
			} else {
				fmt.Println(color.Red("\terror"))
				fmt.Println(fmt.Sprintf("\t\tyour answer is %s", string(out)))
				fmt.Printf("\targument\n\t\t%s\n", item.Content)
			}
		}
	},
}

func init() {
	atcoderCmd.Flags().StringVarP(&url, "url", "u", "", "test case acquisition destination (required)")
	atcoderCmd.Flags().StringVarP(&execCmd, "cmd", "c", "", "your command to pass test case (required)")

	atcoderCmd.MarkFlagRequired("url")
	atcoderCmd.MarkFlagRequired("cmd")

	rootCmd.AddCommand(atcoderCmd)
}
