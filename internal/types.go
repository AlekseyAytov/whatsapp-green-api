package internal

type AuthData struct {
	ID     string `json:"idInstance"`
	Token  string `json:"apiTokenInstance"`
	Number string `json:"number"`
	Text   string `json:"text"`
	File   string `json:"file"`
}
