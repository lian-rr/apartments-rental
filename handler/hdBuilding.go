package handler

type Building struct {
	ID          int    `json:"id, omitempty"`
	sName       string `json:"shortName"`
	fName       string `json:"fullName"`
	addr        string `json:"address"`
	phone       string `json:"phone"`
	mng         string `json:"manager"`
	description string `json:"description"`
}
