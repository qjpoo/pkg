package main

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// 列标题
	titles := []string{
		"姓名", "年龄",
	}
	// 数据源
	data := []map[string]interface{}{
		{"name": "jack", "age": 18},
		{"name": "mary", "age": 28},
	}

	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")

	for clumnNum, v := range titles {
		sheetPosition := Div(clumnNum+1) + "1"
		fmt.Print(sheetPosition)
		f.SetCellValue("Sheet1", sheetPosition, v)
	}
	for lineNum, v := range data {
		// Set value of a cell.
		clumnNum := 0
		for _, vv := range v {
			clumnNum++
			sheetPosition := Div(clumnNum) + strconv.Itoa(lineNum+2)
			switch vv.(type) {
			case string:
				f.SetCellValue("Sheet1", sheetPosition, vv.(string))
				break
			case int:
				f.SetCellValue("Sheet1", sheetPosition, vv.(int))
				break
			case float64:
				f.SetCellValue("Sheet1", sheetPosition, vv.(float64))
				break
			}
		}
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book2.xlsx"); err != nil {
		println(err.Error())
	}
}

// Div 数字转字母
func Div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}
