package rubbish

type Garbage struct {
	Name  string
	Class int
}

var DefaultGarbage = Garbage{}

func NewGarbage(name string) Garbage {
	return Garbage{
		Name: name,
	}
}

func (g Garbage) IsExists() bool {
	return true
}

func (g Garbage) ClassType() string {
	return ""
}
