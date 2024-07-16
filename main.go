package main

import (
	"daftar-hadir-karyawan/repository"
	"fmt"
)

func main() {
	repo := repository.NewKaryawaneRepository()
	var pilihan int

	for {
		fmt.Println("----Menu-----")
		fmt.Println("1. Tambah Karyawan")
		fmt.Println("2. Perbaruhi Status Karyawan")
		fmt.Println("3. Hapus Karyawan")
		fmt.Println("4. Tampilkan Semua Karyawan")
		fmt.Println("5. Exit")
		fmt.Println("6. Find by Id Karyawan")

		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			var id int
			var name string
			var kehadiran bool
			fmt.Println("Masukan nilai ID Karyawan")
			fmt.Scanln(&id)
			fmt.Println("Masukan nama Karyawan")
			fmt.Scanln(&name)
			fmt.Println("Masukan nama kehadiran -> (true/false)")
			fmt.Scanln(&kehadiran)
			repo.Tambahkan(id, name, kehadiran)
		case 2:
			var id int
			var kehadiran bool
			fmt.Println("Masukan Nilai ID Karyawan")
			fmt.Scanln(&id)
			fmt.Println("Apakah karayawan hadir (true/false)")
			fmt.Scanln(&kehadiran)
			repo.Perbaruhi(id, kehadiran)
		case 3:
			var id int
			fmt.Println("Hapus ID karyawan")
			fmt.Scanln(&id)
			repo.Hapus(id)
		case 4:
			repo.Tampilkan()
		case 5:
			fmt.Println("Terimakasih Ok")
			return
		case 6:
			var id int
			fmt.Println("cari id ini :")
			fmt.Scanln(&id)
			nangkap := repo.FindById(id)
			repo.FindById(id)
			fmt.Println(nangkap)
		default:
			fmt.Println("sepertinya kamu sedang mabuk saat memasukan inputan, coba lagi dah")
		}
	}
}
