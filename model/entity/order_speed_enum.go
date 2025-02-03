package entity

type OrderSpeed int

const (
	Regular   OrderSpeed = iota
	Express              = iota
	FullSpeed            = iota
)

func (s OrderSpeed) String() string {
	return [...]string{"Regular", "Express", "FullSpeed"}[s]
}
