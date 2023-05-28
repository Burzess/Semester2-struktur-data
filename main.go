package main

import (
	"fmt"
	"os"
	"os/exec"
	"progres/controller/usercontroller"
	"progres/model/dokter"
	"progres/model/user"
	"progres/views"
	"progres/views/antrianview"
	"progres/views/dokterview"
	"progres/views/pasienview"
	"progres/views/susterview"
	"runtime"
)

func Clear() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	user.DataStatisUser()
	dokter.DataStatisDokter()
	var username, password string
	var pilih int
	for pilih != 2 {
		Clear()
		fmt.Println("1. Login")
		fmt.Println("2. Exit")
		fmt.Print("pilih: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			Clear()
			fmt.Print("username: ")
			fmt.Scan(&username)
			fmt.Print("password: ")
			fmt.Scan(&password)
			status := usercontroller.Login(username, password)
			fmt.Printf("status: %s/n", status)
			if status == "administrator" {
				var menu int
				for menu != 4 {
					Clear()
					views.OlahData()
					fmt.Scan(&menu)
					switch menu {
					case 1:
						var pilih int
						var back string
						for pilih != 6 {
							Clear()
							views.SubMenuOlahData()
							fmt.Scan(&pilih)
							switch pilih {
							case 1:
								Clear()
								dokterview.InsertDokter()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 2:
								Clear()
								dokterview.UpdateDoter()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 3:
								Clear()
								dokterview.DeleteDoter()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 4:
								Clear()
								dokterview.ViewById()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 5:
								Clear()
								dokterview.ViewDokter()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 6:
							default:
								Clear()
								println("pilihan tidak valid")
							}
						}
					case 2:
						var pilih int
						var back string
						for pilih != 6 {
							Clear()
							views.SubMenuOlahData()
							fmt.Scan(&pilih)
							switch pilih {
							case 1:
								Clear()
								susterview.InsertSuster()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 2:
								Clear()
								susterview.UpdateSuster()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 3:
								Clear()
								susterview.DeleteDoter()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 4:
								Clear()
								susterview.ViewById()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 5:
								Clear()
								susterview.ViewSuster()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 6:
							default:
								Clear()
								println("pilihan tidak valid")
							}
						}
					case 3:
						var pilih int
						var back string
						for pilih != 6 {
							Clear()
							views.SubMenuOlahData()
							fmt.Scan(&pilih)
							switch pilih {
							case 1:
								Clear()
								pasienview.InsertPasien()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 2:
								Clear()
								pasienview.UpdatePasien()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 3:
								Clear()
								pasienview.DeletePasien()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 4:
								Clear()
								pasienview.ViewById()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 5:
								Clear()
								pasienview.ViewPasien()
								fmt.Print("kembali (y/t): ")
								fmt.Scan(&back)
								if back == "t" || back == "T" {
									os.Exit(0)
								}
							case 6:
							default:
								Clear()
								println("pilihan tidak valid")
							}
						}
					}
				}

			} else if status == "petugas loket" {
				var pilih int
				Clear()
				for pilih != 4 {
					views.MenuAntrian()
					fmt.Print("pilihan : ")
					fmt.Scan(&pilih)
					switch pilih {
					case 1:
						antrianview.TambahAntrian()
					case 2:
						antrianview.NextAntrian()
					case 3:
						antrianview.ViewAllAntrian()
					case 4:
						fmt.Println("Terima kasih telah menggunakan program antrianview.")
					default:
						fmt.Println("Pilihan tidak valid.")
					}
				}

			} else {
				fmt.Println("username atau password salah")
			}

		} else if pilih == 2 {
			break
		} else {
			fmt.Println("pilihan tidak valid, silahkan pilih menu yang tersedia")
		}
	}
}
