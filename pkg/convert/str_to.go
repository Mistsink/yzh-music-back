package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	return strconv.Atoi(s.String())
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt() (uint, error) {
	v, err := strconv.Atoi(s.String())
	return uint(v), err
}

func (s StrTo) MustUInt() uint {
	v, _ := s.UInt()
	return v
}