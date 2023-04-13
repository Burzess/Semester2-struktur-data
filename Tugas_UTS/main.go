package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"tugas-uts/views"
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

				case 2:

				case 3:

				case 4:

				case 5:

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

				case 2:

				case 3:

				case 4:

				case 5:

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

				case 2:

				case 3:

				case 4:

				case 5:

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
