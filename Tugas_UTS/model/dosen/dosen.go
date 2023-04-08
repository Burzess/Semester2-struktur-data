package dosen

import (
	"fmt"
	"os"
	"tugas-uts/entities"

	"github.com/olekukonko/tablewriter"
)

func InsertDataDosen(g *entities.GerbongDsn, data entities.Dosen) {
	newGerbong := &entities.GerbongDsn{Dosen: data}

	temp := g
	if temp.Next == nil {
		temp.Next = newGerbong
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = newGerbong
	}
}

func Update(g *entities.GerbongDsn, nip string) {
	temp := g

	if temp.Next == nil {
		fmt.Println("Data Kosong")
		return
	} else {
		for temp.Next != nil {
			if temp.Next.Dosen.Nip == nip {
				fmt.Print("Masukan nama baru : ")
				fmt.Scan(&temp.Next.Dosen.Nama)
				return
			}
			temp = temp.Next
		}
	}

	fmt.Printf("data dengan nip %s  tidak ada\n", nip)
}

func Delete(g *entities.GerbongDsn, nip string) {
	temp := g

	if temp.Next == nil {
		fmt.Println("Data Kosong")
		return
	}
	// mengecek apakah data yang di hapus data pertama
	if temp.Dosen.Nip == nip {
		*temp = *temp.Next
		return
	}

	for temp.Next != nil {
		if temp.Next.Dosen.Nip == nip {
			temp.Next = temp.Next.Next
			return
		}
		temp = temp.Next
	}

	fmt.Println("Delete gagal, ID tidak tersedia")
}

// Fungsi untuk mendapatkan daftar NIP dan Nama dosen pada suatu gerbong
func GetList(g *entities.GerbongDsn) {
	// Inisialisasi tabel dan header kolom
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NPM", "Nama"})

	// Iterasi melalui setiap elemen GerbongMhs dan menambahkan data ke dalam tabel
	current := g
	for current.Next != nil {
		dsn := current.Next.Dosen
		row := []string{dsn.Nip, dsn.Nama}
		table.Append(row)
		current = current.Next
	}

	// Render tabel ke terminal
	table.Render()
}

func GetListByNip(g *entities.GerbongDsn, nip string) {
	// Deklarasi variabel data sebagai slice of slice yang akan digunakan sebagai data tabel
	data := [][]string{}
	// Deklarasi variabel curGerbong sebagai pointer ke struct GerbongDsn
	curGerbong := g
	// Looping untuk mencari dosen dengan NIP yang sesuai dengan parameter nip pada setiap gerbong
	for curGerbong.Next != nil {
		if curGerbong.Next.Dosen.Nip == nip {
			// Jika dosen dengan NIP yang sesuai ditemukan, tambahkan informasinya pada variabel data
			data = append(data, []string{curGerbong.Dosen.Nip, curGerbong.Dosen.Nama})
			break
		}
		curGerbong = curGerbong.Next
	}

	if len(data) == 0 { // Jika tidak ditemukan dosen dengan NIP yang sesuai, tampilkan pesan error
		fmt.Println("Dosen dengan NIP", nip, "tidak ditemukan.")
	} else { // Jika ditemukan, tampilkan informasinya dalam bentuk tabel
		// Deklarasi variabel table sebagai tablewriter yang akan digunakan untuk menghasilkan output dalam bentuk tabel
		table := tablewriter.NewWriter(os.Stdout)
		// Mengatur header tabel
		table.SetHeader([]string{"NIP", "Nama"})
		// Mengatur border tabel
		table.SetBorder(false)
		// Mengatur alignment isi data pada tabel ke kiri
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		// Menambahkan data pada tabel
		table.AppendBulk(data)
		// Mengatur output dari tabel ke console
		table.Render()
	}
}
