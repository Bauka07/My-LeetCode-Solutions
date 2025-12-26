package main

func isPalindrome(x int) bool {
	if x < 0 {
			return false
	}

	slice := []int{}
	for x > 0 {
		slice = append(slice, x%10)
		x = x / 10
	}

	reverse := []int{}
	for i := len(slice) - 1; i >= 0; i-- {
		reverse = append(reverse, slice[i])
	}

	for i := range slice {
		if slice[i] != reverse[i] {
			return false
		}
	}

	return true
}
