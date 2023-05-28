package suster

type Suster struct {
	Id    int
	Nama  string
	Tlp   string
	Shift string
}
type LinkedlistSuster struct {
	Data Suster
	Next *LinkedlistSuster
}
