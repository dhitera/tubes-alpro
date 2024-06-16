package main

import (
	"fmt"
	"strings"
)

// RANAH DEKLARASI GLOBAL DAN STRUCT

type camaba struct {
	nama, jurusan, asal, status string
	nohp, nilai int
}

type jurusan struct {
	nama, kode string
}

type tabCamaba = []camaba
type tabJurusan = []jurusan

// RANAH FUNGSI-FUNGSI

func dataJurusan() tabJurusan {
	return tabJurusan{
		{"Informatika", "IF"},
		{"Sistem Informasi", "SI"},
		{"Teknik Elektro", "TE"},
		{"Teknik Telekomunikasi", "TT"},
		{"Desain Komunikasi Visual", "DKV"},
		{"Rekayasa Perangkat Lunak", "RPL"},
		{"Teknologi Informasi", "TI"},
	}
}

func dataCamaba() tabCamaba {
	return tabCamaba{
		{"Radithya", "IF", "Jakarta", "Lulus", 01234567, 90},
		{"Haris", "IF", "Medan", "Lulus", 0121212, 77},
		{"John", "SE", "Bandung", "Ditolak", 010101022, 35},
	}
}

// 		FUNGSI CRUD
// FUNGSI INPUT DATA MAHASISWA DAN JURUSAN

func inputData(TAB *tabCamaba, role string) {
	
	var i camaba
	fmt.Print("Masukkan Nama Lengkap: ")
	fmt.Scan(&i.nama)
	fmt.Print("Masukkan Jurusan anda: ")
	fmt.Scan(&i.jurusan)
	fmt.Print("Masukkan Daerah Asal anda: ")
	fmt.Scan(&i.asal)
	fmt.Print("Masukkan No. Handphone anda: ")
	fmt.Scan(&i.nohp)
	if role == "adm" {
		fmt.Print("Input Nilai: ")
		fmt.Scan(&i.nilai)
		if i.nilai >= 75 {
			i.status = "Lulus"
		}else {
			i.status = "Ditolak"
		}
	}else if role == "cmhs" {
		i.status = "Belum Diputuskan"
		i.nilai = 0
	}
	*TAB = append(*TAB, i)
	fmt.Println("Data berhasil ditambahkan!")
}

func inputDataJurusan(TAB *tabJurusan) {
	
	var i jurusan
	fmt.Print("Masukkan Nama Jurusan: ")
	fmt.Scan(&i.nama)
	fmt.Print("Masukkan Kode Jurusan: ")
	fmt.Scan(&i.kode)
	*TAB = append(*TAB, i)
	fmt.Println("Data berhasil ditambahkan!")
}

// FUNGSI TAMPILKAN DATA MAHASISWA DAN JURUSAN

func tampilkanData(TAB tabCamaba) {
	fmt.Printf("%-3s %-20s %-20s %-15s %-15s %-15s %-20s\n","NO", "NAMA", "JURUSAN", "DAERAH ASAL", "NO.HP", "NILAI TES", "STATUS")
		for i := 0; i < len(TAB); i++ {
			fmt.Printf("%-3d %-20v %-20v %-15v %-15v %-15v %-20v\n", i+1, TAB[i].nama, TAB[i].jurusan, TAB[i].asal, TAB[i].nohp, TAB[i].nilai, TAB[i].status)
		}
		fmt.Println("==================================\n")
}

func tampilkanDataJurusan(TAB tabJurusan) {
	fmt.Printf("%-3s %-35s %-20s\n","NO", "NAMA", "KODE JURUSAN")
		for i := 0; i < len(TAB); i++ {
			fmt.Printf("%-3d %-35v %-20v\n", i+1, TAB[i].nama, TAB[i].kode)
		}
		fmt.Println("==================================\n")
}

// FUNGSI UBAH DATA MAHASISWA DAN JURUSAN

