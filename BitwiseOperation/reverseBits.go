package main

func reverseBits(n uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res << 1
		bit := n % 2
		res += bit
		n = n >> 1
	}

	return res

}
