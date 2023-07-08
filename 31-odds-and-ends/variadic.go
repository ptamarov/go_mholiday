package main

// ! only last parameter of a function can be variadic
// for example, append is of the form append([]T, T...)

func sum(nums ...int) int {
	var total int

	for _, num := range nums {
		total += num
	}

	return total
}
