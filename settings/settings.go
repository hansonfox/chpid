package settings

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
)

//go:embed  stcode01.json
var f embed.FS

var StcodeMap map[string]string

var Weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var CheckList = [12]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2", "1"}

func init() {
	conf_path := "stcode01.json"
	// f, err := ioutil.ReadFile("stcode01.json")
	data, err := f.ReadFile(conf_path)
	if err != nil {
		fmt.Printf("区划代码文件读取失败！\nLoading stcode file fail! ")
		os.Exit(1)
	}

	err = json.Unmarshal(data, &StcodeMap)
	if err != nil {
		fmt.Println(err)
	}
}
