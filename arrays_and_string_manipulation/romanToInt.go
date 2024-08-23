package arrays_and_string_manipulation

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); {
		// check if we're in a substraction case
		if i < len(s)-1 {
			switch s[i : i+2] {
			case "IV":
				sum += 4
				i += 2
				continue
			case "IX":
				sum += 9
				i += 2
				continue
			case "XL":
				sum += 40
				i += 2
				continue
			case "XC":
				sum += 90
				i += 2
				continue
			case "CD":
				sum += 400
				i += 2
				continue
			case "CM":
				sum += 900
				i += 2
				continue
			}
		}

		switch s[i] {
		case 'I':
			sum += 1
		case 'V':
			sum += 5
		case 'X':
			sum += 10
		case 'L':
			sum += 50
		case 'C':
			sum += 100
		case 'D':
			sum += 500
		case 'M':
			sum += 1000
		}
		i++
	}
	return sum
}
