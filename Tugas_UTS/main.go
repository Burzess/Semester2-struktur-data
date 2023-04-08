package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"tugas-uts/entities"
	"tugas-uts/views"
	"tugas-uts/views/dosenview"
	"tugas-uts/views/mhsview"
)

func clear() {
	// hapus isi layar konsol
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	dataDsn := entities.GerbongDsn{}
	dataMhs := entities.GerbongMhs{}
	var menuUtama int = 0

	for menuUtama != 3 { // jika pilihan menuUtama == 3 maka keluar dari program
		clear()
		views.MenuUtama()
		fmt.Scan(&menuUtama)
		scanner.Scan()

		switch menuUtama {
		case 1:
			var input int = 0
			clear()
			for input != 6 { // jika pilihan input == 6 maka kembali ke menu utama
				views.SubMenu() 
				fmt.Scan(&input)

				switch input {
				case 1:
					dosenview.MenuInsertDosen(&dataDsn)
				case 2:
					dosenview.MenuUpdateDosen(&dataDsn)
				case 3:
					dosenview.MenuDeleteDosen(&dataDsn)
				case 4:
					dosenview.MenuViewAll(&dataDsn)
				case 5:
					dosenview.MenuViewByNip(&dataDsn)
				default:
					fmt.Println("pilihan tidak valid")
				}
				if input == 6 {
					continue
				}
			}
		case 2:
			var input int
			for input != 6 {
				views.SubMenu()
				fmt.Print("masukkan pilihan menu :")
				fmt.Scan(&input)
				scanner.Scan()

				switch input {
				case 1:
					mhsview.MenuInsertMhs(&dataMhs)
				case 2:
					mhsview.MenuUpdatetMhs(&dataMhs)
				case 3:
					mhsview.MenuDeletetMhs(&dataMhs)
				case 4:
					mhsview.MenuViewAllMhs(&dataMhs)
				case 5:
					mhsview.MenuViewByNpmMhs(&dataMhs)
				default:
					if input == 6 {
						continue
					}
					fmt.Println("pilihan tidak valid")
				}
			}
		default:
			if menuUtama == 3 {
				break
			}
			fmt.Println("pilihan tidak valid")
		}
	}

}
