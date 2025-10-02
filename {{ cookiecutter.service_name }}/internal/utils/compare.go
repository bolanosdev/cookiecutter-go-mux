package utils

func IF(condition bool, a any, b any) interface{} {
	if condition {
		return a
	}
	return b
}
