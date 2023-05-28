package pasien

type Pasien struct {
	Id       int
	Nama     string
	Penyakit string
	Telpon   string
}

type LinkedlistPasien struct {
	Data Pasien
	Next *LinkedlistPasien
}
