package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type csvRow struct {
	Type    string
	Objects []string
	Removed int
}

type csvFull struct {
}

func main() {

	data, err := os.Open("path/to/AD_diff.html")

	if err != nil {
		fmt.Println(err)
	}

	defer data.Close()

	var lines string

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {

		lines += scanner.Text()

	}

	fe := regexp.MustCompile(`</b>.*?\:`)
	var objectTypes []string
	objectTypes = fe.FindAllString(lines, -1)

	var dataStruct []csvRow

	var fullData []string
	fullTable := regexp.MustCompile(`</b>.*?</table>`)
	fullData = fullTable.FindAllString(lines, -1)

	var rowData []string

	rowReg := regexp.MustCompile(`<tr>.*?</tr>`)
	cellReg := regexp.MustCompile(`<td.*?</td>`)
	objectReg := regexp.MustCompile(`r">.*?<`)

	for i := range fullData {

		var tempData csvRow
		tempData.Type = objectTypes[i][4:]

		rowData = rowReg.FindAllString(fullData[i], -1)
		k := 0
		for j := range rowData {

			matched, _ := regexp.MatchString(`Removed in 12.1.3\s*?<br`, rowData[j])
			if matched {
				var tdArray []string
				tdArray = cellReg.FindAllString(rowData[j], -1)
				objectString := objectReg.FindString(tdArray[0])

				tempData.Objects = append(tempData.Objects, objectString[3:len(objectString)-1])

				k++
			}
		}

		tempData.Removed = k

		dataStruct = append(dataStruct, tempData)
	}
	for i := range dataStruct {
		fmt.Println(dataStruct[i])
	}
}
