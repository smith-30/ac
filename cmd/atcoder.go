package cmd

import (
	"fmt"
	"io"
	"math"
	"os/exec"
	"strconv"
	"strings"

	"github.com/recoilme/pudge"
	"github.com/smith-30/acc/domain"
	"github.com/smith-30/acc/infra/api/http"
	"github.com/smith-30/acc/infra/cache"
	"github.com/smith-30/acc/usecase/atcoder"
	"github.com/smith-30/color"
	"github.com/spf13/cobra"
)

const (
	tab = "\t"
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

		repo, err := cache.NewCacheRepo()
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		// cache db
		defer pudge.CloseAll()

		uc := atcoder.NewUsecase(repo, http.NewHttpClient())

		// test case 取得
		ck, err := domain.NewCacheKey(url)
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		ucmd := atcoder.Command{ck}
		cs, err := uc.Exec(ucmd)
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
			outstr := string(out)

			fmt.Printf("Case [%d] exp: %s", idx, item.Exp)
			if err != nil {
				fmt.Println(color.Redf(tab+"execute error: %s because %s", err, outstr))
				continue
			}
			var ok bool
			outv, err := strconv.ParseFloat(strings.TrimSuffix(outstr, "\n"), 64)
			if err == nil {
				expv, _ := strconv.ParseFloat(strings.TrimSuffix(item.Exp, "\n"), 64)
				if math.Abs(outv-expv) < 0.000000001 {
					ok = true
				}
			} else {
				if outstr == item.Exp {
					ok = true
				}
			}
			if !ok {
				fmt.Println(color.Red(tab + "error"))
				fmt.Println(fmt.Sprintf(tab+tab+"your answer is %s", outstr))
				fmt.Printf(tab+"argument\n%s\n", item.Content)
			} else {
				fmt.Println(tab + color.Green("ok!"))
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
