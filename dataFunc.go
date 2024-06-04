package main

import (
	"fmt"
)

// Function utama untuk pilihan "1. Lihat Kumpulan Data Desa" di Funtion main
func kelolaDataDesa() string {
	for {
		fmt.Println("╔═════════════════════════════════╗")
		fmt.Println("║          Menu Desa              ║")
		fmt.Println("╚═════════════════════════════════╝")
		fmt.Println("Kumpulan Data Desa: ")
		fmt.Println("1. Input Data Desa Baru")
		fmt.Println("2. Periksa Daftar Desa")
		fmt.Println("3. Ubah Data Desa")
		fmt.Println("4. Kembali ke Menu Utama")
		fmt.Print("Silakan pilih opsi yang diinginkan: ")
		var pilih string
		fmt.Scanln(&pilih)
		if pilih == "1" {
			masukkanDataDesa(&desaList, &desaCount)
		} else if pilih == "2" {
			periksaDaftarDesa(&desaList, desaCount)
		} else if pilih == "3" {
			ubahDataDesa(&desaList, desaCount)
		} else if pilih == "4" {
			return "main"
		} else {
			fmt.Println("Pilihan tidak valid.")
		}

	}
}

func masukkanDataDesa(desaList *[MAX]Desa, desaCount *int) string {
	if *desaCount >= MAX {
		return "Tidak bisa menambahkan data desa baru, array penuh."
	}

	var provinsi, kabupaten, kecamatan, nama string
	var jumlahRT, jumlahRW int

	fmt.Println("Masukkan data desa baru:")
	fmt.Print("Provinsi: ")
	fmt.Scanln(&provinsi)
	fmt.Print("Kabupaten: ")
	fmt.Scanln(&kabupaten)
	fmt.Print("Kecamatan: ")
	fmt.Scanln(&kecamatan)
	fmt.Print("Nama Desa: ")
	fmt.Scanln(&nama)
	fmt.Print("Jumlah RT: ")
	fmt.Scanln(&jumlahRT)
	fmt.Print("Jumlah RW: ")
	fmt.Scanln(&jumlahRW)

	desa := Desa{
		Provinsi:  provinsi,
		Kabupaten: kabupaten,
		Kecamatan: kecamatan,
		Nama:      nama,
		JumlahRT:  jumlahRT,
		JumlahRW:  jumlahRW,
	}

	desaList[*desaCount] = desa
	*desaCount++
	return "Data desa berhasil ditambahkan."
}

func periksaDaftarDesa(desaList *[MAX]Desa, desaCount int) string {
	if desaCount == 0 {
		return "Data masih kosong."
	}

	for i := 0; i < desaCount; i++ {
		desa := desaList[i]
		fmt.Printf("%d. %s, %s, %s, %s, RT: %d, RW: %d\n", i+1, desa.Provinsi, desa.Kabupaten, desa.Kecamatan, desa.Nama, desa.JumlahRT, desa.JumlahRW)
	}
	return "Kembali ke menu awal"
}

func ubahDataDesa(desaList *[MAX]Desa, desaCount int) string {
	if desaCount == 0 {
		return "Data masih kosong."
	}

	fmt.Println("Daftar Desa:")
	for i := 0; i < desaCount; i++ {
		desa := desaList[i]
		fmt.Printf("%d. %s, %s, %s, %s\n", i+1, desa.Provinsi, desa.Kabupaten, desa.Kecamatan, desa.Nama)
	}
	fmt.Print("Pilih nomor desa yang ingin diubah: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > desaCount {
		return "Nomor tidak valid."
	}

	desa := &desaList[nomor-1]
	var input string
	fmt.Printf("Ubah data desa %s:\n", desa.Nama)
	fmt.Printf("Provinsi (%s): ", desa.Provinsi)
	fmt.Scanln(&input)
	if input != "" {
		desa.Provinsi = input
	}
	fmt.Printf("Kabupaten (%s): ", desa.Kabupaten)
	fmt.Scanln(&input)
	if input != "" {
		desa.Kabupaten = input
	}
	fmt.Printf("Kecamatan (%s): ", desa.Kecamatan)
	fmt.Scanln(&input)
	if input != "" {
		desa.Kecamatan = input
	}
	fmt.Printf("Nama Desa (%s): ", desa.Nama)
	fmt.Scanln(&input)
	if input != "" {
		desa.Nama = input
	}
	fmt.Printf("Jumlah RT (%d): ", desa.JumlahRT)
	fmt.Scanln(&input)
	if input != "" {
		fmt.Sscanf(input, "%d", &desa.JumlahRT)
	}
	fmt.Printf("Jumlah RW (%d): ", desa.JumlahRW)
	fmt.Scanln(&input)
	if input != "" {
		fmt.Scanf(input, "%d", &desa.JumlahRW)
	}

	return "Data desa berhasil diubah."
}

