package utils

import (
	"fmt"
	"strconv"
	"time"
)

type IDInfo struct {
	stcode    string
	birth     string
	age       int
	gender    string
	checkcode string
}

func (i *IDInfo) Display() string {
	return fmt.Sprintf("区划：%-15s\n出生日期：%-15s\n性别：%-15s\n年龄：%-15d\n", i.stcode, i.birth, i.gender, i.age)
}

func calAge(brith string) int {
	year, _ := time.Parse("2006-01-02", brith)
	return time.Now().Year() - year.Year()
}

//id num parse
func Parse(id string) (string, error) {
	err := ValidatorOne(id)
	if err != nil {
		return "", err
	}
	gender := func() string {
		g, _ := strconv.Atoi(id[len(id)-2 : len(id)-1])
		if err != nil {
			return "params type trans error,please check input"
		}
		if g%2 == 0 {
			return "Femal"
		} else {
			return "Male"
		}
		// return id[len(id)-2 : len(id)-1]
	}()
	age := calAge(id[6:10] + "-" + id[10:12] + "-" + id[12:14])
	info := &IDInfo{
		stcode:    StcodeMap[id[:6]],
		birth:     id[6:10] + "-" + id[10:12] + "-" + id[12:14],
		age:       age,
		gender:    gender,
		checkcode: id[len(id)-1:],
	}
	return info.Display(), nil
}
