package utils

func ConvertToInterfaceArray[T any](objects []T) []interface{} {
	res := make([]interface{}, len(objects))
	for i, object := range objects {
		res[i] = object
	}
	return res
}

func ConvertArrayToAnotherType[J any](objects []any) []J {
	res := make([]J, len(objects))
	for i, object := range objects {
		res[i] = object.(J)
	}
	return res
}