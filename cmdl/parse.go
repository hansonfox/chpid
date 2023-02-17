package cmdl

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"chpid/func/utils"
)

var parsedesc = strings.Join([]string{
	"该子命令解析身份证生成相关信息:",
	"-p:id 身份证号解析",
}, "\n")

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "身份证号解析",
	Long:  parsedesc,
	Run: func(cmd *cobra.Command, args []string) {
		// var gen_res []string
		parse_res, err := utils.Parse(id)
		if err != nil {
			log.Printf("ERROR:%s", err)
			os.Exit(1)
		}
		log.Printf("号码信息解析结果：\n%s", parse_res)
	},
}

var id string

func init() {
	parseCmd.Flags().StringVarP(&id, "id", "p", "", "请输入待解析的身份证号码")
}
