package postgres

import (
	pb "collaboration_service/genproto/collaborations"
	"database/sql"

	"github.com/google/uuid"
)

type CollaborationRepo struct {
	Db *sql.DB
}

func NewCollaborationRepo(db *sql.DB) *CollaborationRepo {
	return &CollaborationRepo{Db: db}
}

func (c *CollaborationRepo) CreateCollaboration(collab *pb.CreateCollaboration) (string, error) {
	query := `
	insert into 
	  collaborations (id, podcast_id, user_id)
	values ($1, $2, $3)
	`
	tx, err := c.Db.Begin()
	if err != nil {
		return "", err
	}
	id := uuid.NewString()
	_, err = tx.Exec(query, id, collab.PodcastId, collab.UserId)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *CollaborationRepo) GetCollaboratorsByPodcastId(PodcastId string) ([]*pb.CollaboratorToGet, error) {
	query := `select user_id, role, created_at from collaborations
	where podcast_id = $1`

	collabrators := []*pb.CollaboratorToGet{}
	rows, err := c.Db.Query(query, PodcastId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c pb.CollaboratorToGet
		err := rows.Scan(&c.UserId, &c.Role, &c.JoinedAt)
		if err != nil {
			return nil, err
		}
		collabrators = append(collabrators, &c)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return collabrators, nil
}

func (c *CollaborationRepo) UpdateCollaboratorByPodcastId(clb *pb.UpdateCollaborator) error {

	query := `
	update 
		set
		   podcast_id = $1,
		  user_id = $2, 
		  role = $3
	from
		collaborations
	where
		id = $4
  `
	tx, err := c.Db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(query, clb.PodcastId, clb.UserId, clb.Role, clb.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (c *CollaborationRepo) DeleteCollaboratorByPodcastId(ids *pb.Ids) (*pb.Void, error) {
	query := `delete from collaborations
	where podcast_id = $1 and user_id = $2`

	tr, err := c.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = tr.Exec(query, ids.PodcastId, ids.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.Void{}, nil
}

func (c *CollaborationRepo) CreateInvitation(invitation *pb.CreateInvite) (string, error) {
	query := `
	insert into 
	  invitations (id, podcast_id, inviter_id, invitee_id)
	values ($1, $2, $3, $4)
	`
	tx, err := c.Db.Begin()
	if err != nil {
		return "", err
	}
	id := uuid.NewString()
	_, err = tx.Exec(query, id, invitation.PodcastId, invitation.InviterId, invitation.InviteeId)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *CollaborationRepo) RespondInvitation(collab *pb.CreateCollaboration) (*pb.ID, error) {
	query := `
	update 
		invitations
	set 
		status = $1
	where 
		id = $2 and deleted_at = null`
	params := []interface{}{collab.Status, collab.InvitationId}

	tr, err := c.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = tr.Exec(query, params...)
	if err != nil {
		return nil, err
	}
	if collab.Status == "accepted" {
		id, err := c.CreateCollaboration(collab)
		if err != nil {
			return nil, err
		}
		return &pb.ID{Id: id}, nil
	}

	return &pb.ID{}, nil
}

func (c *CollaborationRepo) GetCollaboratorsIdByPodcastsId(podcastsId *[]string) (*[]string, error) {
	collaboratorsId := map[string]bool{}

	query := `
		select
			distinct user_id
		from 
			collaborations
		where
			podcast_id = $1	
		`

	for _, podcastId := range *podcastsId {
		collaboratorsIdOfAPodcasts := []string{}

		rows, err := c.Db.Query(query, podcastId)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			collaboratorIdOfAPodcasts := ""
			err := rows.Scan(&collaboratorIdOfAPodcasts)
			if err != nil {
				return nil, err
			}
			collaboratorsIdOfAPodcasts = append(collaboratorsIdOfAPodcasts, collaboratorIdOfAPodcasts)
		}

		for _, val := range collaboratorsIdOfAPodcasts {
			collaboratorsId[val] = true
		}
	}
	res := make([]string, len(collaboratorsId))

	for val := range collaboratorsId {
		res = append(res, val)
	}

	return &res, nil
}

func (c *CollaborationRepo) GetPodcastsIdByCollaboratorsId(colaboratorsId *[]string) (*[]string, error) {
	podcastsId := map[string]bool{}

	query := `
		select
			distinct podcast_id
		from
			collaborations
		where
			user_id = $1
		`

	for _, colaboratorId := range *colaboratorsId {
		podcastsIdOfUser := []string{}

		rows, err := c.Db.Query(query, colaboratorId)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			podcastIdOfUser := ""
			err := rows.Scan(&podcastIdOfUser)
			if err != nil {
				return nil, err
			}
			podcastsIdOfUser = append(podcastsIdOfUser, podcastIdOfUser)
		}

		for _, val := range podcastsIdOfUser {
			podcastsId[val] = true
		}
	}
	res := make([]string, len(podcastsId))

	for val := range podcastsId {
		res = append(res, val)
	}

	return &res, nil
}
