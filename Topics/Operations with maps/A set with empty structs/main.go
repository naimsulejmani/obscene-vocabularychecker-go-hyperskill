package main

func main() {
	// Create a set named 'mySet' with an empty struct as the value type below.
	// The set can have as many keys as you want; you can get creative!
	mySet := map[string]struct{}{}

	// Populate the set with some creative keys. Since we're using empty structs, we're essentially ignoring the values.
	mySet["a"] = struct{}{}
	mySet["b"] = struct{}{}
	mySet["c"] = struct{}{}
	mySet["d"] = struct{}{}
	mySet["e"] = struct{}{}

	checkSetTypeValues(mySet) // DO NOT delete this line!
}
