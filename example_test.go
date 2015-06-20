package stringset_test

import (
	"fmt"
	"sort"
	"strings"

	"github.com/AchievementNetwork/stringset"
)

func ExampleStringSet_Add() {
	greetings := stringset.New().Add("hello")
	greetings.Add("aloha", "bonjour", "g'day")
	// the Strings() function will return results in random order
	output := greetings.Strings()
	sort.Strings(output)
	fmt.Println(output)
	// Output:
	// [aloha bonjour g'day hello]
}

func ExampleStringSet_Delete() {
	nums := stringset.New().Add("1", "2", "3", "4", "5", "6", "7")
	nums.Delete("2", "4", "6")
	output := nums.Strings()
	sort.Strings(output)
	fmt.Println(output)
	// Output:
	// [1 3 5 7]
}

func ExampleStringSet_Intersection() {
	ss1 := stringset.New().Add("a", "b", "c", "d")
	ss2 := stringset.New().Add("c", "d", "e", "f")

	inter := ss1.Intersection(ss2)
	output := inter.Strings()
	sort.Strings(output)
	fmt.Println(output)
	// Output:
	// [c d]
}

func ExampleStringSet_Negate() {
	ss1 := stringset.New().Add("a", "b", "c")
	ss2 := stringset.New().Add("c", "d", "e").Negate()

	// set 2 is negative so its elements will be deleted from set 1
	inter1 := ss1.Intersection(ss2)
	output1 := inter1.Strings()
	sort.Strings(output1)
	fmt.Println(output1)

	ss1.Negate()
	ss2.Negate()

	// set 1 is negative so its elements will be deleted from set 2
	inter2 := ss1.Intersection(ss2)
	output2 := inter2.Strings()
	sort.Strings(output2)
	fmt.Println(output2)

	// Output:
	// [a b]
	// [d e]
}

func ExampleRemoveDuplicates() {
	s := "this is a test it is only a test"
	nodups := stringset.New().Add(strings.Split(s, " ")...).Strings()
	sort.Strings(nodups)
	fmt.Println(nodups)
	// Output:
	// [a is it only test this]
}
