package service

import (
	pb "collaboration_service/genproto/comments"
	pbu "collaboration_service/genproto/user"
	"collaboration_service/pkg"
	"collaboration_service/storage/postgres"
	"context"
	"database/sql"
	"log"
)

type Comments struct {
	pb.UnimplementedCommentsServer
	Repo       *postgres.CommentRepo
	UserClient pbu.UserManagementClient
}

func NewComments(db *sql.DB) *Comments {
	client, err := pkg.CreateUserManagementClient()
	if err != nil {
		log.Fatal(err)
	}

	comments := postgres.NewCommentRepo(db)
	return &Comments{Repo: comments, UserClient: client}
}

func (c *Comments) CreateCommentByPodcastId(ctx context.Context, comment *pb.CreateComment) (*pb.ID, error) {
	id, err := c.Repo.CreateCommentByPodcastId(comment)
	return id, err
}

func (c *Comments) CreateCommentByEpisodeId(ctx context.Context, epComment *pb.EpisodeComment) (*pb.ID, error) {
	id, err := c.Repo.CreateEpisodeComment(epComment)

	return &pb.ID{Id: id}, err
}

func (c *Comments) GetCommentsByPodcastId(ctx context.Context, id *pb.ID) (*pb.AllComments, error) {
	commentInfo, err := c.Repo.GetCommentsByPodcastId(id)
	if err != nil {
		return nil, err
	}
	allComments := pb.AllComments{}
	for i := 0; i < len(commentInfo); i++ {
		comment := pb.Comment{}
		user, err := c.UserClient.GetUserByID(context.Background(), &pbu.ID{Id: (commentInfo)[i].UserId})
		if err != nil {
			return nil, err
		}
		comment.Username = user.Username
		comment.Content = (commentInfo)[i].Content
		comment.CreatedAt = (commentInfo)[i].CreatedAt
		comment.UpdatedAt = (commentInfo)[i].UpdatedAt
		allComments.Comments = append(allComments.Comments, &comment)
	}
	return &allComments, nil
}

func (c *Comments) CountComments(ctx context.Context, filter *pb.CountFilter) (*pb.CommentCount, error) {
	count, err := c.Repo.CountComments(filter)

	return count, err
}