func ubahData(TAB tabCamaba) {
	var i camaba
	var pilihData int
	var namaTabel string
	fmt.Println("=====> Data Calon Mahasiswa <=====")
	if len(TAB) > 0 {
		fmt.Printf("%-3s %-20s %-20s %-15s %-15s %-15s %-20s\n","NO", "NAMA", "JURUSAN", "DAERAH ASAL", "NO.HP", "NILAI TES", "STATUS")
		for i := 0; i < len(TAB); i++ {
			fmt.Printf("%-3d %-20v %-20v %-15v %-15v %-15v %-20v\n", i+1, TAB[i].nama, TAB[i].jurusan, TAB[i].asal, TAB[i].nohp, TAB[i].nilai, TAB[i].status)
		}
		fmt.Println("==================================\n")
		fmt.Print("Data mana yang ingin diubah(no): ")
		fmt.Scan(&pilihData)
		pilihData--
		if pilihData < 0 || pilihData >= len(TAB) {
			fmt.Println("Nomor tidak ada")
			return
		}
		fmt.Printf("%-20s %-20s %-15s %-15s %-15s %-20s\n","NAMA", "JURUSAN", "DAERAH ASAL", "NO.HP", "NILAI TES", "STATUS")
		fmt.Printf("%-20v %-20v %-15v %-15v %-15v %-20v\n", TAB[pilihData].nama, TAB[pilihData].jurusan, TAB[pilihData].asal, TAB[pilihData].nohp, TAB[pilihData].nilai, TAB[pilihData].status)
		fmt.Println()
		fmt.Println("=> Ketik (-) untuk Ubah Semua\n=> Tuliskan Nama Tabel jika ingin ubah spesifik\n(nama, jurusan, asal, nohp, nilai)")
		fmt.Print("Pilih: ")
		fmt.Scan(&namaTabel)
		if namaTabel == "-" {
			fmt.Print("Ubah Nama: ")
			fmt.Scan(&i.nama)
			fmt.Print("Ubah Jurusan: ")
			fmt.Scan(&i.jurusan)
			fmt.Print("Ubah Daerah Asal: ")
			fmt.Scan(&i.asal)
			fmt.Print("Ubah No. Handphone: ")
			fmt.Scan(&i.nohp)
			fmt.Print("Ubah Nilai: ")
			fmt.Scan(&i.nilai)
			TAB[pilihData] = i
		} else {
			switch strings.ToLower(namaTabel) {
			case "nama":
				fmt.Print("Ubah Nama: ")
				fmt.Scan(&TAB[pilihData].nama)
			case "jurusan":
				fmt.Print("Ubah Jurusan: ")
				fmt.Scan(&TAB[pilihData].jurusan)
			case "asal":
				fmt.Print("Ubah Daerah Asal: ")
				fmt.Scan(&TAB[pilihData].asal)
			case "nohp":
				fmt.Print("Ubah No. Handphone: ")
				fmt.Scan(&TAB[pilihData].nohp)
			case "nilai":
				fmt.Print("Ubah Nilai: ")
				fmt.Scan(&TAB[pilihData].nilai)
			default:
				fmt.Println("Nama tabel tidak valid!")
				return
			}
		}

		if TAB[pilihData].nilai >= 75 {
			TAB[pilihData].status = "Lulus"
		} else {
			TAB[pilihData].status = "Ditolak"
		}

		fmt.Println("Data berhasil Diubah!\n")

	} else {
		fmt.Println("\nTidak ada Data!\n")
		fmt.Println("==================================\n")
	}
}

