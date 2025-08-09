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

func FindLongestPalindromeExpandAroundCenter(s string) string {
	if len(s) == 0 {
		return ""
	}

	start := 0
	maxLength := 1
	length := len(s)

	for i := 0; i < length; i++ {
		// Проверяем палиндромы нечетной длины (центр в символе)
		len1 := expandAroundCenter(s, i, i)

		// Проверяем палиндромы четной длины (центр между символами)
		len2 := expandAroundCenter(s, i, i+1)

		// Берем максимальную длину из двух вариантов
		currentMaxLength := max(len1, len2)

		// Обновляем результат, если найден более длинный палиндром
		if currentMaxLength > maxLength {
			maxLength = currentMaxLength
			// Вычисляем начальную позицию палиндрома
			start = i - (currentMaxLength-1)/2
		}
	}

	return s[start : start+maxLength]
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

func expandAroundCenter(s string, left, right int) int {
	length := len(s)

	// Расширяемся пока символы совпадают и не выходим за границы
	for left >= 0 && right < length && s[left] == s[right] {
		left--
		right++
	}

	// Возвращаем длину найденного палиндрома
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
