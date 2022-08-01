package convert

import "strconv"

type IntTo int
type UintTo uint

func (t IntTo) String() string {
	return string(rune(t))
}

func (t UintTo) String() string {
	return strconv.Itoa(int(t))
}