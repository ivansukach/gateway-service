package repositories

type Endpoints struct {
	Version        int32  `db:"version"`
	AuthService    string `db:"authentication_service"`
	BookService    string `db:"book_service"`
	PingPong       string `db:"ping_pong"`
	ProfileService string `db:"profile_service"`
}
type Repository interface {
	GetNewGateways() (*Endpoints, error)
}
