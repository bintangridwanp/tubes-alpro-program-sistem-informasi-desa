package main

const MAX int = 150

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

type StatusChangeRequest struct {
	NIK, Nama, Status string
}

var desaList [MAX]Desa
var pendudukList [MAX][MAX]Penduduk
var pendudukCount [MAX]int
var desaCount int
var statusChangeRequests [MAX]StatusChangeRequest
var statusChangeRequestCount int

func initDefaultData() {
	// Inisialisasi data desa
	desaList[0] = Desa{
		Provinsi:  "Jawa Barat",
		Kabupaten: "Bandung",
		Kecamatan: "Cilengkrang",
		Nama:      "Desa Mekarsari",
		JumlahRT:  10,
		JumlahRW:  5,
	}
	desaList[1] = Desa{
		Provinsi:  "Jawa Tengah",
		Kabupaten: "Sukoharjo",
		Kecamatan: "Baki",
		Nama:      "Desa Bakti",
		JumlahRT:  8,
		JumlahRW:  4,
	}
	desaCount = 2

	// Inisialisasi data penduduk untuk Desa Mekarsari
	pendudukList[0][0] = Penduduk{"3201010101010001", "Andi", "01-01-1990", "Kawin"}
	pendudukList[0][1] = Penduduk{"3201010101010002", "Budi", "02-02-1992", "Belum Kawin"}
	pendudukList[0][2] = Penduduk{"3201010101010003", "Cici", "03-03-1994", "Kawin"}
	pendudukList[0][3] = Penduduk{"3201010101010004", "Dodi", "04-04-1996", "Belum Kawin"}
	pendudukList[0][4] = Penduduk{"3201010101010005", "Edi", "05-05-1998", "Kawin"}
	pendudukCount[0] = 5

	// Inisialisasi data penduduk untuk Desa Bakti
	pendudukList[1][0] = Penduduk{"3302020202020001", "Fani", "06-06-1990", "Kawin"}
	pendudukList[1][1] = Penduduk{"3302020202020002", "Gita", "07-07-1992", "Belum Kawin"}
	pendudukList[1][2] = Penduduk{"3302020202020003", "Hadi", "08-08-1994", "Kawin"}
	pendudukList[1][3] = Penduduk{"3302020202020004", "Iwan", "09-09-1996", "Belum Kawin"}
	pendudukList[1][4] = Penduduk{"3302020202020005", "Joko", "10-10-1998", "Kawin"}
	pendudukCount[1] = 5
}
