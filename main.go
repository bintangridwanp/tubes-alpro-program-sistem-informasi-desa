package main

import (
	"fmt"
	"os"
)

func main() {
	pesanSapaan()
	initDefaultData()
	for {
		fmt.Println("╔═════════════════════════════════╗")
		fmt.Println("║          Menu Utama             ║")
		fmt.Println("╚═════════════════════════════════╝")
		fmt.Println("1. Lihat Kumpulan Data Desa")
		fmt.Println("2. Periksa Data Penduduk Desa")
		fmt.Println("3. Laporan UMKM")
		fmt.Println("4. Hentikan program")
		fmt.Print("Silakan pilih opsi Anda: ")
		var pilih string
		fmt.Scanln(&pilih)

		if pilih == "1" {
			if kelolaDataDesa() == "main" {
				continue
			}
		} else if pilih == "2" {
			if kelolaDataPenduduk() == "main" {
				continue
			}
		} else if pilih == "3" {
			if kelolaLaporanUMKM() == "main" {
				continue
			}
		} else if pilih == "4" {
			fmt.Println("Program dihentikan.")
			os.Exit(0)
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
