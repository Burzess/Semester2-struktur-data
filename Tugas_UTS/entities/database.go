package entities

type Dosen struct {
	Nip  string
	Nama string
}

type GerbongDsn struct {
	Dosen Dosen
	Next  *GerbongDsn
}

type Mahasiswa struct {
	Jurusan string
	Npm     string
	Nama    string
	Ipk     float64
}

type GerbongMhs struct {
	Mahasiswa Mahasiswa
	Next      *GerbongMhs
}
