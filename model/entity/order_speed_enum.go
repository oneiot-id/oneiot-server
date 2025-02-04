package entity

type OrderSpeed int

const (
	Regular   OrderSpeed = iota
	Express   OrderSpeed = iota
	FullSpeed OrderSpeed = iota
)

func (s OrderSpeed) String() string {
	return [...]string{"Regular", "Express", "Full Speed"}[s]
}
