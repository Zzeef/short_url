package link

import (
	"context"
	"errors"
)

type LinkService struct {
	repo *LinkRepo
}

func NewService(repo *LinkRepo) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) GetRecord(ctx context.Context, code string) (*Link, error) {
	if code == "" {
		return nil, errors.New("code is empty")
	}

	record, err := s.repo.GetRecordByColumn(ctx, "shortCode", code)
	if err != nil {
		return nil, err
	}

	if record != nil {
		return record, nil
	}

	return nil, nil
}

func generateShortCode() string {
	return ""
}

func (s *LinkService) Shorten(ctx context.Context, url string) (string, error) {
	if url == "" {
		return "", errors.New("url is empty")
	}

	record, err := s.repo.GetRecordByColumn(ctx, "URL", url)
	if err != nil {
		return "", err
	}
	if record != nil {
		return record.ShortCode, nil
	}

	shortCode := generateShortCode()

	err = s.repo.Insert(ctx, &Link{URL: url, ShortCode: shortCode})
	if err != nil {
		return "", errors.New("error while insert")
	}
	return shortCode, nil
}
