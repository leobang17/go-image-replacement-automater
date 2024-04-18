package atomicity

import (
	"errors"
	"math/rand"
	"testing"
)

func Test_AtomicSection_Success(t *testing.T) {
	randomInts := []int {}
	for i := 0; i < 20; i ++ {
		randomInts = append(randomInts, 100_000 + rand.Intn(1_000_000))
	}
	t.Log(randomInts)

	for _, randNum := range randomInts {
		t.Run("",func(t *testing.T) {
			t.Parallel()
			atomicSection := NewAtomicSection[int]()
			N := randNum
			result, err := atomicSection(func (runnable Runnable[int]) {
				t.Log(N)
				for j := 0; j < N; j ++ {
					runnable(func () (int, error) {
						return j, nil
					})
				}
			})

			if err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if len(result) != N {
				t.Errorf("expected %v length of array, but got %v length", N, len(result))
			}
			t.Logf("Actual Length: %v, Expected Length: %v", len(result), N)
		})
	}
}

func Test_AtomicSection_Failure(t *testing.T) {
	randomInts := []int {}
	for i := 0; i < 20; i ++ {
		randomInts = append(randomInts, 100_000 + rand.Intn(1_000_000))
	}
	t.Log(randomInts)

	for _, randNum := range randomInts {
		t.Run("",func(t *testing.T) {
			t.Parallel()
			atomicSection := NewAtomicSection[int]()
			N := randNum
			_, err := atomicSection(func (runnable Runnable[int]) {
				t.Log(N)
				for j := 0; j < N; j ++ {
					if j == N - 100 {
						runnable(func () (int, error) {
							return 0, errors.New("Error!")
						})
						} else {
						runnable(func () (int, error) {
							return 123, nil
						})
					}
				}
			})

			if err == nil {
				t.Errorf("expected error, but got no error")
			}
		})
	}
}