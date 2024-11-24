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

func (s *SongService) CreateSong(a, song music.Song) error {
	return s.repo.CreateSong(song)
}
