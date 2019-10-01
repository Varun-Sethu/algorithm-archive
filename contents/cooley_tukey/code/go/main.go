package main

import (
	"math/cmplx"
	"math"
)



func dft(x []float64) []complex128 {
	N := float64(len(x))
	res := []complex128{}

	for n := 0.0; n < N; n++ {
		for k := 0.0; k < N; k++ {
			shiftFreq := complex(k*n/N, 0)

			res[int(n)] += cmplx.Exp(-2.0i*math.Pi*shiftFreq)
		}
	}

	return res
} 


func cooley_tukey(x []float64) []complex128 {
	N := len(x)
	if N <= 1.0 {
		return []complex128{complex(x[0], 0)}
	}

	even := []float64{}
	odd  := []float64{}
	// Split into even and odd indicies
	for i, v := range x {
		if i % 2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	even_fourier := cooley_tukey(even)
	odd_fourier  := cooley_tukey(odd)

	temp := make([]complex128, N)
	for k := 0; k < int(N / 2); k++ {
		shiftFreq := complex(float32(k)/float32(N), 0)

		temp[k] = even_fourier[k] + cmplx.Exp(-2.0i*math.Pi*complex128(shiftFreq)) * odd_fourier[k]
		temp[k + int(N / 2)] = even_fourier[k] - cmplx.Exp(-2.0i*math.Pi*complex128(shiftFreq)) * odd_fourier[k]
	}

	return temp
}


func main() {

}