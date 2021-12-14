package mathx

import "fmt"

func conv2int64(i interface{}) (ret int64, err error) {
	switch i.(type) {
	case uint:
		return int64(i.(uint)), nil
	case uint8:
		return int64(i.(uint8)), nil
	case uint16:
		return int64(i.(uint16)), nil
	case uint32:
		return int64(i.(uint32)), nil
	case uint64:
		return int64(i.(uint64)), nil
	case int:
		return int64(i.(int)), nil
	case int8:
		return int64(i.(int8)), nil
	case int16:
		return int64(i.(int16)), nil
	case int32:
		return int64(i.(int32)), nil
	case int64:
		return i.(int64), nil
	default:
		return 0, fmt.Errorf("%v is  not int", i)
	}
}

func SumInt(p ...interface{}) (ret int64, err error) {
	for _, v := range p {
		var curr int64
		if curr, err = conv2int64(v); err != nil {
			ret = 0
			return
		}
		ret += curr
	}
	return
}

func MaxInt(p ...interface{}) (ret int64, err error) {
	for _, v := range p {
		var curr int64
		if curr, err = conv2int64(v); err != nil {
			ret = 0
			return
		}
		if curr > ret {
			ret = curr
		}
	}
	return
}

func MinInt(p ...interface{}) (ret int64, err error) {
	for _, v := range p {
		var curr int64
		if curr, err = conv2int64(v); err != nil {
			ret = 0
			return
		}
		if curr < ret {
			ret = curr
		}
	}
	return
}