func ubahDataJurusan(TAB tabJurusan) {
	var i jurusan
	var pilihData int
	var namaTabel string
	fmt.Println("=====> Data Calon Mahasiswa <=====")
	if len(TAB) > 0 {
		fmt.Printf("%-3s %-35s %-20s\n","NO", "NAMA", "KODE JURUSAN")
		for i := 0; i < len(TAB); i++ {
			fmt.Printf("%-3d %-35v %-20v\n", i+1, TAB[i].nama, TAB[i].kode)
		}
		fmt.Println("==================================\n")
		fmt.Print("Data mana yang ingin diubah(no): ")
		fmt.Scan(&pilihData)
		pilihData--
		if pilihData < 0 || pilihData >= len(TAB) {
			fmt.Println("Nomor tidak ada")
			return
		}
		fmt.Printf("%-3s %-35s %-20s\n","NO", "NAMA", "KODE JURUSAN")
		fmt.Printf("%-35v %-20v\n", TAB[pilihData].nama, TAB[pilihData].kode)
		fmt.Println()
		fmt.Println("=> Ketik (-) untuk Ubah Semua\n=> Tuliskan Nama Tabel jika ingin ubah spesifik\n(nama, jurusan, asal, nohp, nilai)")
		fmt.Print("Pilih: ")
		fmt.Scan(&namaTabel)
		if namaTabel == "-" {
			fmt.Print("Ubah Nama: ")
			fmt.Scan(&i.nama)
			fmt.Print("Ubah Kode: ")
			fmt.Scan(&i.kode)
			TAB[pilihData] = i
		} else {
			switch strings.ToLower(namaTabel) {
			case "nama":
				fmt.Print("Ubah Nama: ")
				fmt.Scan(&TAB[pilihData].nama)
			case "jurusan":
				fmt.Print("Ubah Kode: ")
				fmt.Scan(&TAB[pilihData].kode)
			default:
				fmt.Println("Nama tabel tidak valid!")
				return
			}
		}

		fmt.Println("Data berhasil Diubah!\n")

	} else {
		fmt.Println("\nTidak ada Data!\n")
		fmt.Println("==================================\n")
	}
}

// FUNGSI HAPUS DATA MAHASISWA DAN JURUSAN

func hapusData(TAB *tabCamaba) {
	var pilihData int
	if len(*TAB) > 0 {
		fmt.Print("Data mana yang ingin dihapus(no): ")
		fmt.Scan(&pilihData)
		pilihData--
		if pilihData < 0 || pilihData >= len(*TAB) {
			fmt.Println("Nomor tidak ada")
			return
		}
		*TAB = append((*TAB)[:pilihData], (*TAB)[pilihData+1:]...)
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("\nTidak ada Data!\n")
		fmt.Println("==================================\n")
	}
}

func hapusDataJurusan(TAB *tabJurusan) {
	var pilihData int
	if len(*TAB) > 0 {
		fmt.Print("Data mana yang ingin dihapus(no): ")
		fmt.Scan(&pilihData)
		pilihData--
		if pilihData < 0 || pilihData >= len(*TAB) {
			fmt.Println("Nomor tidak ada")
			return
		}
		*TAB = append((*TAB)[:pilihData], (*TAB)[pilihData+1:]...)
		fmt.Println("Data berhasil dihapus!")
	} else {
		fmt.Println("\nTidak ada Data!\n")
		fmt.Println("==================================\n")
	}
}
// 		FUNGSI SEARCHING

func cariDataJurusan(TAB tabCamaba, x string) {
	var jumDat int
	ketemu := false
	fmt.Printf("=====> Data Jurusan (%s) <=====", x)
	fmt.Println()

	if ketemu {
		fmt.Println("Data tidak ditemukan!")
	}else {
		fmt.Printf("%-3s %-20s %-20s %-15s %-15s %-15s %-20s\n","NO", "NAMA", "JURUSAN", "DAERAH ASAL", "NO.HP", "NILAI TES", "STATUS")
		for i := 0; i < len(TAB); i++ {
			if TAB[i].jurusan == x {
				fmt.Printf("%-3d %-20v %-20v %-15v %-15v %-15v %-20v\n", i+1, TAB[i].nama, TAB[i].jurusan, TAB[i].asal, TAB[i].nohp, TAB[i].nilai, TAB[i].status)
				ketemu = true
				jumDat++
			}
		}
	}

	fmt.Printf("\nTotal Data berjumlah: %d", jumDat)
	fmt.Println()
	fmt.Println("==================================")
}

