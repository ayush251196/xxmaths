package calculate

import (
	"fmt"
	"strings"
	"testing"
)

func TestFibonnaci(t *testing.T) {
	want := "0 1 1 2"
	got := strings.TrimSpace(Fibonnaci(4))
	printErrorMessage(got, want, t)
}

func TestSum(t *testing.T) {
	want := fmt.Sprint(6)
	got := fmt.Sprint(Sum(5, 1))
	printErrorMessage(got, want, t)

}

func TestSumX(t *testing.T) {
	want := fmt.Sprint(10)
	sum, _ := SumX(1, 2, 3, 4)
	got := fmt.Sprint(sum)
	printErrorMessage(got, want, t)
}

func printErrorMessage(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("got =%q and want=%q", got, want)
	}
}
