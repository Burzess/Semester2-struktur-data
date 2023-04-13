package dokter

import (
	"fmt"
	"os"
	"tugas-uts/Tugas_UTS/entities"

	"github.com/olekukonko/tablewriter"
)

func InsertDokter(dataDokter *entities.Dokter, id int, nama string, Tlp string, jamKerja string, spesialis string) {
	newData := entities.Dokter{
		ID:        id,
		Nama:      nama,
		Tlp:       Tlp,
		JamKerja:  jamKerja,
		Spesialis: spesialis,
	}

	temp := dataDokter
	if temp == nil {
		temp.Next = &newData
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newData
	}
}

func ViewAll(dataDokter *entities.Dokter) {
	// inisialisasi table dan kolom header
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Nama", "Telpon", "Jam Kerja", "Spesialis"})
	cur := dataDokter
	for cur.Next != nil {
		dktr := cur
		row := []string{fmt.Sprint(dktr.ID), dktr.Nama, dktr.Tlp, dktr.JamKerja, dktr.Spesialis}
		table.Append(row)
	}
	table.Rander()
}
