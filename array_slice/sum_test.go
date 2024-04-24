package main

import "testing"

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
