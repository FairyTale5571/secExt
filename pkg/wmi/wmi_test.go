package wmi

import "testing"

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
	}
}

func BenchmarkGetCpuId(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetCpuId()
	}
}

func BenchmarkGetCpuName(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetCpuName()
	}
}

func BenchmarkGetMotherId(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetMotherId()
	}
}

func BenchmarkGetMotherName(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetMotherName()
	}
}

func BenchmarkWmi_GetRamSerialNumber(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetRamSerialNumber()
	}
}

func BenchmarkWmi_GetRamPartNumber(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetRamPartNumber()
	}
}

func BenchmarkWmi_GetRamName(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetRamName()
	}
}

func BenchmarkWmi_GetRamCapacity(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetRamCapacity()
	}
}

func BenchmarkWmi_GetProductId(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetProductId()
	}
}

func BenchmarkWmi_GetProductInstallDate(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetProductInstallDate()
	}
}

func BenchmarkWmi_GetProductVersion(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetProductVersion()
	}
}

func BenchmarkWmi_GetBiosId(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetBiosId()
	}
}

func BenchmarkWmi_GetBiosReleaseDate(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetBiosReleaseDate()
	}
}

func BenchmarkWmi_GetBiosVersion(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetBiosVersion()
	}
}

func BenchmarkWmi_GetPcName(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetPcName()
	}
}

func BenchmarkWmi_GetSID(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetSID()
	}
}

func BenchmarkWmi_GetVRAM(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetVRAM()
	}
}

func BenchmarkWmi_GetCSP(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetCSP()
	}
}

func BenchmarkWmi_GetDiskDrives(b *testing.B) {
	w := New()
	for i := 0; i < b.N; i++ {
		w.GetDiskDrives()
	}
}
