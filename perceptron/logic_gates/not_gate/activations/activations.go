package activations

import "math"

func  HeavisideFunc(x float32) int32 {
	if x < 0 {
		return 0
	}
	return 1
}

func  SignumFunc(x float32) float64{
	return 1.0 / (1.0 + math.Exp(-float64(x)))
}
