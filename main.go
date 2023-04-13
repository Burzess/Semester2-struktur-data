package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"tugas-uts/entities"
	"tugas-uts/views"
	"tugas-uts/views/dokterview"
)

func clear() {
	// hapus isi layar konsol
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var menuUtama int = 0
	var dataDokter entities.Dokter

	for menuUtama != 4 {
		clear()
		views.MenuUtama()
		fmt.Scan(&menuUtama)
		scanner.Scan()

		switch menuUtama {
		case 1:
			var input int = 0
			clear()
			for input != 6 {
				views.SubMenu()
				fmt.Scan(&input)

				switch input {
				case 1:
					fmt.Println("MENU INSERT")
					dokterview.MenuInsert(dataDokter)
				case 2:
					fmt.Println("MENU UPDATE")
				case 3:
					fmt.Println("MENU DELETE")
				case 4:
					fmt.Println("MENU VIEW ALL")
				case 5:
					fmt.Println("MENU VIEW BY ID")
				default:
					if input == 6 {
						continue
					}
					fmt.Println("pilihan tidak valid")
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
					fmt.Println("MENU INSERT")
				case 2:
					fmt.Println("MENU UPDATE")
				case 3:
					fmt.Println("MENU DELETE")
				case 4:
					fmt.Println("MENU VIEW ALL")
				case 5:
					fmt.Println("MENU VIEW BY ID")
				default:
					if input == 6 {
						continue
					}
					fmt.Println("pilihan tidak valid")
				}
			}
		case 3:
			var input int
			for input != 6 {
				views.SubMenu()
				fmt.Print("masukkan pilihan menu :")
				fmt.Scan(&input)
				scanner.Scan()

				switch input {
				case 1:
					fmt.Println("MENU INSERT")
				case 2:
					fmt.Println("MENU UPDATE")
				case 3:
					fmt.Println("MENU DELETE")
				case 4:
					fmt.Println("MENU VIEW ALL")
				case 5:
					fmt.Println("MENU VIEW BY ID")
				default:
					if input == 6 {
						continue
					}
					fmt.Println("pilihan tidak valid")
				}
			}
		default:
			if menuUtama == 4 {
				break
			}
			fmt.Println("pilihan tidak valid")
		}
	}

}
