package lists

//liefert die Postion ´von x in nums, falls x darin vorkommt
//liefert -1 falls x nicht vorkommt
//kommt x mehrfach vor wird die postio des letzten Vorkommens geliefert

func Find(nums []int, x int) int {
	if len(nums) == 0 {
		return 0
	}

	pos := -1

	for i := 1; i < len(nums); i++ {
		if nums[i] == x {
			pos = i
		}
	}

	return pos
}
