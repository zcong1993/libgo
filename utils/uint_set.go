package utils

type UintSet map[uint]struct{}

func NewUintSet(ins ...uint) UintSet {
	us := make(UintSet, len(ins))
	for _, in := range ins {
		us.Add(in)
	}
	return us
}

func (us *UintSet) Add(in uint) {
	(*us)[in] = struct{}{}
}

func (us *UintSet) Adds(ins ...uint) {
	for _, in := range ins {
		us.Add(in)
	}
}

func (us *UintSet) ToSlice() []uint {
	res := make([]uint, len(*us))
	i := 0
	for k, _ := range *us {
		res[i] = k
		i++
	}
	return res
}
