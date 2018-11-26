package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Kodepos struct {
	Id        int
	Daerah    string
	Kelurahan string
	Desa      string
	Kodepos   string
}
type Response struct {
	NextPage int
	Data     []Kodepos
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db_kodepos.db")
	if err != nil {
		return db, err
	}

	return db, err
}

func getKodepos(w http.ResponseWriter, r *http.Request) {

	db, errdb := connect()
	if errdb != nil {
		http.Error(w, errdb.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		pages, found := r.URL.Query()["page"]
		search, found := r.URL.Query()["search"]
		var pageNow int
		var kodeposs []Kodepos
		if found {
			nowPage, errPages := strconv.Atoi(pages[0])
			pageNow = nowPage
			if errPages != nil {
				http.Error(w, errPages.Error(), http.StatusInternalServerError)
				return
			}
			page := strconv.Itoa((nowPage - 1) * 10)
			rows, errq := db.Query("SELECT * FROM tb_kodepos WHERE kodepos LIKE '%" + search[0] + "%' ORDER BY tb_kodepos.kodepos ASC LIMIT 10 OFFSET " + page)
			if errq != nil {
				http.Error(w, errq.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			for rows.Next() {
				var kodepos = Kodepos{}
				var err = rows.Scan(&kodepos.Id, &kodepos.Desa, &kodepos.Kelurahan, &kodepos.Desa, &kodepos.Kodepos)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				kodeposs = append(kodeposs, kodepos)
			}
			var response Response
			response.Data = kodeposs

			if kodeposs == nil {
				response.NextPage = 1
			} else {
				response.NextPage = pageNow + 1
			}

			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/kodepos", getKodepos)

	fmt.Println("Starting Server at http://localhost:4848/")
	http.ListenAndServe(":4848", nil)
}
