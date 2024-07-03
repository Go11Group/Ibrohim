package service

import (
	"context"
	"database/sql"
	pb "podcast_service/genproto/episodes"
	"podcast_service/storage/postgres"
)

type EpisodeService struct {
	pb.UnimplementedEpisodesServiceServer
	Episode *postgres.EpisodeRepo
	Podcase *postgres.PodcastRepo
}

func NewEpisodeService(db *sql.DB) *EpisodeService {
	episodeRepo := postgres.NewEpisodeRepo(db)
	podcastRepo := postgres.NewPodcastRepo(db)
	return &EpisodeService{Episode: episodeRepo, Podcase: podcastRepo}
}

func (e *EpisodeService) GetEpisodesByPodcastId(ctx context.Context, req *pb.ID) (*pb.Episodes, error) {
	resp, err := e.Episode.GetEpisodesByPodcastId(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EpisodeService) UpdateEpisode(ctx context.Context, req *pb.IDs) (*pb.Void, error) {
	resp, err := e.Episode.UpdateEpisode(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EpisodeService) CreatePodcastEpisode(ctx context.Context, podcast *pb.EpisodeCreate) (*pb.ID, error) {
	id, err := e.Episode.CreatePodcastEpisode(podcast)

	return &pb.ID{Id: id}, err
}

func (e *EpisodeService) DeleteEpisode(ctx context.Context, ids *pb.IDsForDelete) (*pb.Void, error) {
	err := e.Episode.DeletePodcastEpisode(ids)

	return &pb.Void{}, err
}

func (e *EpisodeService) PublishPodcast(ctx context.Context, id *pb.ID) (*pb.Success, error) {
	success, err := e.Podcase.PublishPodcast(id)

	return success, err
}
