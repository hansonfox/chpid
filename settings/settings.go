package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var StcodeMap map[string]string

// var stcodeMap = map[string]string{"612327": "ly","612328":"zb","612325":"mx","612326":"nq"}

var Weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var CheckList = [12]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2", "1"}

func init() {

	f, err := ioutil.ReadFile("stcode01.json") //go run main.go运行时使用相对路径 且应保证stcode文件存在于指定位置
	// f, err := ioutil.ReadFile("./stcode01.json") //go run main.go运行时使用相对路径 且应保证stcode文件存在于指定位置

	if err != nil {
		fmt.Printf("区划代码文件读取失败！\nLoading stcode file fail! ")
		os.Exit(1)
	}

	err = json.Unmarshal(f, &StcodeMap)
	if err != nil {
		fmt.Println(err)
	}
}
