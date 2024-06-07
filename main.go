package main

import (
	"fmt"
)

func main() {
	pesanSapaan()
	initDefaultData()
	var pilih string
	for pilih != "4" {
		fmt.Println("╔═════════════════════════════════╗")
		fmt.Println("║          Menu Utama             ║")
		fmt.Println("╚═════════════════════════════════╝")
		fmt.Println("1. Lihat Kumpulan Data Desa")
		fmt.Println("2. Periksa Data Penduduk Desa")
		fmt.Println("3. Laporan UMKM")
		fmt.Println("4. Hentikan program")
		fmt.Print("Silakan pilih opsi Anda: ")
		fmt.Scanln(&pilih)

		if pilih == "1" {
			kelolaDataDesa()
		} else if pilih == "2" {
			kelolaDataPenduduk()
		} else if pilih == "3" {
			kelolaLaporanUMKM()
		} else if pilih == "4" {
			fmt.Println("Program dihentikan.")
			continue
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
