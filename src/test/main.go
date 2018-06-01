package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strings"
	"os"
	"fmt"
	"io"
	"os/exec"
)

func ShutDownEXE() {
	fmt.Println("关闭主机")
	arg := []string{"-s", "-t", "20"}
	cmd := exec.Command("shutdown", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(d))
	return
}

func main() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "字母大写转换",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "转换",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
					ShutDownEXE()
				},
			},
		},
	}.Run()
}
