package postgres

import (
	pb "collaboration_service/genproto/comments"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type CommentRepo struct {
	Db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{Db: db}
}

func (c *CommentRepo) CreateEpisodeComment(comment *pb.EpisodeComment) (string, error) {
	tx, err := c.Db.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Commit()

	query := `
	insert into 
	    comments(
		id, episode_id, user_id, content	
	) values (
		$1, $2, $3, $4
	)`

	newId := uuid.NewString()
	_, err = tx.Exec(query, newId, comment.EpisodeId, comment.UserId,
		comment.Content)

	return newId, err
}

func (c *CommentRepo) GetCommentsByPodcastId(id *pb.ID) ([]*pb.CommentInfo, error) {
	query := `
	select
		user_id,
    	content,
    	created_at,
    	updated_at
	from 
	    comments
	where
	    podcast_id = $1
`
	comments := []*pb.CommentInfo{}
	rows, err := c.Db.Query(query, id.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment pb.CommentInfo
		err = rows.Scan(&comment.UserId, &comment.Content, comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (c *CommentRepo) CreateCommentByPodcastId(comment *pb.CreateComment) (*pb.ID, error) {
	tx, err := c.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	query := `
	insert into 
	    comments(
		id, podcast_id, user_id, content	
	) values (
		$1, $2, $3, $4
	)`

	newId := uuid.NewString()
	_, err = tx.Exec(query, newId, comment.PodcastId, comment.UserId,
		comment.Content)

	return &pb.ID{Id: newId}, err
}

func (c *CommentRepo) CountComments(filter *pb.CountFilter) (*pb.CommentCount, error) {
	query := `
	select 
		count(*)
	from
	    comments
	where
	    deleted_at IS NULL 
`
	paramCount := 0
	params := []interface{}{}
	if len(filter.PodcastId) == 32 {
		query += fmt.Sprintf(" and podcast_id = $%d", paramCount)
		paramCount++
		params = append(params, filter.PodcastId)
	}
	if len(filter.EpisodeId) == 32 {
		query += fmt.Sprintf(" and episode_id = $%d", paramCount)
		paramCount++
		params = append(params, filter.EpisodeId)
	}

	row, err := c.Db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	count := pb.CommentCount{}
	err = row.Scan(&count.Count)

	return &count, err
}
