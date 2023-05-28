package dokter

type Dokter struct {
	Id        int
	Nama      string
	Tlp       string
	JamKerja  string
	Spesialis string
}
type LinkedlistDokter struct {
	Data Dokter
	Next *LinkedlistDokter
}
