package cleaninput

func CleanInput(input string) []string {
	result := []string{}

	temp := ""
	for i, r := range input {
		if r == ' ' {
			if temp != "" {
				result = append(result, temp)
				temp = ""
			}
		} else {
			temp += string(r)
		}

		if l := i + 1; len(input) == l && temp != "" {
			result = append(result, temp)
		}
	}
	return result
}
