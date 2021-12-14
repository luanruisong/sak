package mathx

import (
	"fmt"
)

func conv2float64(i interface{}) (ret float64, err error) {
	switch i.(type) {
	case float32:
		return float64(i.(float32)), nil
	case float64:
		return i.(float64), nil
	default:
		return 0, fmt.Errorf("%v is  not float", i)
	}
}

func SumFloat(p ...interface{}) (ret float64, err error) {
	for _, v := range p {
		var curr float64
		if curr, err = conv2float64(v); err != nil {
			ret = 0
			return
		}
		ret += curr
	}
	return
}

func MaxFloat(p ...interface{}) (ret float64, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if ret, err = conv2float64(p[0]); err != nil {
		ret = 0
		return
	}
	for _, v := range p {
		var curr float64
		if curr, err = conv2float64(v); err != nil {
			ret = 0
			return
		}
		if curr > ret {
			ret = curr
		}
	}
	return
}

func MinFloat(p ...interface{}) (ret float64, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	if ret, err = conv2float64(p[0]); err != nil {
		ret = 0
		return
	}
	for _, v := range p {
		var curr float64
		if curr, err = conv2float64(v); err != nil {
			ret = 0
			return
		}
		if curr < ret {
			ret = curr
		}
	}
	return
}
