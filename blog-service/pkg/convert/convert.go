package convert

import "strconv"

type StrTo string

func (s StrTo) String() string  {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	value, err := strconv.Atoi(s.String())
	return value, err
}

func (s StrTo) MustInt() int {
	value, _ := s.Int()
	return value
}

func (s StrTo) Unit32() (uint32, error) {
	value, err := strconv.Atoi(s.String())
	return uint32(value), err
}

func (s StrTo) MustUInt32() uint32 {
	value, _ := s.Unit32()
	return value
}

