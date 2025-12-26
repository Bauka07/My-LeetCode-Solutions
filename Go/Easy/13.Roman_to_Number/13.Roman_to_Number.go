package main

func romanToInt(s string) int {
  maps := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0
	for i := 0; i < len(s); i++  {
		symbol := string(s[i])
		value := maps[symbol]

		if i+1 < len(s) && value < maps[string(s[i+1])] {
			total -= value
		} else {
			total += value
		}
	}

	return total
}
