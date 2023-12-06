package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var ans uint32
	for i := 0; i < 32; i++ {
		// num向右移一位，再和1做与运算，如果(向右移一位后的)num最后一位是0，0和1做与运算还是0，如果最后一位是1，1和1做与运算为1，这样就得到了此时num的最后一位
		k := (num >> i) & 1
		// ans向左移一位，后面补0
		// 0和k做或运算, k 为1时，ans最后一位的0变为1，k为0时，ans最后一位的0变为0
		ans = (ans << 1) | k
	}

	return ans
}

func main() {

	//num := uint32(0b00000010100101000001111010011100)
	num := uint32(0b10000000000000000000000000000000)

	res := reverseBits(num)

	fmt.Printf("res: %v\n", res)
}
