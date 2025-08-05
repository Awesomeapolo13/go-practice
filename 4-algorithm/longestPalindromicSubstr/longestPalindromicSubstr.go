package longestPalindromicSubstr

func FindLongestPalindrome(s string) string {
	var maxLength int
	str := s
	includeLastSymbol := true
	length := len(s)
	i, j := 0, length-1
	if length == 1 {
		return str
	}

	palindromes := map[int]string{}
	for i < length {
		if isPalindrome(str) {
			substrLength := len(str)
			id, _ := palindromes[substrLength]
			if id == "" && substrLength > maxLength {
				maxLength = substrLength
			}
			if id == "" {
				palindromes[substrLength] = str
			}
		}
		if !includeLastSymbol {
			j--
		}
		if i == j {
			i++
			j = length - 1
			str = s[i:]
			includeLastSymbol = true
		} else {
			str = s[i:j]
			includeLastSymbol = false
		}
	}

	return palindromes[maxLength]
}

func isPalindrome(str string) bool {
	i, j := 0, len(str)-1
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func isPalindromeWithExpandAroundCenter(str string) bool {
	length := len(str)
	i, j := 0, length-1
	midIdx := i + (j-i)/2
	if (length % 2) == 0 {
		i = midIdx
		j = midIdx
	} else {
		i = midIdx
		j = i + 1
	}

	for i >= 0 && j != length {
		if str[i] != str[j] {
			return false
		}
		i--
		j++
	}

	return true
}
