package main

import (
	"reflect"
	"testing"
)

//func TestSum(t *testing.T) {
//	numbers := [5]int{1, 2, 3, 4, 5}
//
//	got := Sum(numbers)
//	want := 15
//
//	if got != want {
//		t.Errorf("got %d want %d, %v", got, want, numbers)
//	}
//}

// 在本例中，针对该函数写两个测试其实是多余的，因为切片尺寸并不影响函数的运行。
// Go 有内置的计算测试 覆盖率的工具，它能帮助你发现没有被测试过的区域。
// 我们不需要追求 100% 的测试覆盖率，它只是一个供你获取测试覆盖率的方式。只要你严格遵循 TDD 规范，那你的测试覆盖率就会很接近 100%。
// 运行： go test -cover 可以查看测试覆盖率
func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	//t.Run("collection of any size", func(t *testing.T) {
	//	numbers := []int{1, 2, 3}
	//
	//	got := Sum(numbers)
	//	want := 6
	//
	//	if got != want {
	//		t.Errorf("got %d want %d given, %v", got, want, numbers)
	//	}
	//})

}

func checkSums(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	//want := "bob"   // 此处使用bob可以编译成功，存在类型安全问题；将下面的reflect代码抽取为单独函数，可以避免类型安全问题

	//if !reflect.DeepEqual(got, want) {
	//	t.Errorf("got %v want %v", got, want)
	//}
	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {

	//checkSums := func(t *testing.T, got, want []int) {
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("got %v want %v", got, want)
	//	}
	//}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}
