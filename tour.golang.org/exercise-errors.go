package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x >= 0 {
        z := 1.0
        for c := 0; c < 10; c += 1 {
            z -= (z*z - x) / (2*z)
        }
        return z, nil
    }

    return 0, ErrNegativeSqrt(x)
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
    fmt.Println(ErrNegativeSqrt(-2).Error())
}
