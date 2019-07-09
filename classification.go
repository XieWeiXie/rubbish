package rubbish

import "fmt"

const (
	RECYCKEABLE = iota
	HAZARDOUS
	HOUSEHOLDFOOD
	RESIDUAL
)

// english
const (
	RecyclableWaste    = "Recyclable Waste"
	HazardousWaste     = "Hazardous Waste"
	HouseholdFoodWaste = "Household Food Waste"
	ResidualWaste      = "Residual Waste"
)

// chinese
const (
	RecyclableWasteCN    = "可回收垃圾"
	HazardousWasteCN     = "有害垃圾"
	HouseholdFoodWasteCN = "湿垃圾"
	ResidualWasteCN      = "干垃圾"
)

type Waste interface {
	Define() string
	Requirement() []string
	Help() string
}

type WasteClass struct {
	Name string
}

func (w WasteClass) Define() string {
	return ""
}

func (w WasteClass) Requirement() []string {
	return []string{}
}

func (w WasteClass) Help() string {
	return ""
}

func WhatIsRecW() string {
	return fmt.Sprintf("%s: 废纸张、废塑料、废玻璃制品、废金属、废织物等适宜回收、可循环利用的生活废弃物", RecyclableWasteCN)
}

func HelpRecW() string {
	return fmt.Sprintf("%s投放要求: *轻投轻放\n *清洁干燥、避免污染、废纸尽量平整\n *立体包装物轻清空内容物，清洁后压扁投放\n *有尖锐边角的，应包裹后投放", RecyclableWasteCN)
}

func WhatIsHW() string {
	return fmt.Sprintf("%s: 废电池、废灯管、废药品、废油漆及其容器等对人体健康或者自然环境造成直接或者潜在危害的生活废弃物", HazardousWasteCN)
}

func HelpHw() string {
	return fmt.Sprintf("%s投放要求: *投放时请注意轻放\n *易破损的请连带包装或包裹后轻放\n *如易挥发，请密封后投放", HazardousWasteCN)
}

func WhatIsHFW() string {
	return fmt.Sprintf("%s: 即易腐垃圾，是指食材废料、剩菜剩饭、过期食品、瓜皮果核、花卉绿植、中药药渣等生物质生活废弃物", HouseholdFoodWasteCN)
}

func HelpHFW() string {
	return fmt.Sprintf("%s投放要求: *纯流质的食物垃圾，如牛奶等，应直接倒进下水口\n *有包装的湿垃圾应将包装物去除后分类投放，包装物请投放到对应的可回收来讲或者干垃圾容器", HouseholdFoodWasteCN)
}

func WhatIsResW() string {
	return fmt.Sprintf("%s: 即其他垃圾，是指除可回收垃圾、有害垃圾、湿垃圾织物的其他生活废弃物", ResidualWasteCN)
}

func HelpResW() string {
	return fmt.Sprintf("%s投放要求: *尽量沥干水分\n *难以辨识类别的生活垃圾投入干垃圾容器内")
}
