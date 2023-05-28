package suster

type Suster struct {
	Id    int
	Nama  string
	Tlp   string
	Shift string
	Tugas string
}
type LinkedlistSuster struct {
	Data Suster
	Next *LinkedlistSuster
}
