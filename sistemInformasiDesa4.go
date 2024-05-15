package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Desa struct {
	Provinsi  string
	Kabupaten string
	Kecamatan string
	Nama      string
	JumlahRT  int
	JumlahRW  int
}

type Penduduk struct {
	NIK          string
	Nama         string
	TanggalLahir string
	Status       string
}

var desaList []Desa
var pendudukList map[string][]Penduduk

func main() {
	pendudukList = make(map[string][]Penduduk)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Menu Utama:")
		fmt.Println("1. Kumpulan data desa")
		fmt.Println("2. Periksa data penduduk desa")
		fmt.Println("3. Hentikan program")
		fmt.Print("Pilih opsi: ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			kelolaDataDesa(scanner)
		case "2":
			kelolaDataPenduduk(scanner)
		case "3":
			fmt.Println("Program dihentikan.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func kelolaDataDesa(scanner *bufio.Scanner) {
	for {
		fmt.Println("Kumpulan Data Desa:")
		fmt.Println("1. Masukkan data desa baru")
		fmt.Println("2. Periksa daftar desa")
		fmt.Println("3. Ubah data desa")
		fmt.Println("4. Kembali ke menu awal")
		fmt.Print("Pilih opsi: ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			masukkanDataDesa(scanner)
		case "2":
			periksaDaftarDesa(scanner)
		case "3":
			ubahDataDesa(scanner)
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func masukkanDataDesa(scanner *bufio.Scanner) {
	fmt.Println("Masukkan data desa baru:")
	fmt.Print("Provinsi: ")
	scanner.Scan()
	provinsi := scanner.Text()
	fmt.Print("Kabupaten: ")
	scanner.Scan()
	kabupaten := scanner.Text()
	fmt.Print("Kecamatan: ")
	scanner.Scan()
	kecamatan := scanner.Text()
	fmt.Print("Nama Desa: ")
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Jumlah RT: ")
	scanner.Scan()
	jumlahRT, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Jumlah RW: ")
	scanner.Scan()
	jumlahRW, _ := strconv.Atoi(scanner.Text())

	desa := Desa{
		Provinsi:  provinsi,
		Kabupaten: kabupaten,
		Kecamatan: kecamatan,
		Nama:      nama,
		JumlahRT:  jumlahRT,
		JumlahRW:  jumlahRW,
	}

	desaList = append(desaList, desa)
	fmt.Println("Data desa berhasil ditambahkan.")
}

func periksaDaftarDesa(scanner *bufio.Scanner) {
	if len(desaList) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	for i, desa := range desaList {
		fmt.Printf("%d. %s, %s, %s, %s, RT: %d, RW: %d\n", i+1, desa.Provinsi, desa.Kabupaten, desa.Kecamatan, desa.Nama, desa.JumlahRT, desa.JumlahRW)
	}
	fmt.Println("Kembali ke menu awal")
}

func ubahDataDesa(scanner *bufio.Scanner) {
	if len(desaList) == 0 {
		fmt.Println("Data masih kosong.")
		return
	}

	fmt.Println("Daftar Desa:")
	for i, desa := range desaList {
		fmt.Printf("%d. %s, %s, %s, %s\n", i+1, desa.Provinsi, desa.Kabupaten, desa.Kecamatan, desa.Nama)
	}
	fmt.Print("Pilih nomor desa yang ingin diubah: ")
	scanner.Scan()
	nomor, _ := strconv.Atoi(scanner.Text())
	if nomor < 1 || nomor > len(desaList) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	desa := &desaList[nomor-1]
	fmt.Printf("Ubah data desa %s:\n", desa.Nama)
	fmt.Printf("Provinsi (%s): ", desa.Provinsi)
	scanner.Scan()
	provinsi := scanner.Text()
	if provinsi != "" {
		desa.Provinsi = provinsi
	}
	fmt.Printf("Kabupaten (%s): ", desa.Kabupaten)
	scanner.Scan()
	kabupaten := scanner.Text()
	if kabupaten != "" {
		desa.Kabupaten = kabupaten
	}
	fmt.Printf("Kecamatan (%s): ", desa.Kecamatan)
	scanner.Scan()
	kecamatan := scanner.Text()
	if kecamatan != "" {
		desa.Kecamatan = kecamatan
	}
	fmt.Printf("Nama Desa (%s): ", desa.Nama)
	scanner.Scan()
	nama := scanner.Text()
	if nama != "" {
		desa.Nama = nama
	}
	fmt.Printf("Jumlah RT (%d): ", desa.JumlahRT)
	scanner.Scan()
	jumlahRT, _ := strconv.Atoi(scanner.Text())
	if jumlahRT != 0 {
		desa.JumlahRT = jumlahRT
	}
	fmt.Printf("Jumlah RW (%d): ", desa.JumlahRW)
	scanner.Scan()
	jumlahRW, _ := strconv.Atoi(scanner.Text())
	if jumlahRW != 0 {
		desa.JumlahRW = jumlahRW
	}

	fmt.Println("Data desa berhasil diubah.")
}

func kelolaDataPenduduk(scanner *bufio.Scanner) {
	for {
		fmt.Println("Periksa Data Penduduk Desa:")
		fmt.Println("1. Masukkan data penduduk")
		fmt.Println("2. Periksa daftar penduduk")
		fmt.Println("3. Ubah data penduduk")
		fmt.Println("4. Kembali ke menu awal")
		fmt.Print("Pilih opsi: ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			masukkanDataPenduduk(scanner)
		case "2":
			periksaDaftarPenduduk(scanner)
		case "3":
			ubahDataPenduduk(scanner)
		case "4":
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func masukkanDataPenduduk(scanner *bufio.Scanner) {
	if len(desaList) == 0 {
		fmt.Println("Data desa masih kosong.")
		return
	}

	fmt.Println("Pilih desa:")
	for i, desa := range desaList {
		fmt.Printf("%d. %s\n", i+1, desa.Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	scanner.Scan()
	nomor, _ := strconv.Atoi(scanner.Text())
	if nomor < 1 || nomor > len(desaList) {
		fmt.Println("Nomor desa tidak valid.")
		return
	}
	namaDesa := desaList[nomor-1].Nama

	fmt.Println("Masukkan data penduduk:")
	fmt.Print("NIK: ")
	scanner.Scan()
	nik := scanner.Text()
	fmt.Print("Nama: ")
	scanner.Scan()
	nama := scanner.Text()
	fmt.Print("Tanggal Lahir (dd-mm-yyyy): ")
	scanner.Scan()
	tanggalLahir := scanner.Text()
	fmt.Print("Status Perkawinan (1. Kawin, 2. Belum Kawin): ")
	scanner.Scan()
	status := scanner.Text()

	penduduk := Penduduk{
		NIK:          nik,
		Nama:         nama,
		TanggalLahir: tanggalLahir,
		Status:       status,
	}

	pendudukList[namaDesa] = append(pendudukList[namaDesa], penduduk)
	fmt.Println("Data penduduk berhasil ditambahkan.")
}

func periksaDaftarPenduduk(scanner *bufio.Scanner) {
	if len(desaList) == 0 {
		fmt.Println("Data desa masih kosong.")
		return
	}

	fmt.Println("Pilih desa:")
	for i, desa := range desaList {
		fmt.Printf("%d. %s\n", i+1, desa.Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	scanner.Scan()
	nomor, _ := strconv.Atoi(scanner.Text())
	if nomor < 1 || nomor > len(desaList) {
		fmt.Println("Nomor desa tidak valid.")
		return
	}
	namaDesa := desaList[nomor-1].Nama

	pendudukDesa, exists := pendudukList[namaDesa]
	if !exists || len(pendudukDesa) == 0 {
		fmt.Println("Data penduduk masih kosong.")
		return
	}

	for i, penduduk := range pendudukDesa {
		fmt.Printf("%d. NIK: %s, Nama: %s, Tanggal Lahir: %s, Status: %s\n", i+1, penduduk.NIK, penduduk.Nama, penduduk.TanggalLahir, penduduk.Status)
	}
	fmt.Println("Kembali ke menu awal")
}

func ubahDataPenduduk(scanner *bufio.Scanner) {
	if len(desaList) == 0 {
		fmt.Println("Data desa masih kosong.")
		return
	}

	fmt.Println("Pilih desa:")
	for i, desa := range desaList {
		fmt.Printf("%d. %s\n", i+1, desa.Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	scanner.Scan()
	nomor, _ := strconv.Atoi(scanner.Text())
	if nomor < 1 || nomor > len(desaList) {
		fmt.Println("Nomor desa tidak valid.")
		return
	}
	namaDesa := desaList[nomor-1].Nama

	pendudukDesa, exists := pendudukList[namaDesa]
	if !exists || len(pendudukDesa) == 0 {
		fmt.Println("Data penduduk masih kosong.")
		return
	}

	fmt.Println("Pilih penduduk:")
	for i, penduduk := range pendudukDesa {
		fmt.Printf("%d. NIK: %s, Nama: %s\n", i+1, penduduk.NIK, penduduk.Nama)
	}
	fmt.Print("Masukkan nomor penduduk: ")
	scanner.Scan()
	nomorPenduduk, _ := strconv.Atoi(scanner.Text())
	if nomorPenduduk < 1 || nomorPenduduk > len(pendudukDesa) {
		fmt.Println("Nomor penduduk tidak valid.")
		return
	}

	penduduk := &pendudukDesa[nomorPenduduk-1]
	fmt.Printf("Ubah data penduduk NIK %s:\n", penduduk.NIK)
	fmt.Printf("Nama (%s): ", penduduk.Nama)
	scanner.Scan()
	nama := scanner.Text()
	if nama != "" {
		penduduk.Nama = nama
	}
	fmt.Printf("Tanggal Lahir (%s): ", penduduk.TanggalLahir)
	scanner.Scan()
	tanggalLahir := scanner.Text()
	if tanggalLahir != "" {
		penduduk.TanggalLahir = tanggalLahir
	}
	fmt.Printf("Status Perkawinan (%s): ", penduduk.Status)
	scanner.Scan()
	status := scanner.Text()
	if status != "" {
		penduduk.Status = status
	}

	fmt.Println("Data penduduk berhasil diubah.")
}
