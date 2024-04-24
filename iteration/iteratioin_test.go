package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// 基准测试
// testing.B 可使你访问隐性命名（cryptically named）b.N。
// 基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
// 代码运行的次数不会对你产生影响，测试框架会选择一个它所认为的最佳值，以便让你获得更合理的结果。
// 用 go test -bench=. 来运行基准测试。 (如果在 Windows Powershell 环境下使用 go test -bench=".")
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
