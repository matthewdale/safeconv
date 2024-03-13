// Copied from https://github.com/g-utils/overflow/blob/b0c53d3efc6f4599ff58a858bbf9d91ec7a1a359/overflow_test.go

package safeconv_test

import (
	"math"
	"testing"

	"github.com/matthewdale/safeconv"
)

// sample all possibilities of 8 bit numbers
// by checking against 64 bit numbers

func TestAlgorithms(t *testing.T) {

	errors := 0

	for a64 := int64(math.MinInt8); a64 <= int64(math.MaxInt8); a64++ {

		for b64 := int64(math.MinInt8); b64 <= int64(math.MaxInt8) && errors < 10; b64++ {

			a8 := int8(a64)
			b8 := int8(b64)

			if int64(a8) != a64 || int64(b8) != b64 {
				t.Fatal("LOGIC FAILURE IN TEST")
			}

			// ADDITION
			{
				r64 := a64 + b64

				// now the verification
				result, ok := safeconv.AddOK(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v + %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// SUBTRACTION
			{
				r64 := a64 - b64

				// now the verification
				result, ok := safeconv.SubOK(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v - %v = %v instead of %v\n",
						a8, b8, result, r64)
				}
				if !ok && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// MULTIPLICATION
			{
				r64 := a64 * b64

				// now the verification
				result, ok := safeconv.MulOK(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v * %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// DIVISION
			if b8 != 0 {
				r64 := a64 / b64

				// now the verification
				result, ok := safeconv.DivOK(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v / %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && result != 0 && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}
		}
	}

}

func TestUnsignedAlgorithms(t *testing.T) {

	errors := 0

	for a64 := uint64(0); a64 <= uint64(math.MaxUint8); a64++ {

		for b64 := uint64(0); b64 <= uint64(math.MaxUint8) && errors < 10; b64++ {

			a8 := uint8(a64)
			b8 := uint8(b64)

			if uint64(a8) != a64 || uint64(b8) != b64 {
				t.Fatal("LOGIC FAILURE IN TEST")
			}

			// ADDITION
			{
				r64 := a64 + b64

				// now the verification
				result, ok := safeconv.AddOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v + %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// SUBTRACTION
			{
				r64 := a64 - b64

				// now the verification
				result, ok := safeconv.SubOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v - %v = %v instead of %v\n",
						a8, b8, result, r64)
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// MULTIPLICATION
			{
				r64 := a64 * b64

				// now the verification
				result, ok := safeconv.MulOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v * %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// DIVISION
			if b8 != 0 {
				r64 := a64 / b64

				// now the verification
				result, ok := safeconv.DivOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v / %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && result != 0 && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}
		}
	}
}

// Second set of tests for unsigned overflows in case my implementation is incorrect
// Source: github.com/rwxe/overflow
func TestAlgorithmsUintImpl2(t *testing.T) {

	errors := 0

	for a64 := uint64(0); a64 <= uint64(math.MaxUint8); a64++ {

		for b64 := uint64(0); b64 <= uint64(math.MaxUint8) && errors < 10; b64++ {

			a8 := uint8(a64)
			b8 := uint8(b64)

			if uint64(a8) != a64 || uint64(b8) != b64 {
				t.Fatal("LOGIC FAILURE IN TEST")
			}

			// ADDITION
			{
				r64 := a64 + b64

				// now the verification
				result, ok := safeconv.AddOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v + %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// SUBTRACTION
			{
				r64 := a64 - b64

				// now the verification
				result, ok := safeconv.SubOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v - %v = %v instead of %v\n",
						a8, b8, result, r64)
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// MULTIPLICATION
			{
				r64 := a64 * b64

				// now the verification
				result, ok := safeconv.MulOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v * %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// DIVISION
			if b8 != 0 {
				r64 := a64 / b64

				// now the verifiggcation
				result, ok := safeconv.DivOK(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v / %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && result != 0 && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}
		}
	}

}
