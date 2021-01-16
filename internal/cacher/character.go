package cacher

import (
	"net/http"

	"github.com/rl404/go-malscraper/model"
)

// GetCharacter to get character detail information.
func (c *Cacher) GetCharacter(id int) (data *model.Character, code int, err error) {
	// Get from cache.
	key := getKey(keyCharacter, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetCharacter(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetCharacterArticle to get character featured article list.
func (c *Cacher) GetCharacterArticle(id int) (data []model.ArticleItem, code int, err error) {
	// Get from cache.
	key := getKey(keyCharacterArticle, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetCharacterArticle(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetCharacterOgraphy to get character animeography/mangaography list.
func (c *Cacher) GetCharacterOgraphy(t string, id int) (data []model.Role, code int, err error) {
	// Get from cache.
	key := getKey(keyCharacterOgraphy, t, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetCharacterOgraphy(t, id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetCharacterPicture to get character picture list.
func (c *Cacher) GetCharacterPicture(id int) (data []string, code int, err error) {
	// Get from cache.
	key := getKey(keyCharacterPicture, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetCharacterPicture(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}

// GetCharacterClub to get character club list.
func (c *Cacher) GetCharacterClub(id int) (data []model.ClubItem, code int, err error) {
	// Get from cache.
	key := getKey(keyCharacterClub, id)
	if c.cacher.Get(key, &data) == nil {
		return data, http.StatusOK, nil
	}

	// Parse.
	data, code, err = c.api.GetCharacterClub(id)
	if err != nil {
		return nil, code, err
	}

	// Save to cache. Won't return error.
	_ = c.cacher.Set(key, data)
	return data, http.StatusOK, nil
}
