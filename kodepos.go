package main

import "fmt"

type kecamatan struct {
	idKecamatan     string
	kecamatanDaerah string
}

var dataKecamatan = []kecamatan{
	kecamatan{"A", "Kota Kediri"},
	kecamatan{"B", "Kabupaten Kediri"},
}

type kodepos struct {
	idKodePos   int
	idKecamatan string
	kelurahan   string
	kodepos     string
}

func main() {
	fmt.Println("Hello World!")
}
