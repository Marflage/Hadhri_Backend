package utils

// func GroupBy[T any, K comparable, V comparable](slice []T, groupFn func(T) (K,V)) map[K][]T {
// 	result := make(map[K][]T)

// 	for _, item := range slice {
// 		key := groupFn(item)
// 		result[key] = append(result[key], item)
// 	}

// 	return result
// }

// func GroupBy[T any, K comparable, X any](slice []T, groupFn func(T) (K, X)) Grouping[K, []T, X] {
// 	result := []Grouping[K, []T, X]{}

// 	for _, item := range slice {
// 		key, name := groupFn(item)
// 		result[key] = append(result[key], item)
// 	}

// 	return result
// }

func GroupBy[T any, K comparable, X any](slice []T, groupFn func(T) (K, X)) map[K]GroupData[X, []T] {
	result := make(map[K]GroupData[X, []T])

	for _, item := range slice {
		key, name := groupFn(item)
		value := GroupData[X, []T]{Name: name, Data: append(result[key].Data, item)}
		result[key] = value
	}

	return result
}

type GroupData[X any, T any] struct {
	Name X
	Data T
}
