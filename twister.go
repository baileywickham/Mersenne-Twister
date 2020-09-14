package main

// Magic values taken from python source code
// https://github.com/python/cpython/blob/457d4e97de0369bc786e363cb53c7ef3276fdfcd/Modules/_randommodule.c#L75

// Step 0 in MT
const N = 624
const M = 397

// b and c represent "tempering masks" defined in MT
const b = 0x9d2c5680
const c = 0xefc60000

var UPPER_MASK uint = 0x80000000 // U in MT
var LOWER_MASK uint = 0x7fffffff // ll in MT
var MATRIX_A uint = 0x9908b0df   // a in MT
var MAG01 [2]uint
var mt [N]uint

func random(i int) uint {
	var y uint
	var kk uint
	// Sets matrix MT to random value
	initMT(19650218)

	// Used in XOR, defined to make & easier.
	MAG01 = [...]uint{0x0, MATRIX_A}
	// Step 2 in MT.
	for kk = 0; kk < N-M; kk++ {
		y = (mt[kk] & UPPER_MASK) | (mt[kk+1] & LOWER_MASK)
		mt[kk] = mt[kk+M] ^ (y >> 1) ^ MAG01[y&uint(0x1)]
	}

	// Step 3 in MT
	for ; kk < N-1; kk++ {
		y = (mt[kk] & UPPER_MASK) | (mt[kk+1] & LOWER_MASK)
		mt[kk] = mt[int(kk)+int((M-N))] ^ (y >> 1) ^ MAG01[y&uint(0x1)]
	}

	y = (mt[N-1] & UPPER_MASK) | (mt[0] & LOWER_MASK)
	mt[N-1] = mt[M-1] ^ (y >> 1) ^ MAG01[y&0x1]

	// This should cycle through MT, not be assigned at 0
	y = mt[i]
	y ^= (y >> UPPER_MASK)
	y ^= (y << 7) & b
	y ^= (y << 15) & c
	y ^= (y >> 18)
	return y
}

func initMT(seed uint) {
	// Step 1 in MT. Set inital values of x
	var i uint
	mt[0] = seed
	for i = 1; i < N; i++ {
		mt[i] = (uint(1812433253)*(mt[i-1]^(mt[i-1]>>30)) + i)
	}
}

func main() {
	println(random(2))
}
