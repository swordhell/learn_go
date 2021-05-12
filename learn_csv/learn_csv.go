package main

import (
	"encoding/csv"
	"os"
)

type saveCSV struct {
	*os.File
	*csv.Writer
}

func (s *saveCSV) InitEnv(filename string) bool {

	if _, errStat := os.Stat(filename); errStat != nil && os.IsNotExist(errStat) {
		if file, errOpenfile := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644); errOpenfile != nil {
			return false
		} else {
			s.File = file
			file.WriteString("\xEF\xBB\xBF")
			s.Writer = csv.NewWriter(file)

			feildName := [][]string{
				{"LogType", "ID", "ParamNum1", "ParamNum2", "ParamNum3", "Value"},
			}
			s.WriteAll(feildName)
			s.Flush()
		}
	} else if file, errOpenfile := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644); errOpenfile == nil {
		s.File = file
		s.Writer = csv.NewWriter(file) //创建一个新的写入文件流
	}
	return true
}

func (s *saveCSV) AppendData(data [][]string) {
	s.WriteAll(data)
	s.Flush()
}

func (s *saveCSV) TearDown() {
	s.File.Close()
}

func main() {
	s := &saveCSV{}
	s.InitEnv("statCE.csv")

	data := [][]string{
		{"10", "2004", "3001", "1", "0", "8"},
	}

	s.AppendData(data)

	s.TearDown()
}
