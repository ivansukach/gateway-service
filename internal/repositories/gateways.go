package repositories

type Endpoints struct {
	Id      int32  `db:"id"`
	Version string `db:"version"`
	Name    string `db:"name"`
	Url     string `db:"url"`
}
type Repository interface {
	GetNewGateways() ([]*Endpoints, error)
	GetLatestId() int32
	SetLatestId(latestId int32)
}