func cariDataKelulusan(TAB tabCamaba, x string) {
	var jumDat int
	ketemu := false
	fmt.Printf("=====> Data %s <=====", x)
	fmt.Println()

	if ketemu {
		fmt.Println("Data tidak ditemukan!")
	}else {
		fmt.Printf("%-3s %-20s %-20s %-15s %-15s %-15s %-20s\n","NO", "NAMA", "JURUSAN", "DAERAH ASAL", "NO.HP", "NILAI TES", "STATUS")
		for i := 0; i < len(TAB); i++ {
			if TAB[i].status == x {
				fmt.Printf("%-3d %-20v %-20v %-15v %-15v %-15v %-20v\n", i+1, TAB[i].nama, TAB[i].jurusan, TAB[i].asal, TAB[i].nohp, TAB[i].nilai, TAB[i].status)
				ketemu = true
				jumDat++
			}
		}
	}

	fmt.Printf("\nTotal Data berjumlah: %d", jumDat)
	fmt.Println()
	fmt.Println("==================================")
}

// FUNGSI SORTING

func urutkanData() {
	
}

// FUNGSI TAMPILKAN DATA AND SEARCH DATA MAHASISWA DAN JURUSAN

func tampilDanCariData(TAB tabCamaba) {
	var pilihan, pilihanCari int
	var jurusan string
	fmt.Println("=====> Data Calon Mahasiswa <=====")
	if len(TAB) > 0 {
		tampilkanData(TAB)
		fmt.Println("Pilih Opsi!\n 1. Cari Data\n 2. Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			fmt.Printf("\x1bc")
			fmt.Println("Tampilkan Data Berdasarkan?\n 1. Jurusan\n 2. Status(Lulus)\n 3. Status(Ditolak)")
			fmt.Print("Pilih: ")
			fmt.Scan(&pilihanCari)
		if pilihanCari == 1 {
			fmt.Print("Jurusan apa yang ingin dicari?: ")
			fmt.Scan(&jurusan)
			cariDataJurusan(TAB, jurusan)
			fmt.Print("Kembali?(y): ")
			fmt.Scan(&pilihan)
		}else if pilihanCari == 2 {
			cariDataKelulusan(TAB, "Lulus")
			fmt.Print("Kembali?(y): ")
			fmt.Scan(&pilihan)
			}else if pilihanCari == 3 {
				cariDataKelulusan(TAB, "Ditolak")
				fmt.Print("Kembali?(y): ")
				fmt.Scan(&pilihan)
		}
		}else {
			fmt.Printf("\x1bc")
		}
	} else {
		fmt.Println("\nData masih kosong!\n")
		fmt.Println("==================================\n")
	}
}

func tampilDanCariDataJurusan(TAB tabJurusan) {
	var pilihan int
	fmt.Println("=====> Data Jurusan <=====")
	if len(TAB) > 0 {
		tampilkanDataJurusan(TAB)
		fmt.Print("Kembali?(y): ")
		fmt.Scan(&pilihan)
	} else {
		fmt.Println("\nData masih kosong!\n")
		fmt.Println("==================================\n")
	}
}

// RANAH MENU

func mainMenu(TABM *tabCamaba, TABJ *tabJurusan) {
	fmt.Printf("\x1bc")
	var pilihan int
	var pass string
	// var tab tabCamaba

	fmt.Println("---------------------")
	fmt.Println("      Menu Utama     ")
	fmt.Println("---------------------")
	fmt.Println("1. Calon Mahasiswa\n2. Admin\n3. Keluar Aplikasi")

	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)
	switch pilihan {
	case 1:
		camabaMenu(TABM, TABJ)
	case 2:
		fmt.Print("Masukkan Password: ")
		fmt.Scan(&pass)
		if pass == "admin" {
			adminMenu(TABM, TABJ)
		} else {
			fmt.Println("Maaf Password salah!")
			mainMenu(TABM, TABJ)
		}
	case 3:
		fmt.Println("Keluar Aplikasi...")
	default:
		fmt.Println("Input salah!")
		mainMenu(TABM, TABJ)
	}
}

