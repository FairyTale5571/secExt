package wmi

import "testing"

func BenchmarkGetCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCPU()
	}
}

func BenchmarkGetMother(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getMother()
	}
}

func BenchmarkGetBios(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getBios()
	}
}

func BenchmarkGetRAM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRAM()
	}
}

func BenchmarkGetOS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getOS()
	}
}

func BenchmarkGetCSP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getCSP()
	}
}

func BenchmarkGetVRAM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getVRAM()
	}
}

func BenchmarkGetHDDs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getDiskDrive()
	}
}

func BenchmarkGetUserAccount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getUserAccount()
	}
}
