import "maps"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	setS := map[string]int{}

	for _, char := range s {
		if v, ok := setS[string(char)]; ok {
			setS[string(char)] += v + 1
			continue
		}
		setS[string(char)] = 0
	}	

	setT := map[string]int{}

	for _, char := range t {
		if v, ok := setT[string(char)]; ok {
			setT[string(char)] += v + 1
			continue
		}
		setT[string(char)] = 0
	}	

	fmt.Printf("setS: %v", setS)
	fmt.Printf("setT: %v", setT)

	if !maps.Equal(setS, setT) {
		return false
	}

	return true
}
