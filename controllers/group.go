package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/SIMPLYBOYS/column_grouping/models"
	"github.com/SIMPLYBOYS/column_grouping/utils/mysql"
	"github.com/julienschmidt/httprouter"
)

type GroupController struct {
}

func NewGroupController() *GroupController {
	return &GroupController{}
}

func (gc GroupController) GetCorrespond(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	db, err := utils.GetOpenMysqlClient()
	check(err)
	defer db.Close()
	rows, err := db.Query(`SELECT DISTINCT CorrespondInterface, Field_Name FROM T_EA_DocsD WHERE CorrespondInterface <> ""`)
	check(err)

	var correspondList []models.Correspond

	for rows.Next() {
		var (
			CorrespondInterface string
			Field_Name          string
			foo                 models.Correspond
		)

		err := rows.Scan(&CorrespondInterface, &Field_Name)
		foo.Correspond = CorrespondInterface
		foo.FieldName = Field_Name
		check(err)
		fmt.Print(CorrespondInterface)
		correspondList = append(correspondList, foo)
	}

	result, _ := json.Marshal(correspondList)

	w.WriteHeader(http.StatusOK) //200
	w.Write(result)
}

func (gc GroupController) BulkInsertFieldGrp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("\n\n --- BulkInsertFieldGrp ---\n\n")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	db, err := utils.GetOpenMysqlClient()
	check(err)
	defer db.Close()

	grps := req.FormValue("groups")
	formBody := []byte(grps)
	groups := make([]models.FiledGroup, 0)
	grp_eng := req.FormValue("grpeng")
	json.Unmarshal(formBody, &groups)

	// -------------------------------

	for _, item := range groups {
		stmt, err := db.Prepare("INSERT column_group SET grp=?, CorrespondInterface=?, Field_Name=?, grp_eng=?, US_ID=?, date_time=?")
		check(err)
		defer stmt.Close()

		r, err := stmt.Exec(item.Grp, item.Correspond, item.FieldName, item.GrpEng, item.UserId, item.Date)
		check(err)

		n, err := r.RowsAffected()
		fmt.Println(n)
		check(err)
	}

	// -------------------------------

	stmt, err := db.Prepare("ALTER TABLE t_sys_role ADD " + strings.ToLower(grp_eng) + " INT NOT NULL DEFAULT 0")
	check(err)

	a, err := stmt.Exec()
	check(err)

	n2, err := a.RowsAffected()
	check(err)
	fmt.Println(n2)

	fmt.Fprintln(w, "Bulk INSERT Grp Success!", nil)

}

func (gc GroupController) InsertFieldGrp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("--- InsertFieldGrp ---")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	db, err := utils.GetOpenMysqlClient()
	check(err)
	defer db.Close()

	grp := req.FormValue("grp")
	correspond := req.FormValue("correspond")
	field := req.FormValue("field")
	grp_eng := req.FormValue("grpEng")
	userId := req.FormValue("userId")
	date := req.FormValue("date")

	// rows, err := db.Query("SELECT * FROM column_group WHERE CorrespondInterface =" + correspond + " AND Field_Name = " + field)
	// fmt.Println(rows)

	stmt, err := db.Prepare("INSERT column_group SET grp=?, CorrespondInterface=?, Field_Name=?, grp_eng=?, US_ID=?, date_time=?")
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec(grp, correspond, field, grp_eng, userId, date)
	check(err)

	n, err := r.RowsAffected()
	fmt.Println(n)
	check(err)

	rows, err := db.Query("SELECT " + strings.ToLower(grp_eng) + " FROM t_sys_role")
	if rows == nil {
		stmt, err = db.Prepare("ALTER TABLE t_sys_role ADD " + strings.ToLower(grp_eng) + " INT NOT NULL DEFAULT 0")
		check(err)

		a, err := stmt.Exec()
		check(err)

		n2, err := a.RowsAffected()
		check(err)
		fmt.Println(n2)
	} else {
		fmt.Println("Mow")
	}

	fmt.Fprintln(w, "INSERT Grp Success!", n)
}

