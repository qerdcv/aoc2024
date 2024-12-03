package day3

func ResolvePartOne(bytes []byte) int {
	return parseResult(bytes, false)
}

func ResolvePartTwo(bytes []byte) int {
	return parseResult(bytes, true)
}

func parseResult(bytes []byte, doDont bool) int {
	total := 0
	do := true
	for i := 0; i+8 < len(bytes); i++ {
		if doDont {
			if string(bytes[i:i+4]) == "do()" {
				do = true
				continue
			}

			if !do {
				continue
			}

			if string(bytes[i:i+7]) == "don't()" {
				do = false
				continue
			}
		}

		if string(bytes[i:i+4]) == "mul(" {
			offset := i + 4
			nI := 0
			num := 0
			n1 := 0
			for j := 0; j+offset < len(bytes); j++ {
				b := bytes[j+offset]
				if b >= '0' && b <= '9' {
					num = (num * 10) + int(b-'0')
				} else if b == ',' {
					n1 = num
					num = 0
					nI++
					if nI > 1 {
						break
					}
				} else if b == ')' {
					total += n1 * num
					break
				} else {
					break
				}
			}

		}
	}

	return total
}
