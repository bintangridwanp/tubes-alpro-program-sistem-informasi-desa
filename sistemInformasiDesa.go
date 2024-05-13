package main

import (
	"fmt"
	"os"
)

type Desa struct {
	NamaLokasi string
	NamaDesa   string
	NamaRT     string
	NamaRW     string
}

type Penduduk struct {
	NIK              string
	Nama             string
	TanggalLahir     string
	StatusPerkawinan string
}

var dataDesa Desa
var dataPenduduk []Penduduk

func main() {
	var pilihan int

	for {
		fmt.Println("\nPilihan:")
		fmt.Println("1. Masukan data desa")
		fmt.Println("2. Masukan data penduduk")
		fmt.Println("3. Edit data desa")
		fmt.Println("4. Hentikan program")
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			if dataDesa.NamaDesa != "" {
				fmt.Println("Data desa sudah dimasukkan sebelumnya.")
				fmt.Println("Apakah Anda ingin mengedit data desa? (y/n)")
				var edit string
				fmt.Scanln(&edit)
				if edit == "y" {
					editDataDesa()
				}
			} else {
				masukkanDataDesa()
			}
		case 2:
			if dataDesa.NamaDesa == "" {
				fmt.Println("Masukkan data desa terlebih dahulu.")
			} else {
				masukkanDataPenduduk()
			}
		case 3:
			if dataDesa.NamaDesa == "" {
				fmt.Println("Masukkan data desa terlebih dahulu.")
			} else {
				editDataPenduduk()
			}
		case 4:
			fmt.Println("Program dihentikan.")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan masukkan pilihan yang benar.")
		}
	}
}

func masukkanDataDesa() {
	fmt.Println("\nMasukkan data desa:")
	fmt.Print("Nama lokasi desa: ")
	fmt.Scanln(&dataDesa.NamaLokasi)
	fmt.Print("Nama desa: ")
	fmt.Scanln(&dataDesa.NamaDesa)
	fmt.Print("Nama RT: ")
	fmt.Scanln(&dataDesa.NamaRT)
	fmt.Print("Nama RW: ")
	fmt.Scanln(&dataDesa.NamaRW)
	fmt.Println("Data desa berhasil dimasukkan.")
}

func editDataDesa() {
	fmt.Println("\nEdit data desa:")
	fmt.Print("Nama lokasi desa: ")
	fmt.Scanln(&dataDesa.NamaLokasi)
	fmt.Print("Nama desa: ")
	fmt.Scanln(&dataDesa.NamaDesa)
	fmt.Print("Nama RT: ")
	fmt.Scanln(&dataDesa.NamaRT)
	fmt.Print("Nama RW: ")
	fmt.Scanln(&dataDesa.NamaRW)
	fmt.Println("Data desa berhasil diubah.")
}

func masukkanDataPenduduk() {
	var penduduk Penduduk
	fmt.Println("\nMasukkan data penduduk:")
	fmt.Print("NIK: ")
	fmt.Scanln(&penduduk.NIK)
	fmt.Print("Nama: ")
	fmt.Scanln(&penduduk.Nama)
	fmt.Print("Tanggal lahir (dd MMMM yyyy): ")
	fmt.Scanln(&penduduk.TanggalLahir)
	fmt.Print("Status perkawinan (1. Kawin, 2. Belum kawin): ")
	fmt.Scanln(&penduduk.StatusPerkawinan)
	dataPenduduk = append(dataPenduduk, penduduk)
	fmt.Println("Data penduduk berhasil dimasukkan.")
}

func editDataPenduduk() {
	if len(dataPenduduk) == 0 {
		fmt.Println("Belum ada data penduduk yang dimasukkan.")
		return
	}

	fmt.Println("\nData penduduk:")
	for i, penduduk := range dataPenduduk {
		fmt.Printf("%d. NIK: %s, Nama: %s, Tanggal Lahir: %s, Status Perkawinan: %s\n", i+1, penduduk.NIK, penduduk.Nama, penduduk.TanggalLahir, penduduk.StatusPerkawinan)
	}

	fmt.Print("Masukkan nomor data penduduk yang ingin diubah: ")
	var index int
	fmt.Scanln(&index)
	index--

	if index < 0 || index >= len(dataPenduduk) {
		fmt.Println("Nomor data penduduk tidak valid.")
		return
	}

	var penduduk Penduduk
	fmt.Println("\nMasukkan data penduduk baru:")
	fmt.Print("NIK: ")
	fmt.Scanln(&penduduk.NIK)
	fmt.Print("Nama: ")
	fmt.Scanln(&penduduk.Nama)
	fmt.Print("Tanggal lahir (dd MMMM yyyy): ")
	fmt.Scanln(&penduduk.TanggalLahir)
	fmt.Print("Status perkawinan (1. Kawin, 2. Belum kawin): ")
	fmt.Scanln(&penduduk.StatusPerkawinan)

	dataPenduduk[index] = penduduk
	fmt.Println("Data penduduk berhasil diubah.")
}
