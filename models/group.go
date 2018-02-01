package models

type Correspond struct {
	Correspond string `json:"correspond"`
	FieldName  string `json:"fieldname"`
}

type Grp struct {
	Grp    string `json:"group"`
	GrpEng string `json:"grpeng"`
}

type FiledGroup struct {
	Grp        string `json:"grp"`
	Correspond string `json:"correspond"`
	FieldName  string `json:"fieldname"`
	GrpEng     string `json:"grpeng"`
}
