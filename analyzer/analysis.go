package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	re "regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

var nF *os.File

var mainBranches int = 0

var objects []string

func scanDir(dir string) {
	mainBranches++
	fileList, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, e := range fileList {
		f := e.Name()

		switch mode := e.Mode(); {
		case mode.IsDir():
			wg.Add(1)
			go scanDir(dir + "/" + f)

		case mode.IsRegular():
			wg.Add(1)
			go scanFile(dir + "/" + f)
		default:
			fmt.Printf("Encountered an incompatible filetype. Ignoring %s", f)
		}
	}
	mainBranches--
	wg.Done()

}

func scanFile(path string) {

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fullDatas, err := ioutil.ReadAll(file)
	part := string(fullDatas)
	//fullDatas = nil
	if err != nil {
		fmt.Println(err)
		wg.Done()
		return
	}

	for _, obj := range objects {

		if strings.Contains(part, obj) {

			r := re.MustCompile(`(?i)\b` + obj + `\b`)

			m := len(r.FindAllString(part, -1))
			if m > 0 {
				nF.WriteString(path + "," + obj + "," + strconv.Itoa(m) + "\n")
			}
		}
	}

	fullDatas = nil
	wg.Done()
	mainBranches--
}

func check(err error) {

	if err != nil {
		fmt.Println(err)
	}

}

var shortest int = 0

func main() {

	startTime := time.Now()

	cpus := runtime.NumCPU()
	fmt.Println(cpus)
	fmt.Println(runtime.GOMAXPROCS(cpus))
	objF, err := os.Open("export.csv")
	check(err)
	csvR := csv.NewReader(objF)
	check(err)
	allObj, err := csvR.ReadAll()

	for i, r := range allObj {
		if i > 0 {
			if len(string(r[1])) < shortest || shortest == 0 {
				shortest = len(string(r[1]))
			}
			objects = append(objects, string(r[1]))
		}
	}

	wg.Add(1)
	curdir, _ := os.Getwd()

	nF, err = os.Create("output.csv")
	go scanDir(curdir + "/src")
	wg.Wait()
	nF.Close()
	duration := time.Since(startTime).Seconds()
	fmt.Printf("Process Complete in %d seconds", duration)

}
