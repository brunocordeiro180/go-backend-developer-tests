package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {

	checkResults := func(t testing.TB, got, want []string) {
		t.Helper()

		if len(got) != len(want) {
			t.Errorf("got %d words, but wanted %d", len(got), len(want))
		}

		for i, g := range got {
			if want[i] != g {
				t.Errorf("got %v words, but wanted %v", got, want)
			}
		}
	}

	t.Run("Just Fizz", func(t *testing.T) {
		want := []string{"Fizz"}
		got := FizzBuzz(1, 1, 10)

		checkResults(t, got, want)
	})

	t.Run("Just Buzz", func(t *testing.T) {
		want := []string{"Buzz"}
		got := FizzBuzz(1, 10, 1)

		checkResults(t, got, want)
	})

	t.Run("Fizz and Buzz", func(t *testing.T) {
		want := []string{"1", "2", "Fizz", "4", "Buzz"}
		got := FizzBuzz(5, 3, 5)

		checkResults(t, got, want)
	})
}
