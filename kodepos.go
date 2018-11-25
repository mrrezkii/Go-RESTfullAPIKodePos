package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Daerah struct {
	IDDaerah   string
	NamaDaerah string
}

var dataDaerah = []Daerah{
	Daerah{"A", "Kota Kediri"},
	Daerah{"B", "Kabupaten Kediri"},
}

type Kodepos struct {
	IDKodepos    int
	Daerah       string
	Kelurahan    string
	Desa         string
	NomorKodepos string
}

var dataKodepos = []Kodepos{
	Kodepos{1, "A", "Kediri Kota", "Semampir", "64121"},
	Kodepos{2, "A", "Kediri Kota", "Dandangan", "64122"},
	Kodepos{3, "A", "Kediri Kota", "Ngadirejo", "64122"},
	Kodepos{4, "A", "Kediri Kota", "Poncaran", "64123"},
	Kodepos{5, "A", "Kediri Kota", "Banjaran", "64123"},
	Kodepos{6, "A", "Kediri Kota", "Jagalan", "64124"},
	Kodepos{7, "A", "Kediri Kota", "Kemasan", "64125"},
	Kodepos{8, "A", "Kediri Kota", "Kaliombo", "64125"},
	Kodepos{9, "A", "Kediri Kota", "Kali Dalem", "64126"},
	Kodepos{10, "A", "Kediri Kota", "Ngronggo", "64126"},
	Kodepos{11, "A", "Kediri Kota", "Manisrenggo", "64127"},
	Kodepos{12, "A", "Mojoroto", "Bangsal", "64131"},
	Kodepos{13, "A", "Mojoroto", "Burengan", "64131"},
	Kodepos{14, "A", "Mojoroto", "Pesantren", "64131"},
	Kodepos{15, "A", "Mojoroto", "Jamsaren", "64132"},
	Kodepos{16, "A", "Mojoroto", "Pakunden", "64132"},
	Kodepos{17, "A", "Mojoroto", "Tinalan", "64134"},
	Kodepos{18, "A", "Mojoroto", "Banaran", "64134"},
	Kodepos{19, "A", "Mojoroto", "Tosaren", "64135"},
	Kodepos{20, "B", "Badas", "Krecek", "64135"},
	Kodepos{21, "B", "Badas", "Canggu", "64135"},
	Kodepos{22, "B", "Badas", "Lamong", "64135"},
	Kodepos{23, "B", "Banyakan", "Jabon", "64135"},
	Kodepos{24, "B", "Banyakan", "Jatirego", "64135"},
	Kodepos{25, "B", "Banyakan", "Maron", "64135"},
	Kodepos{26, "B", "Gampingrejo", "Kepuhrejo", "64135"},
	Kodepos{27, "B", "Gampingrejo", "Putih", "64135"},
	Kodepos{28, "B", "Grogol", "Grogol", "64135"},
	Kodepos{29, "B", "Grogol", "Kalimpang", "64135"},
	Kodepos{30, "B", "Kandangan", "Bukur", "64135"},
	Kodepos{31, "B", "Kandangan", "Banaran", "64135"},
	Kodepos{32, "B", "Kandat", "Kandat", "64135"},
	Kodepos{33, "B", "Kandat", "Pule", "64135"},
	Kodepos{34, "B", "Kras", "Banjaranyar", "64135"},
	Kodepos{35, "B", "Kras", "Bleber", "64136"},
	Kodepos{36, "B", "Kras", "Krandang", "6419"},
	Kodepos{37, "B", "Kras", "Purwodadi", "64150"},
	Kodepos{38, "B", "Kras", "Setonorejo", "64130"},
	Kodepos{39, "B", "Kras", "Pelas", "641334"},
}

func getDaerah(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(dataDaerah)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func getKodepos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(dataKodepos)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/Kodepos", getKodepos)
	http.HandleFunc("/Daerah", getDaerah)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