func camabaMenu(TABM *tabCamaba, TABJ *tabJurusan) {
	fmt.Printf("\x1bc")
	var pilihan, mabaOrJurusan int
	for {
		fmt.Println("---------------------")
		fmt.Println("    Menu Mahasiswa   ")
		fmt.Println("---------------------")
		fmt.Println("1. Tampilkan Data\n2. Tambah Data\n3. Ubah Data\n4. Keluar Menu")


		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Printf("\x1bc")
			fmt.Println("1. Tampilkan Data Calon Mahasiswa\n2. Tampilkan Data Jurusan")
			fmt.Print("Pilih: ")
			fmt.Scan(&mabaOrJurusan)
			if mabaOrJurusan == 1 {
				tampilDanCariData(*TABM)
			}else if mabaOrJurusan == 2 {
				tampilDanCariDataJurusan(*TABJ)
			}
		case 2:
			fmt.Printf("\x1bc")
			inputData(TABM, "cmhs")
		case 3:
			fmt.Printf("\x1bc")
			fmt.Println("Ubah Data")
        case 4:
            mainMenu(TABM, TABJ)
            return
        default:
            fmt.Println("Input salah!")
		}
	}
}

func adminMenu(TABM *tabCamaba, TABJ *tabJurusan) {
	fmt.Printf("\x1bc")
	var pilihan, mabaOrJurusan int
	for {
		fmt.Println("---------------------")
		fmt.Println("      Menu Admin     ")
		fmt.Println("---------------------")
		fmt.Println("1. Tampilkan Data\n2. Tambah Data\n3. Ubah Data\n4. Hapus Data\n5. Keluar Menu")


		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Printf("\x1bc")
			fmt.Println("1. Tampilkan Data Calon Mahasiswa\n2. Tampilkan Data Jurusan")
			fmt.Print("Pilih: ")
			fmt.Scan(&mabaOrJurusan)
			if mabaOrJurusan == 1 {
				tampilDanCariData(*TABM)
			}else if mabaOrJurusan == 2 {
				tampilDanCariDataJurusan(*TABJ)
			}
		case 2:
			fmt.Printf("\x1bc")
			fmt.Println("1. Tambah Data Calon Mahasiswa\n2. Tambah Data Jurusan")
			fmt.Print("Pilih: ")
			fmt.Scan(&mabaOrJurusan)
			if mabaOrJurusan == 1 {	
				inputData(TABM, "adm")
				}else if mabaOrJurusan == 2 {
				inputDataJurusan(TABJ)

			}
		case 3:
			fmt.Printf("\x1bc")
			fmt.Println("1. Ubah Data Calon Mahasiswa\n2. Ubah Data Jurusan")
			fmt.Print("Pilih: ")
			fmt.Scan(&mabaOrJurusan)
			if mabaOrJurusan == 1 {	
				ubahData(*TABM)
			}else if mabaOrJurusan == 2 {
				ubahDataJurusan(*TABJ)

			}
		case 4:
			fmt.Printf("\x1bc")
			fmt.Println("1. Hapus Data Calon Mahasiswa\n2. Hapus Data Jurusan")
			fmt.Print("Pilih: ")
			fmt.Scan(&mabaOrJurusan)
			if mabaOrJurusan == 1 {	
				tampilkanData(*TABM)
				hapusData(TABM)
			}else if mabaOrJurusan == 2 {
				tampilkanDataJurusan(*TABJ)
				hapusDataJurusan(TABJ)
			}
		case 5:
            mainMenu(TABM, TABJ)
            return
        default:
            fmt.Println("Input salah!")
		}
	}
}

func main() {
	fmt.Printf("\x1bc")
	var tabMaba tabCamaba
	var tabJurusan tabJurusan
	tabMaba = append(tabMaba, dataCamaba()...)
	tabJurusan = append(tabJurusan, dataJurusan()...)
	mainMenu(&tabMaba, &tabJurusan)
}