package antrian

type Antrian struct {
	NamaDokter string
	NoAntrian  int
	Name       string
	Keluhan    string
}

type LinkedlistAntrian struct {
	Data Antrian
	Next *LinkedlistAntrian
}
