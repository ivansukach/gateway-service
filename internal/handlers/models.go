package handlers

type BookModel struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Genre         string `json:"genre"`
	Edition       string `json:"edition"`
	NumberOfPages int32  `json:"numberOfPages"`
	Year          int32  `json:"year"`
	Amount        int32  `json:"amount"`
	IsPopular     bool   `json:"isPopular"`
	InStock       bool   `json:"inStock"`
}

type ProfileModel struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Age        int32  `json:"age"`
	Gender     bool   `json:"gender"`
	HasAnyPets bool   `json:"hasAnyPets"`
	Employed   bool   `json:"employed"`
}
type UserModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenModel struct {
	AccessToken  string `json:"Authorization"`
	RefreshToken string `json:"RefreshToken"`
}
type ClaimsModel struct {
	Claims map[string]string `json:"claims"`
}
