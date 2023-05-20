package model

type Person struct {
	Id           string   `json:"id"`
	Guid         string   `json:"guid"`
	UserName     string   `json:"user_name"`
	Mobile       string   `json:"mobile"`
	Address      string   `json:"address"`
	Brthday      string   `json:"brthday"`
	Email        string   `json:"email"`
	Weight       float64  `json:"weight"`
	Height       float64  `json:"height"`
	Gender       string   `json:"gender"`
	Introduce    string   `json:"introduce"`
	Motto        string   `json:"motto"`
	Genres       []string `json:"genres"`
	Professional string   `json:"professional"`
	Title        string   `json:"title"`
}
