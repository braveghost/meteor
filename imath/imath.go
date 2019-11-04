package imath

// 取余数和商
func Divmod(numerator, denominator int64) (int64, int64) {
	return numerator / denominator, numerator % denominator
}
