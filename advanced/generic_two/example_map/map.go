package _map

func Map[T, M any](a []T, f func(b T) M) []M {
	rst := make([]M, len(a))
	for index, value := range a {
		rst[index] = f(value)
	}
	return rst
}
