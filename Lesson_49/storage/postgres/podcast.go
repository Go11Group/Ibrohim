package postgres

import (
	"fmt"
	pb "podcast_service/genproto/podcasts"
	"time"

	"database/sql"
	"github.com/google/uuid"
)

type PodcastRepo struct {
	Db *sql.DB
}

func NewPodcastRepo(db *sql.DB) *PodcastRepo {
	return &PodcastRepo{Db: db}
}

func (p *PodcastRepo) CreatePodcast(podcast *pb.PodcastCreate) (*string, error) {
	tx, err := p.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	newId := uuid.NewString()
	query := `insert into podcasts (id, user_id, title, description, 
	status) values ($1,$2,$3,$4,$5)`

	res, err := tx.Exec(query, newId, podcast.UserId, podcast.Title,
		podcast.Description, podcast.Status)
	if err != nil {
		return nil, err
	}
	if n, _ := res.RowsAffected(); n <= 0 {
		return nil, fmt.Errorf("cannot created")
	}
	return &newId, nil
}

func (p *PodcastRepo) UpdatePodcast(poadcast *pb.PodcastUpdate) error {
	query := `
	update 
		podcasts
	set
	    user_id  = $1,
    	title = $2,
    	description = $3,
   	 	status = $4,
    	updated_at = $5
	where
	    id = $6
`
	tx, err := p.Db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec(query, poadcast.UserId, poadcast.Title, poadcast.Description, poadcast.Status,
		time.Now(), poadcast.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	ra, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if ra == 0 {
		return fmt.Errorf("no rows affected")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PodcastRepo) DeletePodcast(podcastId *pb.ID) (*pb.Void, error) {
	tx, err := p.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	query := `update podcasts set deleted_at = $1 where deleted_at 
	is null`
	res, err := tx.Exec(query, podcastId)
	if err != nil {
		return nil, err
	}
	if n, _ := res.RowsAffected(); n <= 0 {
		return nil, fmt.Errorf("nothing is deleted")
	}
	return &pb.Void{}, nil
}
func (p *PodcastRepo) GetPodcastById(in *pb.ID) (*pb.Podcast, error) {
	podcast := &pb.Podcast{Id: in.Id}

	query := `select user_id, title, description, created_at, updated_at
	from podcasts where id = $1 and deleted_at = null`

	err := p.Db.QueryRow(query, in.Id).Scan(&podcast.UserId, &podcast.Title, &podcast.Description, &podcast.CreatedAt, &podcast.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return podcast, nil
}

func (p *PodcastRepo) GetUserPodcasts(in *pb.ID) (*pb.UserPodcasts, error) {
	query := `select id, user_id, title, description, created_at, updated_at
	from podcasts where user_id = $1 and deleted_at = null`

	rows, err := p.Db.Query(query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var podcasts []*pb.Podcast
	for rows.Next() {
		var podcast pb.Podcast
		err := rows.Scan(&podcast.Id, &podcast.UserId, &podcast.Title, &podcast.Description, &podcast.CreatedAt, &podcast.UpdatedAt)
		if err != nil {
			return nil, err
		}
		podcasts = append(podcasts, &podcast)
	}

	return &pb.UserPodcasts{Podcasts: podcasts}, nil
}
