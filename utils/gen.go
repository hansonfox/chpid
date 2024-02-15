package utils

import (
	"chpid/settings"
	"fmt"
	// "log"

	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

//random generator
func RandGen() string {
	stcode := Randstcode()
	date := RandDate()
	seri := RandSeri()
	gender := RandGendercode()
	base := stcode + date + seri + gender
	cbit := GenCheckbit(base)
	return fmt.Sprintf("%s%s", base, cbit)
}

func Randstcode() string {
	rand.Seed(time.Now().UnixNano())

	var stcodeArray []string
	for key := range settings.StcodeMap {
		stcodeArray = append(stcodeArray, key)
	}
	rIndex := rand.Intn(len(stcodeArray))
	return stcodeArray[rIndex]
}

func RandDate() string {
	d, _ := RandBirthday_v2()
	return fmt.Sprintf("%d%02d%02d", d.year, d.month, d.day)
}

func RandSeri() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%02d", rand.Intn(100))
}

func RandGendercode() string {
	rand.Seed(time.Now().UnixNano())
	genderCode := rand.Intn(10)
	return strconv.Itoa(genderCode)
}

func GenCheckbit(base string) string {
	ss := strings.Split(base, "")
	return CalCheckbit(ss)
}

func randGenCo(c chan string, wg *sync.WaitGroup) {
	s := RandGen()
	c <- s
	wg.Done()
}

func writeToFileCo(c chan string, done chan bool, fn string) {
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("ERROR open file error:", err)
		return
	}
	for d := range c {
		_, err = fmt.Fprintln(f, d)

		// fmt.Println("write one to file", d)
		if err != nil {
			fmt.Println("ERROR file write error:", err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println("ERROR file closed error:", err)
		done <- false
		return
	}
	done <- true
}

func RandGenNCo(n int, fn string) {
	c := make(chan string, 100)
	done := make(chan bool)
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go randGenCo(c, &wg)
	}
	go writeToFileCo(c, done, fn)

	go func() {
		wg.Wait()
		close(c)
	}()
	d := <-done
	if d {
		fmt.Println("File written success!")
	} else {
		fmt.Println("File written failed!")
	}
}
