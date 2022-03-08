package modul2_testing

type mahasiswa struct {
	nama string
	nim  string
}

type sekolah struct {
	nama      string
	daftarMhs []mahasiswa
}

type School sekolah

func (s *School) TambahMahasiswa() {
	s.daftarMhs = append(s.daftarMhs, mahasiswa{
		nama: "maba 1",
		nim:  "001",
	})
}

func (s *School) GetJumlahDaftarMhs() int {
	return len(s.daftarMhs)
}

func SayHi(who string) string {
	return "hi," + who
}
