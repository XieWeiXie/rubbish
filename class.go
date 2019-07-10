package rubbish

import "fmt"

const (
	RECYCKEABLE = iota + 1
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
	Help() []string
}

var Recyclable = WasteClass{Name: ResidualWasteCN}
var Hazardous = WasteClass{Name: HazardousWasteCN}
var HouseHoldFood = WasteClass{Name: HouseholdFoodWasteCN}
var Residual = WasteClass{Name: ResidualWasteCN}

type WasteClass struct {
	Name string
}

func define(name string) string {
	if name == RecyclableWaste || name == RecyclableWasteCN {
		return fmt.Sprintf("%s: 废纸张、废塑料、废玻璃制品、废金属、废织物等适宜回收、可循环利用的生活废弃物", RecyclableWasteCN)
	}
	if name == HazardousWaste || name == HazardousWasteCN {
		return fmt.Sprintf("%s: 废电池、废灯管、废药品、废油漆及其容器等对人体健康或者自然环境造成直接或者潜在危害的生活废弃物", HazardousWasteCN)
	}
	if name == HouseholdFoodWaste || name == HouseholdFoodWasteCN {
		return fmt.Sprintf("%s: 即易腐垃圾，是指食材废料、剩菜剩饭、过期食品、瓜皮果核、花卉绿植、中药药渣等生物质生活废弃物", HouseholdFoodWasteCN)
	}
	if name == ResidualWaste || name == ResidualWasteCN {
		return fmt.Sprintf("%s: 即其他垃圾，是指除可回收垃圾、有害垃圾、湿垃圾织物的其他生活废弃物", ResidualWasteCN)
	}
	return ""
}

func (w WasteClass) Define() string {
	return define(w.Name)
}

func requirement(name string) []string {
	if name == RecyclableWaste || name == RecyclableWasteCN {
		return []string{"轻投轻放", "清洁干燥、避免污染、废纸尽量平整", "立体包装物轻清空内容物，清洁后压扁投放", "有尖锐边角的，应包裹后投放"}
	}
	if name == HazardousWaste || name == HazardousWasteCN {
		return []string{"投放时请注意轻放", "易破损的请连带包装或包裹后轻放", "如易挥发，请密封后投放", HazardousWasteCN}
	}
	if name == HouseholdFoodWaste || name == HouseholdFoodWasteCN {
		return []string{"纯流质的食物垃圾，如牛奶等，应直接倒进下水口", "有包装的湿垃圾应将包装物去除后分类投放，包装物请投放到对应的可回收来讲或者干垃圾容器"}
	}
	if name == ResidualWaste || name == ResidualWasteCN {
		return []string{"尽量沥干水分", "难以辨识类别的生活垃圾投入干垃圾容器内"}
	}
	return []string{}
}

func (w WasteClass) Requirement() []string {
	return requirement(w.Name)
}

func (w WasteClass) Help() []string {
	return w.Requirement()
}
