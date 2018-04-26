package setrg

import (
	"net/http"
	"strconv"
)

type mahasiswa struct {
	no      int
	nama    string
	umur    int
	jurusan string
}

type SetrgModule struct {
	something []mahasiswa
}

func NewSetrgModule() *SetrgModule {

	return &SetrgModule{
		something: []mahasiswa{
			mahasiswa{
				no:      1,
				nama:    "Andi",
				umur:    12,
				jurusan: "Matematika",
			},
			mahasiswa{
				no:      2,
				nama:    "Anton",
				umur:    13,
				jurusan: "Teknik Informatika",
			},
			mahasiswa{
				no:      3,
				nama:    "Budi",
				umur:    17,
				jurusan: "Sistem Informasi",
			},
			mahasiswa{
				no:      4,
				nama:    "Calvin",
				umur:    23,
				jurusan: "Teknik Informatika",
			},
			mahasiswa{
				no:      5,
				nama:    "Dennis",
				umur:    12,
				jurusan: "Matematika",
			},
		},
	}

}

func (hlm *SetrgModule) ManusiaGanjil(w http.ResponseWriter, r *http.Request) {
	for _, v := range hlm.something {
		if v.no%2 != 0 {
			w.Write([]byte(strconv.Itoa(v.no) + ". Nama: " + v.nama + " umur: " + strconv.Itoa(v.umur) + "\n"))
		}
	}

}

func (hlm *SetrgModule) ProgrammerMuda(w http.ResponseWriter, r *http.Request) {
	for _, v := range hlm.something {
		if v.umur < 14 && v.jurusan == "Teknik Informatika" {
			w.Write([]byte(strconv.Itoa(v.no) + ". Nama: " + v.nama + " umur: " + strconv.Itoa(v.umur) + "\n"))
		}
	}

}