func (gc GroupController) BulkDelFieldGrp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("--- BulkDelFieldGrp ---")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	db, err := utils.GetOpenMysqlClient()
	// check(err)
	defer db.Close()
	grp := req.FormValue("grp")
	grpEng := req.FormValue("grpEng")
	stmt, err := db.Prepare("DELETE FROM column_group WHERE grp=?")
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec(grp)
	check(err)

	n, err := r.RowsAffected()
	check(err)

	stmt, err = db.Prepare("ALTER TABLE t_sys_role DROP COLUMN " + grpEng)
	check(err)

	a, err := stmt.Exec()
	check(err)

	_, err = a.RowsAffected()

	fmt.Fprintln(w, "Bulk DELETE Grp Success!", n)
}

func (gc GroupController) DelFiledGrp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Println("--- DelFiledGrp ---")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	db, err := utils.GetOpenMysqlClient()
	// check(err)
	defer db.Close()
	grp := req.FormValue("grp")
	grpEng := req.FormValue("grpEng")

	stmt, err := db.Prepare("DELETE FROM column_group WHERE grp=? AND grp_eng=?")
	check(err)
	defer stmt.Close()

	r, err := stmt.Exec(grp, grpEng)
	check(err)

	n, err := r.RowsAffected()
	check(err)

	// fmt.Println(grpEng)

	rows, err := db.Query("SELECT " + grpEng + " FROM t_sys_role")
	// fmt.Print(rows)
	check(err)

	if rows != nil {
		stmt, err = db.Prepare("ALTER TABLE t_sys_role DROP COLUMN " + grpEng)
		check(err)

		a, err := stmt.Exec()
		check(err)

		_, err = a.RowsAffected()
	} else {
		fmt.Println("Mow")
	}

	fmt.Fprintln(w, "DELETE Grp Success!", n)
}

func (gc GroupController) GetFieldGrp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Println("\n\n --- GetFieldGrp --- \n\n")
	db, err := utils.GetOpenMysqlClient()
	check(err)
	defer db.Close()
	rows, err := db.Query(`SELECT grp, CorrespondInterface, Field_Name, grp_eng FROM column_group`)
	check(err)

	var fieldGrpList []models.FiledGroup

	for rows.Next() {
		var (
			grp                 string
			CorrespondInterface string
			Field_Name          string
			grp_eng             string
			foo                 models.FiledGroup
		)

		err = rows.Scan(&grp, &CorrespondInterface, &Field_Name, &grp_eng)
		check(err)
		fmt.Print("\n\n ==== rows.Scan ==== \n")
		fmt.Println(grp)
		fmt.Println(CorrespondInterface)
		foo.Correspond = CorrespondInterface
		foo.Grp = grp
		foo.FieldName = Field_Name
		foo.GrpEng = grp_eng
		fieldGrpList = append(fieldGrpList, foo)
	}

	fmt.Print(rows)

	result, _ := json.Marshal(fieldGrpList)

	w.WriteHeader(http.StatusOK) //200
	w.Write(result)
}

func (gc GroupController) GetGrp(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	db, err := utils.GetOpenMysqlClient()
	check(err)
	defer db.Close()
	rows, err := db.Query("SELECT DISTINCT grp, grp_eng FROM column_group")
	check(err)

	var grpList []models.Grp

	for rows.Next() {
		var (
			grp     string
			grp_eng string
		)
		var foo models.Grp
		err = rows.Scan(&grp, &grp_eng)
		foo.Grp = strings.Trim(grp, " ")
		foo.GrpEng = grp_eng
		grpList = append(grpList, foo)
		check(err)
		fmt.Println(grp)
	}

	result, _ := json.Marshal(grpList)
	// fmt.Fprintln(w, "RETRIEVED RECORD:", result)
	w.WriteHeader(http.StatusOK) //200
	w.Write(result)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
