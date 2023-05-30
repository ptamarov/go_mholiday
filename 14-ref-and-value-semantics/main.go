package main

import "fmt"

type User struct {
	Name  string
	Count int
}

func addTo(u *User) {
	u.Count++
}

func main() {
	stalePointerExample()
	pointerInsideForLoop()
}

func stalePointerExample() {
	var users []User

	users = []User{{"Alice", 0}, {"Bob", 2}}
	alice := &users[0] // Bad idea.
	amy := User{"Amy", 1}
	users = append(users, amy)

	addTo(alice)       // Stale pointer, slice had to change due to append.
	fmt.Println(users) // [{Alice 0} {Bob 2} {Amy 1}]
}

func pointerInsideForLoop() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}
	b := [][]byte{}

	for _, item := range items {
		a = append(a, item[:]) // Bad. Passing pointer to loop variable item, will only keep last value.
	}

	fmt.Println(a) // [[5 6] [5 6] [5 6]]

	for _, item := range items {
		i := make([]byte, len(item)) // Good. Make a local variable, do not use the loop variable...
		copy(i, item[:])             // ...and copy the data there.
		b = append(b, i)
	}

	fmt.Println(b) // [[1 2] [3 4] [5 6]]

}
