package database

type TypeInt int

const (
	H1   TypeInt = iota // EnumIndex = 1
	H2                  // EnumIndex = 2
	H3                  // EnumIndex = 3
	p                   // EnumIndex = 4
	link                // EnumIndex = 5
	code                // EnumIndex = 6
)

func (w TypeInt) String() string {
	return [...]string{"H1", "H2", "H3", "p", "link", "code"}[w]
}

func (w TypeInt) TypeEnumIndex() int {
	return int(w)
}
