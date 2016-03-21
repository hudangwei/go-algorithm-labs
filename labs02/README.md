# strings源码(go version go1.6rc2 windows/amd64)

`go
// Contains reports whether substr is within s.
// 判断s字符串中是否包含substr子串
func Contains(s, substr string) bool {
	return Index(s, substr) >= 0
}

// Index returns the index of the first instance of sep in s, or -1 if sep is not present in s.
func Index(s, sep string) int {
	n := len(sep)
	switch {
	case n == 0: //子串长度为0，也返回true
		return 0
	case n == 1: //子串长度为1，判断s是否包含byte
		return IndexByte(s, sep[0])
	case n <= shortStringLen: //const shortStringLen = 31 子串长度小于31时，排序
		return indexShortStr(s, sep)
	case n == len(s):
		if sep == s {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	}
	// 当子串长度大于31时，用Rabin-Karp算法
	// Rabin-Karp search
	hashsep, pow := hashStr(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashsep && s[:n] == sep {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashsep && s[i-n:i] == sep {
			return i - n
		}
	}
	return -1
}

`go