package repositories

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func New(db *sqlx.DB) *gatewayRepository {
	return &gatewayRepository{db: db}
}

type gatewayRepository struct {
	db *sqlx.DB
}

func (gr *gatewayRepository) GetNewGateways() ([]Endpoints, error) {
	u := make([]Endpoints, 0)
	rows, err := gr.db.Queryx("SELECT * FROM gateways")
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for rows.Next() {
		record := new(Endpoints)
		err = rows.StructScan(&record)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		u = append(u, *record)
	}
	return u, err

}
