package service

import (
	pb "collaboration_service/genproto/collaborations"
	pbu "collaboration_service/genproto/user"
	"collaboration_service/pkg"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
	"log"
)

type Collaborations struct {
	pb.UnimplementedCollaborationsServer
	Repo       *postgres.CollaborationRepo
	UserClient pbu.UserManagementClient
}

func NewCollaborations(db *sql.DB) *Collaborations {
	client, err := pkg.CreateUserManagementClient()
	if err != nil {
		log.Fatal(err)
	}

	collaborations := postgres.NewCollaborationRepo(db)
	return &Collaborations{Repo: collaborations, UserClient: client}
}

func (c *Collaborations) CreateInvitation(ctx context.Context, invitation *pb.CreateInvite) (*pb.ID, error) {
	id, err := c.Repo.CreateInvitation(invitation)
	if err != nil {
		return nil, err
	}
	return &pb.ID{Id: id}, err
}

func (c *Collaborations) RespondInvitation(ctx context.Context, req *pb.CreateCollaboration) (*pb.ID, error) {
	resp, err := c.Repo.RespondInvitation(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Collaborations) GetCollaboratorsByPodcastId(ctx context.Context, id *pb.ID) (*pb.Collaborators, error) {
	res := pb.Collaborators{}

	collaboratorsId, err := c.Repo.GetCollaboratorsByPodcastId(id.Id)
	if err != nil {
		return nil, err
	}

	for _, col := range collaboratorsId {
		userInfo, err := c.UserClient.GetUserByID(context.Background(), &pbu.ID{Id: col.UserId})
		if err != nil {
			return nil, err
		}
		colaborator := pb.Collaborator{
			Username: userInfo.Username,
			Email:    userInfo.Email,
			Role:     col.Role,
			JoinedAt: col.JoinedAt,
		}
		res.Collaborators = append(res.Collaborators, &colaborator)
	}
	return &res, nil
}

func (c *Collaborations) UpdateCollaboratorByPodcastId(ctx context.Context, clb *pb.UpdateCollaborator) (*pb.Void, error) {
	err := c.Repo.UpdateCollaboratorByPodcastId(clb)

	return &pb.Void{}, err
}

func (c *Collaborations) DeleteCollaboratorByPodcastId(ctx context.Context, req *pb.Ids) (*pb.Void, error) {
	resp, err := c.Repo.DeleteCollaboratorByPodcastId(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Collaborations) GetAllPodcastsUsersWorkedOn(ctx context.Context, podcastsId *pb.PodcastsId) (*pb.PodcastsId, error) {
	colaborators, err := c.Repo.GetCollaboratorsIdByPodcastsId(&podcastsId.PodcastsId)
	if err != nil {
		return nil, err
	}

	podcastsIdToReturn, err := c.Repo.GetPodcastsIdByCollaboratorsId(colaborators)
	return &pb.PodcastsId{PodcastsId: *podcastsIdToReturn}, err
}
