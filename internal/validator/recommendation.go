package validator

import (
	"net/http"

	"github.com/rl404/go-malscraper/errors"
	"github.com/rl404/go-malscraper/model"
)

// GetRecommendation to get recommendation detail information.
func (v *Validator) GetRecommendation(t string, id1, id2 int) (*model.Recommendation, int, error) {
	if t != AnimeType && t != MangaType {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if id1 <= 0 || id2 <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidID
	}
	if t == AnimeType && (v.isEmptyID(getKey(keyEmptyAnime, id1)) || v.isEmptyID(getKey(keyEmptyAnime, id2))) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}
	if t == MangaType && (v.isEmptyID(getKey(keyEmptyManga, id1)) || v.isEmptyID(getKey(keyEmptyManga, id2))) {
		return nil, http.StatusNotFound, errors.ErrNot200
	}
	return v.api.GetRecommendation(t, id1, id2)
}

// GetRecommendations to get anime/manga recommendation list.
func (v *Validator) GetRecommendations(t string, page int) ([]model.Recommendation, int, error) {
	if t != AnimeType && t != MangaType {
		return nil, http.StatusBadRequest, errors.ErrInvalidType
	}
	if page <= 0 {
		return nil, http.StatusBadRequest, errors.ErrInvalidPage
	}
	return v.api.GetRecommendations(t, page)
}
