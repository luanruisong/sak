package mathx

import "fmt"

func splits(p ...interface{}) (fs []interface{}, is []interface{}, err error) {
	for i, v := range p {
		switch v.(type) {
		case float32, float64:
			fs = append(fs, v)
		case int, uint, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			is = append(is, v)
		default:
			err = fmt.Errorf("%v is  not float or int", i)
			return
		}
	}
	return
}

func Sum(p ...interface{}) (ret float64, err error) {
	fs, is, err := splits(p...)
	if err != nil {
		return 0, err
	}
	ret, _ = SumFloat(fs...)
	iRet, _ := SumInt(is...)
	ret += float64(iRet)
	return
}

func Max(p ...interface{}) (ret float64, err error) {
	fs, is, err := splits(p...)
	if err != nil {
		return 0, err
	}
	ret, _ = MaxFloat(fs...)
	iRet, _ := MaxInt(is...)
	if fIRet := float64(iRet); ret < fIRet {
		ret = fIRet
	}
	return
}

func Min(p ...interface{}) (ret float64, err error) {
	fs, is, err := splits(p...)
	if err != nil {
		return 0, err
	}
	ret, _ = MinFloat(fs...)
	iRet, _ := MinInt(is...)
	if fIRet := float64(iRet); ret > fIRet {
		ret = fIRet
	}
	return
}
