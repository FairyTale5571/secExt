package files

import "testing"

func BenchmarkWriteFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := WriteFile("~\\AppData\\Local\\Steam\\htmlcache\\test.db", "test"); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := ReadFile("~\\AppData\\Local\\Steam\\htmlcache\\test.db"); err != nil {
			b.Error(err)
		}
	}
}
