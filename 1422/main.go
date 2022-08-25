package _422

func main() {

}

func maxScore(s string) int {
	max := 0
	n := len(s)

	for i := 1; i < n; i++ {
		score := 0
		for j := 0; j < i; j++ {
			if s[j] == '0' {
				score++
			}
		}
		for j := i; j < n; j++ {
			if s[j] == '1' {
				score++
			}
		}

		if max < score {
			max = score
		}
	}
	return max
}
