package client

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gfbatista/xy-inc/database"
	"github.com/gfbatista/xy-inc/model"
	"github.com/gfbatista/xy-inc/util"
	_ "github.com/go-sql-driver/mysql" //driver de conexão mysql
)

// ListByProximity Calcular as coordenadas X, Y e D e retornar os POIs correspondentes
func ListByProximity(w http.ResponseWriter, r *http.Request) {

	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))
	d, errD := strconv.Atoi(r.URL.Query().Get("d"))

	if errX != nil || errY != nil || errD != nil || x < 0 || y < 0 || d < 0 {
		w.WriteHeader(http.StatusNotFound)
		json, _ := json.Marshal(model.Erro{Descricao: "Envie os valores das coordenadas corretamente"})

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(json))
	}

	db, err := database.OpenConection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pois := getPoiList(db)

	newPois := []model.Poi{}
	for _, poi := range pois {
		vlCalc := calc(poi, float64(x), float64(y))
		if vlCalc <= float64(d) {
			newPois = append(newPois, poi)
		}
	}

	json, err := util.ObjToJSON(newPois)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

// ListInsert conforme a requisição, realiza o insert ou lista os registros
func ListInsert(w http.ResponseWriter, r *http.Request) {
	db, err := database.OpenConection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch {
	case r.Method == "POST":
		p := []model.Poi{}

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		insertPoi(db, p)

		w.WriteHeader(http.StatusCreated)
	case r.Method == "GET":
		pois := getPoiList(db)

		json, err := util.ObjToJSON(pois)

		if err != nil {
			fmt.Println("Erro ao converter para Json. Erro: ", err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(json))
	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		json, _ := json.Marshal(model.Erro{Descricao: "Método inexistente"})
		fmt.Fprint(w, string(json))
	}
}

func getPoiList(db *sql.DB) []model.Poi {
	rows, _ := db.Query("select id, nome, coordenada_x, coordenada_y from poi")
	defer rows.Close()

	pois := []model.Poi{}
	for rows.Next() {
		poi := model.Poi{}
		rows.Scan(&poi.ID, &poi.Nome, &poi.CoordenadaX, &poi.CoordenadaY)
		pois = append(pois, poi)
	}

	return pois
}

func insertPoi(db *sql.DB, pois []model.Poi) {

	stmt, err := db.Prepare("INSERT INTO poi (nome, coordenada_X, coordenada_y) VALUES( ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, poi := range pois {
		if _, err := stmt.Exec(poi.Nome, poi.CoordenadaX, poi.CoordenadaY); err != nil {
			log.Fatal(err)
		}
	}
}

func calc(poi model.Poi, x float64, y float64) float64 {

	return math.Sqrt(exponent(float64(poi.CoordenadaX)-x, 2) + exponent(float64(poi.CoordenadaY)-y, 2))
}

func exponent(x, y float64) float64 {

	return math.Pow(x, y)

}
