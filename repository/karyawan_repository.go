package repository

import (
	"daftar-hadir-karyawan/model"
	"fmt"
)

type KaryawanRepository struct {
	Karyawan []model.Karyawan
}

func NewKaryawaneRepository() *KaryawanRepository {
	return &KaryawanRepository{
		Karyawan: []model.Karyawan{},
	}
}

func (repo *KaryawanRepository) Tambahkan(id int, name string, kehadiran bool) {
	repo.Karyawan = append(repo.Karyawan, model.Karyawan{Id: id, Nama: name, Kehadiran: kehadiran})
}

func (repo *KaryawanRepository) Perbaruhi(id int, kehadiran bool) {
	for i := range repo.Karyawan {
		if repo.Karyawan[i].Id == id {
			repo.Karyawan[i].Kehadiran = kehadiran
			break
		}
	}
}

func (repo *KaryawanRepository) Hapus(id int) {
	for i := range repo.Karyawan {
		if repo.Karyawan[i].Id == id {
			repo.Karyawan = append(repo.Karyawan[:i], repo.Karyawan[1+i:]...)
			
			break
		}
	}

	for i := range repo.Karyawan {
		repo.Karyawan[i].Id = i + 1
	}
}

func (repo *KaryawanRepository) Tampilkan() {
	fmt.Println("=================== Data Karyawan ============================")
	for _, k := range repo.Karyawan {
		fmt.Printf("ID : %d , Nama: %s , Kehadiran: %t\n", k.Id, k.Nama, k.Kehadiran)
	}
}


func (repo *KaryawanRepository) FindById(id int) *model.Karyawan {
	for _ , value := range repo.Karyawan {
		if value.Id == id {
			return &value
		}else{
			return nil
		}
	}

	return nil
}
// func (repo *KaryawanRepository) FindById(id int, karyawan *model.Karyawan){
// 	for _ , value := range repo.Karyawan {
// 		if value.Id == id {
// 			*karyawan == &value
			
// 		}
// 	}
// }