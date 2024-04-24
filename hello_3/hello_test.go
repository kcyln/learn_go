package main

import "testing"

// TDD 测试驱动开发，或成为 敏捷开发 中的一项核心实践和技术，也是一种设计方法论。
// 旨在 明确要开发的功能后，先编写测试代码，再编写功能代码然后用测试代码验证，
// 确保测试代码通过，再修改功能代码，再重新测试，直到测试通过为止。
// 如此循环，直到完成全部功能的开发。

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
