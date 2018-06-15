package main

import (
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/walk"
	"math/rand"
)

func main() {
	msgmap := make(map[int]string)
	msgmap [0] = "狼烟风沙口"
	msgmap [1] = "让我做你的眼睛"
	msgmap [2] = "that girl"
	msgmap [3] = "萨瓦迪卡"
	msgmap [4] = "继续点吧！"
	msgmap [5] = "go go go ole ole ole!"
	msgmap [6] = "点够五十次你就可以出去了"

	MainWindow{
		Title:   "测试",
		MinSize: Size{270, 290},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 1},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							Label{
								Text: "一般人都没勇气点下面的按钮！你呢？",
							},
						},
					},
				},
			},

			PushButton{
				Text:    "开始冒险",
				MinSize: Size{60, 40},
				OnClicked: func() {
					flag := 0
					var tmp walk.Form
					var k int
					for {
						if flag > 50 {
							break
						}
						k = rand.Intn(6)
						walk.MsgBox(tmp, "exe", msgmap[k], walk.MsgBoxYesNo)
						flag ++
					}
					return
				},
			},
		},
	}.Run()
}
