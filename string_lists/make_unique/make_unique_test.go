package make_unique

import "fmt"

func ExampleMakeUnique() {
	s1 := []string{"a", "b", "a", "c", "b", "a"}

	MakeUnique(s1)
	fmt.Println(s1)

	// Output: [a b a_2 c b_2 a_3]
}
