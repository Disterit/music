package service

import (
	"music"
	"music/pkg/repository"
)

type SongService struct {
	repo repository.Song
}

func NewSongService(repo repository.Song) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(song music.Song, artistId int) (int, error) {
	return s.repo.CreateSong(song, artistId)
}
