package util

func CalculateModifier(abilityScore int) (modifier int) {
	if abilityScore <= 0 || abilityScore >= 30 {
		return -1
	}

	return ((abilityScore - 10) / 2)
}