// Function utama untuk pilihan "2. Periksa Data Penduduk Desa" di Funtion main
func kelolaDataPenduduk() string {
	for {
		fmt.Println("╔═════════════════════════════════╗")
		fmt.Println("║          Menu Penduduk          ║")
		fmt.Println("╚═════════════════════════════════╝")
		fmt.Println("Periksa Data Penduduk Desa:")
		fmt.Println("1. Masukkan data penduduk")
		fmt.Println("2. Periksa daftar penduduk")
		fmt.Println("3. Ubah data penduduk")
		fmt.Println("4. Ajukan pergantian status perkawinan")
		fmt.Println("5. Kembali ke menu awal")
		fmt.Print("Pilih opsi: ")
		var pilih string
		fmt.Scanln(&pilih)

		if pilih == "1" {
			masukkanDataPenduduk(&pendudukList, &pendudukCount)
		} else if pilih == "2" {
			if periksaDaftarPenduduk(&pendudukList, pendudukCount) == "periksa" {

			}
		} else if pilih == "3" {
			if ubahDataPenduduk(&pendudukList, &pendudukCount) == "ubah" {

			}
		} else if pilih == "4" {
			if ajukanPergantianStatusPerkawinan(&desaList, desaCount, &pendudukList, &pendudukCount) == "gantiStatus" {

			}
		} else if pilih == "5" {
			return "main"
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func masukkanDataPenduduk(pendudukList *[MAX][MAX]Penduduk, pendudukCount *[MAX]int) string {
	if desaCount == 0 {
		return "Data desa masih kosong."
	}

	fmt.Println("Pilih desa:")
	for i := 0; i < desaCount; i++ {
		fmt.Printf("%d. %s\n", i+1, desaList[i].Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > desaCount {
		return "Nomor desa tidak valid."
	}

	var nik, nama, tanggalLahir, status string

	fmt.Println("Masukkan data penduduk:")
	fmt.Print("NIK: ")
	fmt.Scanln(&nik)
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Tanggal Lahir (dd-mm-yyyy): ")
	fmt.Scanln(&tanggalLahir)
	fmt.Print("Status Perkawinan (1. Kawin, 2. Belum Kawin): ")
	fmt.Scanln(&status)

	penduduk := Penduduk{
		NIK:          nik,
		Nama:         nama,
		TanggalLahir: tanggalLahir,
		Status:       status,
	}

	desaIdx := nomor - 1
	if pendudukCount[desaIdx] < MAX {
		pendudukList[desaIdx][pendudukCount[desaIdx]] = penduduk
		pendudukCount[desaIdx]++
		return "Data penduduk berhasil ditambahkan."
	} else {
		return "Tidak bisa menambahkan data penduduk baru, array penuh."
	}
}

func periksaDaftarPenduduk(pendudukList *[MAX][MAX]Penduduk, pendudukCount [MAX]int) string {
	if desaCount == 0 {
		fmt.Println("Data desa masih kosong.")
		return "periksa"
	}

	fmt.Println("Pilih desa:")
	for i := 0; i < desaCount; i++ {
		fmt.Printf("%d. %s\n", i+1, desaList[i].Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > desaCount {
		fmt.Println("Nomor desa tidak valid.")
		return "periksa"
	}

	desaIdx := nomor - 1
	if pendudukCount[desaIdx] == 0 {
		fmt.Println("Data penduduk masih kosong.")
		return "periksa"
	}

	for i := 0; i < pendudukCount[desaIdx]; i++ {
		penduduk := pendudukList[desaIdx][i]
		fmt.Printf("%d. NIK: %s, Nama: %s, Tanggal Lahir: %s, Status: %s\n", i+1, penduduk.NIK, penduduk.Nama, penduduk.TanggalLahir, penduduk.Status)
	}

	for {
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Cari data penduduk")
		fmt.Println("2. Kembali ke menu awal")
		fmt.Print("Pilih opsi: ")
		var pilih string
		fmt.Scanln(&pilih)
		if pilih == "1" {
			if cariDataPenduduk(pendudukList[desaIdx], pendudukCount[desaIdx]) == "cari" {

			}
		} else if pilih == "2" {
			return "periksa"
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func cariDataPenduduk(pendudukDesa [MAX]Penduduk, count int) string {
	fmt.Println("Cari data penduduk berdasarkan:")
	fmt.Println("1. NIK")
	fmt.Println("2. Nama")
	fmt.Print("Pilih opsi: ")
	var pilih string
	fmt.Scanln(&pilih)

	if pilih != "1" && pilih != "2" {
		fmt.Println("Pilihan tidak valid.")
		return "Pilihan tidak valid."
	}

	// menggunakan algoritma Pencarian sekuensial
	var cari string
	fmt.Print("Masukkan pencarian: ")
	fmt.Scanln(&cari)

	found := false
	if pilih == "1" {
		for i := 0; i < count && !found; i++ {
			if pendudukDesa[i].NIK == cari {
				fmt.Printf("NIK: %s, Nama: %s, Tanggal Lahir: %s, Status: %s\n", pendudukDesa[i].NIK, pendudukDesa[i].Nama, pendudukDesa[i].TanggalLahir, pendudukDesa[i].Status)
				found = true
			}
		}
	} else if pilih == "2" {
		for i := 0; i < count && !found; i++ {
			if pendudukDesa[i].Nama == cari {
				fmt.Printf("NIK: %s, Nama: %s, Tanggal Lahir: %s, Status: %s\n", pendudukDesa[i].NIK, pendudukDesa[i].Nama, pendudukDesa[i].TanggalLahir, pendudukDesa[i].Status)
				found = true
			}
		}
	}

	if !found {
		fmt.Println("Data penduduk tidak ditemukan.")
	}
	return "cari"
}

func ubahDataPenduduk(pendudukList *[MAX][MAX]Penduduk, pendudukCount *[MAX]int) string {
	if desaCount == 0 {
		fmt.Println("Data desa masih kosong.")
		return "ubah"
	}

	fmt.Println("Pilih desa:")
	for i := 0; i < desaCount; i++ {
		fmt.Printf("%d. %s\n", i+1, desaList[i].Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > desaCount {
		fmt.Println("Nomor desa tidak valid.")
		return "ubah"
	}

	desaIdx := nomor - 1
	if pendudukCount[desaIdx] == 0 {
		fmt.Println("Data penduduk masih kosong.")
		return "ubah"
	}

	fmt.Println("Pilih penduduk:")
	for i := 0; i < pendudukCount[desaIdx]; i++ {
		penduduk := pendudukList[desaIdx][i]
		fmt.Printf("%d. NIK: %s, Nama: %s\n", i+1, penduduk.NIK, penduduk.Nama)
	}
	fmt.Print("Masukkan nomor penduduk: ")
	var nomorPenduduk int
	fmt.Scanln(&nomorPenduduk)
	if nomorPenduduk < 1 || nomorPenduduk > pendudukCount[desaIdx] {
		fmt.Println("Nomor penduduk tidak valid.")
		return "ubah"
	}

	penduduk := &pendudukList[desaIdx][nomorPenduduk-1]
	fmt.Println("Pilih opsi:")
	fmt.Println("1. Edit data penduduk")
	fmt.Println("2. Hapus data penduduk")
	fmt.Println("3. Kembali ke menu awal")
	fmt.Print("Pilih opsi: ")
	var pilih string
	fmt.Scanln(&pilih)

	if pilih == "1" {
		editDataPenduduk(penduduk)
	} else if pilih == "2" {
		for i := nomorPenduduk - 1; i < pendudukCount[desaIdx]-1; i++ {
			pendudukList[desaIdx][i] = pendudukList[desaIdx][i+1]
		}
		pendudukCount[desaIdx]--
		fmt.Println("Data penduduk berhasil dihapus.")
	} else if pilih == "3" {
		return "cari"
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
	return "ubah"
}

func editDataPenduduk(penduduk *Penduduk) {
	var input string
	fmt.Printf("Ubah data penduduk NIK %s:\n", penduduk.NIK)
	fmt.Printf("Nama (%s): ", penduduk.Nama)
	fmt.Scanln(&input)
	if input != "" {
		penduduk.Nama = input
	}
	fmt.Printf("Tanggal Lahir (%s): ", penduduk.TanggalLahir)
	fmt.Scanln(&input)
	if input != "" {
		penduduk.TanggalLahir = input
	}
	fmt.Printf("Status Perkawinan (%s): ", penduduk.Status)
	fmt.Scanln(&input)
	if input != "" {
		penduduk.Status = input
	}

	fmt.Println("Data penduduk berhasil diubah.")
}

func ajukanPergantianStatusPerkawinan(desaList *[MAX]Desa, desaCount int, pendudukList *[MAX][MAX]Penduduk, pendudukCount *[MAX]int) string {
	if desaCount == 0 {
		fmt.Println("Data desa masih kosong.")
		return "gantiStatus"
	}

	fmt.Println("Pilih desa:")
	for i := 0; i < desaCount; i++ {
		fmt.Printf("%d. %s\n", i+1, desaList[i].Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > desaCount {
		fmt.Println("Nomor desa tidak valid.")
		return "gantiStatus"
	}
	desaIdx := nomor - 1

	if pendudukCount[desaIdx] == 0 {
		fmt.Println("Data penduduk masih kosong.")
		return "gantiStatus"
	}

	fmt.Println("Pilih penduduk:")
	for i := 0; i < pendudukCount[desaIdx]; i++ {
		penduduk := pendudukList[desaIdx][i]
		fmt.Printf("%d. NIK: %s, Nama: %s, Status: %s\n", i+1, penduduk.NIK, penduduk.Nama, penduduk.Status)
	}
	fmt.Print("Masukkan nomor penduduk: ")
	var nomorPenduduk int
	fmt.Scanln(&nomorPenduduk)
	if nomorPenduduk < 1 || nomorPenduduk > pendudukCount[desaIdx] {
		fmt.Println("Nomor penduduk tidak valid.")
		return "gantiStatus"
	}

	penduduk := &pendudukList[desaIdx][nomorPenduduk-1]
	fmt.Println("Ajukan pergantian status perkawinan:")
	fmt.Print("Status Perkawinan Baru (1. Kawin, 2. Belum Kawin): ")
	var statusBaru string
	fmt.Scanln(&statusBaru)

	if statusBaru == "1" {
		penduduk.Status = "Kawin"
	} else if statusBaru == "2" {
		penduduk.Status = "Belum Kawin"
	} else {
		fmt.Println("Status tidak valid.")
		return "gantiStatus"
	}

	fmt.Println("Pengajuan pergantian status perkawinan berhasil diajukan.")
	return "gantiStatus"
}

// Function utama untuk pilihan "3. Laporan UMKM" di Funtion main
func kelolaLaporanUMKM() string {
	initDefaultDataUMKM()
	for {
		fmt.Println("╔═════════════════════════════════╗")
		fmt.Println("║        Menu Laporan UMKM        ║")
		fmt.Println("╚═════════════════════════════════╝")
		fmt.Println("1. Laporkan UMKM")
		fmt.Println("2. Lihat Riwayat Laporan")
		fmt.Println("3. Kembali ke menu awal")
		fmt.Print("Pilih opsi: ")
		var pilih string
		fmt.Scanln(&pilih)

		if pilih == "1" {
			if laporkanUMKM(&umkmList, &umkmCount) == "laporan" {

			}
		} else if pilih == "2" {
			if lihatRiwayatLaporan(&umkmList, &umkmCount) == "riwayat" {

			}
		} else if pilih == "3" {
			return "main"
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func laporkanUMKM(umkmList *[MAX][MAX]UMKM, umkmCount *[MAX]int) string {
	if desaCount == 0 {
		fmt.Println("Data desa masih kosong.")
		return "laporan"
	}

	fmt.Println("Pilih desa:")
	for i := 0; i < desaCount; i++ {
		fmt.Printf("%d. %s\n", i+1, desaList[i].Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > desaCount {
		fmt.Println("Nomor desa tidak valid.")
		return "laporan"
	}

	var namaUMKM string
	var totalUang int

	fmt.Println("Masukkan laporan UMKM:")
	fmt.Print("Nama UMKM: ")
	fmt.Scanln(&namaUMKM)
	fmt.Print("Total Uang: ")
	fmt.Scanln(&totalUang)

	umkm := UMKM{
		Nama:      namaUMKM,
		TotalUang: totalUang,
	}

	desaIdx := nomor - 1
	if umkmCount[desaIdx] < MAX {
		umkmList[desaIdx][umkmCount[desaIdx]] = umkm
		umkmCount[desaIdx]++
		fmt.Println("Laporan UMKM berhasil ditambahkan.")
	} else {
		fmt.Println("Tidak bisa menambahkan laporan UMKM baru, array penuh.")
	}

	return "laporan"
}

func lihatRiwayatLaporan(umkmList *[MAX][MAX]UMKM, umkmCount *[MAX]int) string {
	if desaCount == 0 {
		fmt.Println("Data desa masih kosong.")
		return "riwayat"
	}

	fmt.Println("Pilih desa:")
	for i := 0; i < desaCount; i++ {
		fmt.Printf("%d. %s\n", i+1, desaList[i].Nama)
	}
	fmt.Print("Masukkan nomor desa: ")
	var nomor int
	fmt.Scanln(&nomor)
	if nomor < 1 || nomor > desaCount {
		fmt.Println("Nomor desa tidak valid.")
		return "riwayat"
	}

	desaIdx := nomor - 1
	if umkmCount[desaIdx] == 0 {
		fmt.Println("Data masih kosong.")
		return "riwayat"
	}

	insertionSortUMKM(umkmList[desaIdx], umkmCount[desaIdx])

	for i := 0; i < umkmCount[desaIdx]; i++ {
		umkm := umkmList[desaIdx][i]
		fmt.Printf("%d. Nama UMKM: %s, Total Uang: %d\n", i+1, umkm.Nama, umkm.TotalUang)
	}

	var totalUang int
	for i := 0; i < desaCount; i++ {
		for j := 0; j < umkmCount[i]; j++ {
			totalUang += umkmList[i][j].TotalUang
		}
	}

	fmt.Println("Total Uang Laporan: ", totalUang)
	return "riwayat"
}

func insertionSortUMKM(umkmList [MAX]UMKM, umkmCount int) {
	for i := 1; i < umkmCount; i++ {
		key := umkmList[i]
		j := i - 1
		for j >= 0 && umkmList[j].Nama > key.Nama {
			umkmList[j+1] = umkmList[j]
			j = j - 1
		}
		umkmList[j+1] = key
	}
}
