package service

import (
	"context"
	"database/sql"
	pb "podcast_service/genproto/podcasts"
	"podcast_service/storage/postgres"
)

type PodcastService struct {
	pb.UnimplementedPodcastsServer
	Podcast *postgres.PodcastRepo
}

func NewPodcastService(db *sql.DB) *PodcastService {
	podcast := postgres.NewPodcastRepo(db)
	return &PodcastService{Podcast: podcast}
}

func (p *PodcastService) CreatePodcast(ctx context.Context, req *pb.PodcastCreate) (*pb.ID, error) {
	resp, err := p.Podcast.CreatePodcast(req)
	if err != nil {
		return nil, err
	}

	return &pb.ID{Id: *resp}, nil
}

func (p *PodcastService) GetPodcastById(ctx context.Context, req *pb.ID) (*pb.Podcast, error) {
	resp, err := p.Podcast.GetPodcastById(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) UpdatePodcast(ctx context.Context, podcast *pb.PodcastUpdate) (*pb.Void, error) {
	err := p.Podcast.UpdatePodcast(podcast)

	return &pb.Void{}, err
}

func (p *PodcastService) DeletePodcast(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	resp, err := p.Podcast.DeletePodcast(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PodcastService) GetUserPodcasts(ctx context.Context, req *pb.ID) (*pb.UserPodcasts, error) {
	resp, err := p.Podcast.GetUserPodcasts(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}