package utils

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func ContainsStr(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ContainsInt(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}

	return false
}
