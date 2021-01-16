package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
)

// GetUser to get user detail information.
func (c *Cacher) GetUser(user string) (data *model.User, code int, err error) {
	// Get from cache.
	key := getKey(keyUser, user)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUser(user)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserStats to get user stats detail information.
func (c *Cacher) GetUserStats(user string) (data *model.UserStats, code int, err error) {
	// Get from cache.
	key := getKey(keyUserStats, user)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserStats(user)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserFavorite to get user favorite list.
func (c *Cacher) GetUserFavorite(user string) (data *model.UserFavorite, code int, err error) {
	// Get from cache.
	key := getKey(keyUserFavorite, user)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserFavorite(user)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserFriend to get user friend list.
func (c *Cacher) GetUserFriend(user string, page int) (data []model.UserFriend, code int, err error) {
	// Get from cache.
	key := getKey(keyUserFriend, user, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserFriend(user, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserHistory to get user history list.
func (c *Cacher) GetUserHistory(user string, t string) (data []model.UserHistory, code int, err error) {
	// Get from cache.
	key := getKey(keyUserHistory, user, t)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserHistory(user, t)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserReview to get user review list.
func (c *Cacher) GetUserReview(user string, page int) (data []model.Review, code int, err error) {
	// Get from cache.
	key := getKey(keyUserReview, user, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserReview(user, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserRecommendation to get user recommendation list.
func (c *Cacher) GetUserRecommendation(user string, page int) (data []model.Recommendation, code int, err error) {
	// Get from cache.
	key := getKey(keyUserRecommendation, user, page)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserRecommendation(user, page)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserClub to get user club list.
func (c *Cacher) GetUserClub(user string) (data []model.Item, code int, err error) {
	// Get from cache.
	key := getKey(keyUserClub, user)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserClub(user)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserAnime to get user anime list.
func (c *Cacher) GetUserAnime(query model.UserListQuery) (data []model.UserAnime, code int, err error) {
	// Get from cache.
	key := getKey(keyUserAnime, query.Username, query.Page, query.Status, query.Order, query.Tag)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserAnime(query)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetUserManga to get user manga list.
func (c *Cacher) GetUserManga(query model.UserListQuery) (data []model.UserManga, code int, err error) {
	// Get from cache.
	key := getKey(keyUserManga, query.Username, query.Page, query.Status, query.Order, query.Tag)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetUserManga(query)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
