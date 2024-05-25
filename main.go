package main

import "fmt"

func main() {
	pesanSapaan()
	initDefaultData()
	for {
		fmt.Println("╔═════════════════════════════════╗")
		fmt.Println("║          Menu Utama             ║")
		fmt.Println("╚═════════════════════════════════╝")
		fmt.Println("1. Lihat Kumpulan Data Desa")
		fmt.Println("2. Periksa Data Penduduk Desa")
		fmt.Println("3. Hentikan program")
		fmt.Print("Silakan pilih opsi Anda: ")
		var pilih string
		fmt.Scanln(&pilih)

		if pilih == "1" {
			kelolaDataDesa()
		} else if pilih == "2" {
			kelolaDataPenduduk()
		} else if pilih == "3" {
			fmt.Println("Program dihentikan.")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func pesanSapaan() {
	fmt.Println("\n╔═════════════════════════════════╗")
	fmt.Println("║                                 ║")
	fmt.Println("║        Selamat Datang!          ║")
	fmt.Println("║                                 ║")
	fmt.Println("╚═════════════════════════════════╝")
}
