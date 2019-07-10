package main

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/rubbish"
)

func main() {
	var a rubbish.Garbage
	a.Name = "八宝粥"
	fmt.Println(a.IsExists())
	fmt.Println(a.ClassType())
	fmt.Println(a.Help())
	fmt.Println(a.Requirement())
	fmt.Println(a.Define())
	var exampleRubbish = []string{"塑料袋", "西瓜皮", "桌子", "瓜子壳", "湿巾纸"}
	for _, i := range exampleRubbish {
		tempRubbish := rubbish.NewGarbage(i)
		if tempRubbish.IsExists() {
			fmt.Println(i, tempRubbish.ClassType())
		} else {
			fmt.Println(i, "No Data")
		}
	}
}
