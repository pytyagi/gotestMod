package prime

func CheckPrime(n int64) bool {
	if n <= 1 {
		return false
	}

	var i int64
	for i = 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
