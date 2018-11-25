package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type daerah struct {
	idDaerah   string
	namaDaerah string
}

var dataDaerah = []daerah{
	daerah{"A", "Kota Kediri"},
	daerah{"B", "Kabupaten Kediri"},
}

type kodepos struct {
	idKodePos    int
	idDaerah     string
	kelurahan    string
	desa         string
	nomorkodepos string
}

var dataKodepos = []kodepos{
	kodepos{1, "A", "Kediri Kota", "Semampir", "64121"},
	kodepos{2, "A", "Kediri Kota", "Dandangan", "64122"},
	kodepos{3, "A", "Kediri Kota", "Ngadirejo", "64122"},
	kodepos{4, "A", "Kediri Kota", "Poncaran", "64123"},
	kodepos{5, "A", "Kediri Kota", "Banjaran", "64123"},
	kodepos{6, "A", "Kediri Kota", "Jagalan", "64124"},
	kodepos{7, "A", "Kediri Kota", "Kemasan", "64125"},
	kodepos{8, "A", "Kediri Kota", "Kaliombo", "64125"},
	kodepos{9, "A", "Kediri Kota", "Kali Dalem", "64126"},
	kodepos{10, "A", "Kediri Kota", "Ngronggo", "64126"},
	kodepos{11, "A", "Kediri Kota", "Manisrenggo", "64127"},
	kodepos{12, "A", "Mojoroto", "Bangsal", "64131"},
	kodepos{13, "A", "Mojoroto", "Burengan", "64131"},
	kodepos{14, "A", "Mojoroto", "Pesantren", "64131"},
	kodepos{15, "A", "Mojoroto", "Jamsaren", "64132"},
	kodepos{16, "A", "Mojoroto", "Pakunden", "64132"},
	kodepos{17, "A", "Mojoroto", "Tinalan", "64134"},
	kodepos{18, "A", "Mojoroto", "Banaran", "64134"},
	kodepos{19, "A", "Mojoroto", "Tosaren", "64135"},
	kodepos{20, "B", "Badas", "Krecek", "64135"},
	kodepos{21, "B", "Badas", "Canggu", "64135"},
	kodepos{22, "B", "Badas", "Lamong", "64135"},
	kodepos{23, "B", "Banyakan", "Jabon", "64135"},
	kodepos{24, "B", "Banyakan", "Jatirego", "64135"},
	kodepos{25, "B", "Banyakan", "Maron", "64135"},
	kodepos{26, "B", "Gampingrejo", "Kepuhrejo", "64135"},
	kodepos{27, "B", "Gampingrejo", "Putih", "64135"},
	kodepos{28, "B", "Grogol", "Grogol", "64135"},
	kodepos{29, "B", "Grogol", "Kalimpang", "64135"},
	kodepos{30, "B", "Kandangan", "Bukur", "64135"},
	kodepos{31, "B", "Kandangan", "Banaran", "64135"},
	kodepos{32, "B", "Kandat", "Kandat", "64135"},
	kodepos{33, "B", "Kandat", "Pule", "64135"},
	kodepos{34, "B", "Kras", "Banjaranyar", "64135"},
	kodepos{35, "B", "Kras", "Bleber", "64136"},
	kodepos{36, "B", "Kras", "Krandang", "6419"},
	kodepos{37, "B", "Kras", "Purwodadi", "64150"},
	kodepos{38, "B", "Kras", "Setonorejo", "64130"},
	kodepos{39, "B", "Kras", "Pelas", "641334"},
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

func main() {
	fmt.Println("Hello World!")
}
