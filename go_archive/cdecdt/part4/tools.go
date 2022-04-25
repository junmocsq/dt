package part4

import "errors"

func Compare(a, b interface{}) (int, error) {
	switch a := a.(type) {
	case int:
		if b, ok := b.(int); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int8:
		if b, ok := b.(int8); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int16:
		if b, ok := b.(int16); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int32:
		if b, ok := b.(int32); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case int64:
		if b, ok := b.(int64); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint:
		if b, ok := b.(uint); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint8:
		if b, ok := b.(uint8); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint16:
		if b, ok := b.(uint16); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint32:
		if b, ok := b.(uint32); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case uint64:
		if b, ok := b.(uint64); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case string:
		if b, ok := b.(string); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case float32:
		if b, ok := b.(float32); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	case float64:
		if b, ok := b.(float64); ok {
			if a > b {
				return 1, nil
			} else if a == b {
				return 0, nil
			} else {
				return -1, nil
			}
		}
	}
	return 0, errors.New("数据类型不一样或不是支持的类型")
}
