package safeconv

import "fmt"

func AddOK[L, R number](l L, r R) (R, bool) {
	lAsR, ok := ConvertOK[L, R](l)
	if !ok {
		return 0, false
	}

	c := lAsR + r
	if (c > lAsR) == (r > 0) {
		return c, true
	}
	return c, false
}

func SubOK[L, R number](l L, r R) (R, bool) {
	lAsR, ok := ConvertOK[L, R](l)
	if !ok {
		return 0, false
	}

	c := lAsR - r
	if (c < lAsR) == (r > 0) {
		return c, true
	}
	return c, false
}

func MulOK[L, R number](l L, r R) (R, bool) {
	if l == 0 || r == 0 {
		return 0, true
	}

	lAsR, ok := ConvertOK[L, R](l)
	if !ok {
		return 0, false
	}

	c := lAsR * r
	if (c < 0) == ((l < 0) != (r < 0)) {
		if c/r == lAsR {
			return c, true
		}
	}
	return 0, false
}

func DivOK[L, R number](l L, r R) (R, bool) {
	if r == 0 {
		return 0, false
	}

	lAsR, ok := ConvertOK[L, R](l)
	if !ok {
		return 0, false
	}

	c := lAsR / r
	if (c < 0) == ((l < 0) != (r < 0)) {
		return c, true
	}
	return 0, false
}

func Add[L, R number](l L, r R) R {
	c, ok := AddOK(l, r)
	if !ok {
		panic(fmt.Sprintf("%[1]T(%[1]v) + %[2]T(%[2]v) cannot be computed safely", l, r))
	}
	return c
}

func Sub[L, R number](l L, r R) R {
	c, ok := SubOK(l, r)
	if !ok {
		panic(fmt.Sprintf("%[1]T(%[1]v) - %[2]T(%[2]v) cannot be computed safely", l, r))
	}
	return c
}

func Mul[L, R number](l L, r R) R {
	c, ok := MulOK(l, r)
	if !ok {
		panic(fmt.Sprintf("%[1]T(%[1]v) * %[2]T(%[2]v) cannot be computed safely", l, r))
	}
	return c
}

func Div[L, R number](l L, r R) R {
	c, ok := DivOK(l, r)
	if !ok {
		panic(fmt.Sprintf("%[1]T(%[1]v) / %[2]T(%[2]v) cannot be computed safely", l, r))
	}
	return c
}
