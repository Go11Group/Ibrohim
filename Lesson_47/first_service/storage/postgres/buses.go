package postgres

import (
	"database/sql"
	"math/rand"
	pb "weather-transport/genproto/transport"

	"github.com/pkg/errors"
)

type BusRepo struct {
	DB *sql.DB
}

func NewBusRepo(db *sql.DB) *BusRepo {
	return &BusRepo{DB: db}
}

func (br *BusRepo) GetBusSchedule(n *pb.Number) (*pb.Schedule, error) {
	sch := &pb.Schedule{}

	query := `
	select monday, tuesday, wednesday, thursday, friday, saturday, sunday
	from schedule s
	join buses b on s.bus_id = b.id
	where b.number = $1
	group by s.monday, s.tuesday, s.wednesday, s.thursday, s.friday, s.saturday, s.sunday`
	err := br.DB.QueryRow(query, n.Number).Scan(&sch.Monday, &sch.Tuesday, &sch.Wednesday, &sch.Thursday, &sch.Friday, &sch.Saturday, &sch.Sunday)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get schedule")
	}

	return sch, nil
}

func (br *BusRepo) TrackBusLocation(n *pb.Number) (*pb.GPS, error) {
	return &pb.GPS{
		Latitude:  rand.Float32()*180 - 90,
		Longitude: rand.Float32()*360 - 180,
	}, nil
}

func (br *BusRepo) ReportTrafficJam(r *pb.Route) (*pb.Traffic, error) {
	t := &pb.Traffic{}

	if r.Transports > 100 {
		t.Level = "heavy"
	} else if r.Transports > 50 {
		t.Level = "moderate"
	} else {
		t.Level = "light"
	}

	return t, nil
}
