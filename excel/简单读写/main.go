package main

import (
	"github.com/tealeg/xlsx"
	"strconv"
)

func Makefile() error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("sheet1")
	if err != nil {
		return err
	}

	style := xlsx.NewStyle()
	alignment := xlsx.Alignment{
		Horizontal: "center",
		Vertical:   "center",
	}
	style.Alignment = alignment
	style.ApplyAlignment = true

	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "测试一下"
	cell.SetStyle(style)

	cell.Merge(4, 0)
	row2 := sheet.AddRow()
	for i := 0; i < 5; i++ {
		name := row2.AddCell()
		name.Value = "掌声"
		name.SetStyle(style)
	}

	row3 := sheet.AddRow()
	mergeCell := row3.AddCell()
	mergeCell.Value = "合并"
	mergeCell.SetStyle(style)
	mergeCell.Merge(1, 1)

	row4 := sheet.AddRow()
	row4.AddCell()

	for i := 0; i < 10; i++ {
		row := sheet.AddRow()
		nameCell := row.AddCell()
		nameCell.Value = "xxx" + strconv.Itoa(i)

		ageCell := row.AddCell()
		ageCell.Value = strconv.Itoa(18)

		phoneCell := row.AddCell()
		phoneCell.Value = "stu.Phone"

		genderCell := row.AddCell()
		genderCell.Value = "stu.Gender"

		mailCell := row.AddCell()
		mailCell.Value = "stu.Mail"
	}

	err = file.Save("a.xlsl")
	if err != nil {
		return err
	}

	return nil
}

func main() {
	Makefile()
}
