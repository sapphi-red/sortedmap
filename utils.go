package sortedmap

func insertAt[T any](slice []T, pos int, v T) []T {
	slice = append(slice, v)
	copy(slice[pos+1:], slice[pos:])
	slice[pos] = v
	return slice
}

func deleteAt[T any](slice []T, pos int) []T {
	return append(slice[:pos], slice[pos+1:]...)
}
