package parser

import (
	"net/http"
	"testing"
	"time"

	"github.com/rl404/go-malscraper/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProducers(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetProducers()
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyCount := true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Count != 0 {
			emptyCount = false
		}
	}
	assert.False(t, emptyCount)
	time.Sleep(sleepDur)
}

func TestGetProducer(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetProducer(0, 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetProducer(1, 1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyY, emptyM, emptyD, emptyLic, emptyEp := true, true, true, true, true
		for _, p := range d {
			assert.NotZero(t, p.ID)
			assert.NotEmpty(t, p.Image)
			assert.NotEmpty(t, p.Title)
			assert.NotEmpty(t, p.Source)
			assert.NotEmpty(t, p.Synopsis)
			assert.NotEmpty(t, p.Type)
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
			assert.NotZero(t, p.Score)
			assert.NotZero(t, len(p.Genres))
			for _, g := range p.Genres {
				assert.NotZero(t, g.ID)
				assert.NotEmpty(t, g.Name)
			}
			assert.NotZero(t, len(p.Producers))
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
		assert.False(t, emptyY)
		assert.False(t, emptyM)
		assert.False(t, emptyD)
		assert.False(t, emptyEp)
		assert.False(t, emptyLic)
	})
	time.Sleep(sleepDur)
}

func TestGetMagazines(t *testing.T) {
	parser := New(true, true, log)
	d, code, err := parser.GetMagazines()
	require.NotNil(t, d)
	require.Equal(t, code, http.StatusOK)
	require.NoError(t, err)

	assert.NotZero(t, len(d))
	emptyCount := true
	for _, p := range d {
		assert.NotZero(t, p.ID)
		assert.NotEmpty(t, p.Name)
		if p.Count != 0 {
			emptyCount = false
		}
	}
	assert.False(t, emptyCount)
	time.Sleep(sleepDur)
}

func TestGetMagazine(t *testing.T) {
	parser := New(true, true, log)
	t.Run("invalid-id", func(t *testing.T) {
		d, code, err := parser.GetMagazine(0, 1)
		assert.Nil(t, d)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		d, code, err := parser.GetMagazine(83, 1)
		require.NotNil(t, d)
		require.Equal(t, code, http.StatusOK)
		require.NoError(t, err)

		assert.NotZero(t, len(d))
		emptyY, emptyM, emptyD, emptySer, emptyVol := true, true, true, true, true
		for _, p := range d {
			assert.NotZero(t, p.ID)
			assert.NotEmpty(t, p.Image)
			assert.NotEmpty(t, p.Title)
			assert.NotEmpty(t, p.Type)
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
			assert.NotZero(t, p.Score)
			assert.NotEmpty(t, p.Synopsis)
			assert.NotZero(t, len(p.Genres))
			for _, g := range p.Genres {
				assert.NotZero(t, g.ID)
				assert.NotEmpty(t, g.Name)
			}
			assert.NotZero(t, len(p.Authors))
			for _, pp := range p.Authors {
				assert.NotZero(t, pp.ID)
				assert.NotEmpty(t, pp.Name)
			}
			if p.Volume != 0 {
				emptyVol = false
			}
			for _, l := range p.Serializations {
				assert.NotEmpty(t, l)
				emptySer = false
			}
		}
		assert.False(t, emptyY)
		assert.False(t, emptyM)
		assert.False(t, emptyD)
		assert.False(t, emptyVol)
		assert.False(t, emptySer)
	})
	time.Sleep(sleepDur)
}
