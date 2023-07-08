package main

type Shoe int

const (
	Tennis Shoe = iota
	Dress
	Sandal
	Clog
)

// Get assigned int values 0, 1, 2, 3.
// Can also do _ Shoe = iota to get 1, 2, 3, 4.

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FalgPointToPoint
	FlagMulticast
)

// Get powers of two, so easy to combine flags by adding binary numbers.
