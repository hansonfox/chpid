package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"

	// "os/exec"
	// "path/filepath"
	"chpid/settings"
)

// func init() {
// 	// 解除下三行注释 并引入 exec, filepath 处理go build 后读取json文件路径问题
// 	// cmd_f, _ := exec.LookPath(os.Args[0])
// 	// path, _ := filepath.Abs(cmd_f)
// 	// i := strings.LastIndex(path, string(os.PathSeparator))
// 	// f, err := ioutil.ReadFile(path[:i] + "/stcode01.json")

// 	f, err := ioutil.ReadFile("G:/study_pro/go_project/chpid/stcode01.json") //go run main.go运行时使用相对路径 且应保证stcode文件存在于指定位置
// 	// f, err := ioutil.ReadFile("./stcode01.json") //go run main.go运行时使用相对路径 且应保证stcode文件存在于指定位置

// 	if err != nil {
// 		fmt.Printf("区划代码文件读取失败！\n Loading stcode file fail! ")
// 		os.Exit(1)
// 	}

// 	err = json.Unmarshal(f, &StcodeMap)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// for key, value := range StcodeMap {
// 	// 	fmt.Printf("%v--%v", key, value)
// 	// }
// }

//id string length check
func lenC(s string) error {
	if len(s) != 18 {
		return errors.New("incorrect Id string length")
	}
	return nil
}

//stcode check
func stcodeC(s []string) error {
	if _, ok := settings.StcodeMap[strings.Join(s[0:6], "")]; ok {
		return nil
	} else {
		return errors.New("district code error or not exsit")
	}
}

//Date validate
func dateC(s []string) []error {
	var errArray []error
	yearStr, _ := strconv.Atoi(strings.Join(s[6:10], ""))
	monthStr, _ := strconv.Atoi(strings.Join(s[10:12], ""))
	dayStr, _ := strconv.Atoi(strings.Join(s[12:14], ""))
	d := Date{}
	yerr := d.SetYear(yearStr)
	if yerr != nil {
		errArray = append(errArray, yerr)
	}
	merr := d.SetMonth(monthStr)
	if merr != nil {
		errArray = append(errArray, merr)
	}
	derr := d.SetDay(dayStr)
	if derr != nil {
		errArray = append(errArray, derr)
	}
	return errArray
}

//calculate the check bit number
func CalCheckbit(s []string) string {
	var sum int
	for i, s := range s {
		value, _ := strconv.Atoi(s)
		isum := value * settings.Weight[i]
		sum += isum
	}
	var icheck int = sum % 11
	return settings.CheckList[icheck]
}

//valid the check bit
func checkbitC(s []string) (string, error) {
	checkbit := s[17]
	base := s[:17]

	calcheckbit := CalCheckbit(base)
	if calcheckbit == checkbit {
		return calcheckbit, nil
	} else {
		return calcheckbit, errors.New("wrong check bit")
	}
}

func ValidatorOne(s string) error {
	//length check
	lenErr := lenC(s)
	if lenErr != nil {
		// log.Println(lenErr)
		return fmt.Errorf("validation error:%s", lenErr)
	}
	ss := strings.Split(s, "") //split string to slice

	//stcode check
	stcodeErr := stcodeC(ss)
	if stcodeErr != nil {
		// log.Println(stcodeErr)
		return fmt.Errorf("validation error:%s", stcodeErr)
	}

	//valid date
	dateErr := dateC(ss)
	if dateErr != nil {
		return fmt.Errorf("valiadtion error:%s", dateErr)
	}

	checkbit, checkbitErr := checkbitC(ss)
	if checkbitErr != nil {
		return fmt.Errorf("validation error:%s correct should be %s", checkbitErr, checkbit)
	}

	fmt.Println("Validation passed.")
	return nil
}

func ValidatorBatch(fn string) error {
	target_path := path.Dir(fn)
	fd, err := os.OpenFile(fn, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)

	var result []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		res := ValidatorOne(string(line))
		result = append(result, fmt.Sprintf("%s---%v\n", line, res))
	}
	target_path = path.Join(target_path, "valid_res.txt")
	_, err = os.Create(target_path)
	if err != nil {
		fmt.Println("file creation failed.")
		return err
	}
	f, err := os.OpenFile(target_path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, v := range result {
		_, err := w.Write([]byte(v))
		if err != nil {
			fmt.Println(err)
		}
	}
	w.Flush()
	fmt.Println("batch check results have been writeen completely.")
	return nil
}

func Validator(s string, fn string) error {
	if s != "" {
		err := ValidatorOne(s)
		return err
	}
	if fn != "" {
		ValidatorBatch(fn)
	}
	return nil
}
