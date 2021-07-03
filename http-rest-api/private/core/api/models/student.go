package model

type Student struct {
	Id            int    `json:"id"`
	Surname       string `json:"surname"`
	MiddleName    string `json:"middleName"`
	Name          string `json:"name"`
	Date_b        string `json:"date_b"`
	City          string `json:"city"`
	Street        string `json:"street"`
	House         string `json:"house"`
	Flat          string `json:"flat"`
	Phone         string `json:"phone"`
	Info          string `json:"info"`
	NumGroup      string `json:"numGroup"`
	Activs        string `json:"activs"`
	Gender        string `json:"gender"`
	Status        string `json:"status"`
	Orphan        bool   `json:"orphan"`
	Invalid       bool   `json:"invalid"`
	Low_income    bool   `json:"low_income"`
	Low_parents   bool   `json:"low_parents"`
	Idn           bool   `json:"idn"`
	Kdn           bool   `json:"kdn"`
	Many_children bool   `json:"many_children"`
	Login         string `json:"login"`
	Password      string `json:"password"`
	Budget        string `json:"budget"`
}
