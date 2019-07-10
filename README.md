# Rubbish

> 垃圾分类库


## Install

```
go get github.com/wuxiaoxiaoshen/rubblish
```

## Basic Usage

```
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

>>
true
湿垃圾
纯流质的食物垃圾，如牛奶等，应直接倒进下水口; 有包装的湿垃圾应将包装物去除后分类投放，包装物请投放到对应的可回收来讲或者干垃圾容器
纯流质的食物垃圾，如牛奶等，应直接倒进下水口; 有包装的湿垃圾应将包装物去除后分类投放，包装物请投放到对应的可回收来讲或者干垃圾容器
湿垃圾: 即易腐垃圾，是指食材废料、剩菜剩饭、过期食品、瓜皮果核、花卉绿植、中药药渣等生物质生活废弃物
塑料袋 干垃圾
西瓜皮 湿垃圾
桌子 No Data
瓜子壳 湿垃圾
湿巾纸 干垃圾

```

## API


1. 类别
2. 投放要求
3. 查询
4. 常见的垃圾


``` 
// 具体的垃圾实例
type Garbage struct {
	Name string
}
```

- IsExists: 数据集中是否存在
- ClassType： 类别：干垃圾、湿垃圾、可回收垃圾、有害垃圾
- Requirement： 投放要求
- Define: 概念，比如干垃圾是什么
- Help: 投放要求


## Data

> 参考至诸多小程序。[数据集](https://github.com/wuxiaoxiaoshen/Collection)


Author: @wuxiaoxioshen



