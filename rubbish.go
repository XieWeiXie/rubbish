package rubbish

import (
	"strings"

	"github.com/wuxiaoxiaoshen/rubbish/data/Collection/waste"
)

type Garbage struct {
	Name string
}

var DefaultGarbage = Garbage{}

func NewGarbage(name string) Garbage {
	return Garbage{
		Name: name,
	}
}

func (g Garbage) IsExists() bool {
	if _, ok := waste.Waste[g.Name]; !ok {
		return ok
	}
	return true
}

func (g Garbage) ClassType() string {
	if g.IsExists() {
		return waste.Waste[g.Name]
	}
	return "None"
}

func (g Garbage) Classification() WasteClass {
	var w WasteClass
	if g.IsExists() {
		w.Name = g.ClassType()
	}
	return w
}

func (g Garbage) Requirement() string {
	return strings.Join(g.Classification().Requirement(), "; ")
}

func (g Garbage) Define() string {
	return g.Classification().Define()
}

func (g Garbage) Help() string {
	return g.Requirement()
}
