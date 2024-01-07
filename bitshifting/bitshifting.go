package bitshifting

func MultiplyByTwo(x int) int {
	return x << 1
}

func DivideByTwo(x int) int {
	return x >> 1
}

func ToPower(x int, y int) int {
	return x << uint(y)
}

func ToRoot(x int, y int) int {
	return x >> uint(y)
}

func IsEven(x int) bool {
	return x&1 == 0
}

func IsOdd(x int) bool {
	return x&1 == 1
}

func IsPowerOfTwo(x int) bool {
	return x != 0 && x&(x-1) == 0
}

func Swap(x int, y int) (int, int) {
	x = x ^ y
	print(x, "\n")
	print(y, "\n")
	y = x ^ y
	print(x, "\n")
	print(y, "\n")
	x = x ^ y
	print(x, "\n")
	print(y, "\n")
	return x, y
}
