package vote

func HasToVote(age int) bool {
	if age >= 18 && age <= 69 {
		return true
	}

	return false
}
