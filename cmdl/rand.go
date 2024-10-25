package cmdl

import (
	// "log"
	"strings"

	"github.com/spf13/cobra"

	"chpid/utils"
)

var randdesc = strings.Join([]string{
	"该子命令支持身份证随机生成:",
	"-r:num 身份证号随机生成",
	"-o:out 输出文件名称",
}, "\n")

var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "身份证随机生成",
	Long:  randdesc,
	Run: func(cmd *cobra.Command, args []string) {
		// var gen_res []string
		utils.RandGenNCo(num, out)
	},
}

var num int
var out string

func init() {
	randCmd.Flags().IntVarP(&num, "num", "r", 1, "请输入随机生成数量")
	randCmd.Flags().StringVarP(&out, "out", "o", "default_res.txt", "输出至文件")
}
