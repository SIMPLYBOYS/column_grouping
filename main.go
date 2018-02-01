package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"column_grouping/controllers"
	"column_grouping/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var db *sql.DB
var gc *controllers.GroupController

// var err error

func init() {
	log.Println("=== init ===")
}

func main() {
	gc := controllers.NewGroupController()
	fmt.Println("=== main ===")
	mux := httprouter.New()
	mux.GET("/", helloworld)
	mux.GET("/getGrp", gc.GetGrp)
	mux.GET("/getFieldGrp", gc.GetFieldGrp)
	mux.GET("/getCorrespond", gc.GetCorrespond)
	mux.POST("/delGrp", gc.DelFiledGrp)
	mux.POST("/bulkDelGrp", gc.BulkDelFieldGrp)
	mux.POST("/bulkAddGrp", gc.BulkInsertFieldGrp)
	mux.POST("/addGrp", gc.InsertFieldGrp)
	fmt.Println("Play Golang Now ...")
	http.ListenAndServe(":"+models.HttpPort, mux)
}

func helloworld(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("--- helloworld ---")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Println("Hello, playground")
	foo_marshalled, _ := json.Marshal(`{name: hanter}`)
	fmt.Fprint(w, string(foo_marshalled))
}
