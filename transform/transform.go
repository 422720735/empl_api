package transform

import (
	"errors"
	"fmt"
	"strconv"
)

// 接口转string
func InterToString(inter interface{}) (string, error) {
	if inter == nil {
		return "", errors.New("参数错误")
	}
	switch v := inter.(type) {
	case []byte:
		return string(v), nil
	case int:
		str := strconv.Itoa(v)
		return str, nil
	case string:
		return v, nil
	case float64:
		str := strconv.FormatFloat(v, 'f', -1, 64)
		return str, nil
	default:
		return "", fmt.Errorf("type %T", v)
	}
}

// 接口转int
func InterToInt(inter interface{}) (int, error) {
	if inter == nil {
		return -1, errors.New("参数错误")
	}
	switch v := inter.(type) { //v表示b1 接口转换成Bag对象的值
	case []byte:
		return strconv.Atoi(string(v))
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	case float64:
		return int(v), nil
	default:
		return -1, fmt.Errorf("type %T", v)
	}
}
