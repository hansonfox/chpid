package cmdl

import (
	"log"
	"strings"

	"chpid/utils"

	"github.com/spf13/cobra"
)

var validdesc = strings.Join([]string{
	"该子命令支持身份证校验:",
	"-v:idstr 身份证号校验",
	"-f:fn 批量校验文件",
}, "\n")

var validCmd = &cobra.Command{
	Use:   "valid",
	Short: "身份证校验",
	Long:  validdesc,
	Run: func(cmd *cobra.Command, args []string) {
		// var res error
		res := utils.Validator(idstr, fn)
		if res != nil {
			// fmt.Errorf("校验错误:%s", res)
			log.Printf("校验错误:%s", res)
		} else {
			log.Printf("校验完成")
		}
	},
}

var idstr string
var fn string

func init() {
	validCmd.Flags().StringVarP(&idstr, "idstr", "v", "", "请输入需校验的身份证号")
	validCmd.Flags().StringVarP(&fn, "fn", "f", "", "请输入需要批量校验身份证号txt文件路径")
}
