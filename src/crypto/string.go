package crypto

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789._-"
	letterBytes2 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	letterIdxBits = 7 // 7 bits to represent a Letter index
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func String(n int) string {
	b := make([]byte, n)
	l := len(letterBytes)
	// A randpool.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, randpool.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randpool.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < l {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func StringReadable(n int) string {
	b := make([]byte, n)
	l := len(letterBytes2)
	for i, cache, remain := n-1, randpool.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randpool.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < l {
			b[i] = letterBytes2[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
