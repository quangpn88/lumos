package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var (
	systems   = []string{"VISAP", "CASAP", "SCARD", "BID", "BIB", "CONSOL"}
	dateStart = time.Date(2015, time.September, 14, 0, 0, 0, 0, time.UTC)
)

func main() {
	Checklog()
	ReadMenuFile(FindLunchmenuFile())
}

func Checklog() {
	duration := time.Since(dateStart)
	days := int(duration.Hours() / 24)
	weekth := days/7 + 1
	position := weekth % len(systems)
	fmt.Println("days:" + strconv.Itoa(days) + "  weekth:" + strconv.Itoa(weekth))
	fmt.Println(systems[position-1])
}

func FindLunchmenuFile() string {
	basePath := "X:/"
	today := time.Now()
	basePath += strconv.Itoa(today.Year())
	monthDirs, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	var monthDirName string
	for _, monthDir := range monthDirs {
		if strings.Contains(today.Month().String(), monthDir.Name()) {
			monthDirName = monthDir.Name()
			basePath += "/" + monthDirName
			break
		}
	}
	weekFiles, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, weekFile := range weekFiles {
		weekFileName := weekFile.Name()
		if !strings.Contains(weekFileName, "~$") {
			indexOfMoth := strings.LastIndex(weekFileName, monthDirName)
			firstDay := weekFileName[indexOfMoth+len(monthDirName)+1 : indexOfMoth+len(monthDirName)+1+2]
			firstDayInt, _ := strconv.Atoi(firstDay)
			if firstDayInt <= today.Day() && today.Day() < firstDayInt+7 {
				return basePath + "/" + weekFileName
			}
		}
	}
	return ""
}

func ReadMenuFile(filePath string) {
	today := time.Now()
	xlFile, _ := xlsx.OpenFile(filePath)
	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			for cellIndex, cell := range row.Cells {
				if strings.Contains(cell.String(), today.Weekday().String()) {
					fmt.Printf("%s\n", cell.String())
					fmt.Println(sheet.Rows[rowIndex+1].Cells[cellIndex].String())
					return
				}

			}
		}
	}
}
