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
	UserId     string `json:"userId"`
	Date       string `json:"date"`
}

type MysqlConfig struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Configs struct {
	Configs []MysqlConfig `json:"mysql"`
}
