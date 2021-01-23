package malscraper

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSeason(t *testing.T) {
	d, code, err := mal.GetSeason("fall", 2017)
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyImg, emptySyn, emptyScore, emptyLic, emptyEp, emptyGenre, emptyProd := true, true, true, true, true, true, true
	emptyY, emptyM, emptyD := true, true, true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		if p.Image != "" {
			emptyImg = false
		}
		assert.NotEmpty(t, p.Title)
		assert.NotEmpty(t, p.Source)
		if p.Synopsis != "" {
			emptySyn = false
		}
		assert.NotEmpty(t, p.Type)
		assert.NotNil(t, p.StartDate)
		if p.StartDate.Year != 0 {
			emptyY = false
		}
		if p.StartDate.Month != 0 {
			emptyM = false
		}
		if p.StartDate.Day != 0 {
			emptyD = false
		}
		assert.NotZero(t, p.Member)
		if p.Score != 0.0 {
			emptyScore = false
		}
		if len(p.Genres) > 0 {
			emptyGenre = false
		}
		for _, g := range p.Genres {
			assert.NotZero(t, g.ID)
			assert.NotEmpty(t, g.Name)
		}
		if len(p.Producers) > 0 {
			emptyProd = false
		}
		for _, pp := range p.Producers {
			assert.NotZero(t, pp.ID)
			assert.NotEmpty(t, pp.Name)
		}
		if p.Episode != 0 {
			emptyEp = false
		}
		for _, l := range p.Licensors {
			assert.NotEmpty(t, l)
			emptyLic = false
		}
	}
	assert.False(t, emptyImg)
	assert.False(t, emptySyn)
	assert.False(t, emptyScore)
	assert.False(t, emptyY)
	assert.False(t, emptyM)
	assert.False(t, emptyD)
	assert.False(t, emptyEp)
	assert.False(t, emptyLic)
	assert.False(t, emptyGenre)
	assert.False(t, emptyProd)
	time.Sleep(sleepDur)
}
