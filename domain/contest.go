package domain

//go:generate enumer -type=ContestName
type ContestName int

const (
	Atcoder ContestName = iota
)