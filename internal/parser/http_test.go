package parser

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/errors"
	iMock "github.com/rl404/go-malscraper/internal/mocks"
	rMock "github.com/rl404/go-malscraper/internal/parser/mocks"
	"github.com/rl404/go-malscraper/model"
	"github.com/stretchr/testify/assert"
)

type responseMock struct{}

func (r *responseMock) Read([]byte) (int, error) { return 0, nil }
func (r *responseMock) Close() error             { return nil }

func TestGetBody(t *testing.T) {
	rMock := new(rMock.Requester)
	lMock := new(iMock.Logger)
	t.Run("error-prepare-request", func(t *testing.T) {
		httpRequest = func(string, string, io.Reader) (*http.Request, error) {
			return nil, errDummy
		}
		lMock.On("Error", "failed preparing request: %s", errDummy.Error()).Once()
		parser := &Parser{logger: lMock}

		data, code, err := parser.getBody("")
		assert.Nil(t, data)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrPrepareRequest.Error())
	})

	t.Run("error-http-request", func(t *testing.T) {
		httpRequest = func(string, string, io.Reader) (*http.Request, error) {
			return &http.Request{}, nil
		}
		rMock.On("Do", &http.Request{}).Return(nil, errDummy).Once()
		parser := &Parser{logger: lMock, http: rMock}

		data, code, err := parser.getBody("")
		assert.Nil(t, data)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrHTTPRequest.Error())
	})

	t.Run("error-not-200", func(t *testing.T) {
		httpRequest = func(string, string, io.Reader) (*http.Request, error) {
			return &http.Request{}, nil
		}
		timeSince = func(time.Time) time.Duration {
			return time.Second
		}
		rMock.On("Do", &http.Request{}).Return(&http.Response{
			StatusCode: http.StatusNotFound,
			Body:       &responseMock{},
		}, nil).Once()
		lMock.On("Debug", "%s %v (%s)", "", http.StatusNotFound, time.Second.Truncate(time.Microsecond)).Once()
		parser := &Parser{logger: lMock, http: rMock}

		data, code, err := parser.getBody("")
		assert.Nil(t, data)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("ok", func(t *testing.T) {
		httpRequest = func(string, string, io.Reader) (*http.Request, error) {
			return &http.Request{}, nil
		}
		timeSince = func(time.Time) time.Duration {
			return time.Second
		}
		rMock.On("Do", &http.Request{}).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       &responseMock{},
		}, nil).Once()
		lMock.On("Debug", "%s %v (%s)", "", http.StatusOK, time.Second.Truncate(time.Microsecond)).Once()
		parser := &Parser{logger: lMock, http: rMock}

		data, code, err := parser.getBody("")
		assert.NotNil(t, data)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestGetDoc(t *testing.T) {
	rMock := new(rMock.Requester)
	lMock := new(iMock.Logger)

	httpRequest = func(string, string, io.Reader) (*http.Request, error) {
		return &http.Request{}, nil
	}
	timeSince = func(time.Time) time.Duration {
		return time.Second
	}

	t.Run("error-not-200", func(t *testing.T) {
		rMock.On("Do", &http.Request{}).Return(&http.Response{
			StatusCode: http.StatusNotFound,
			Body:       &responseMock{},
		}, nil).Once()
		lMock.On("Debug", "%s %v (%s)", "", http.StatusNotFound, time.Second.Truncate(time.Microsecond)).Once()
		lMock.On("Trace", "parsing %s", "").Once()
		parser := &Parser{logger: lMock, http: rMock}

		data, code, err := parser.getDoc("", "")
		assert.Nil(t, data)
		assert.Equal(t, code, http.StatusNotFound)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrNot200.Error())
	})

	t.Run("error-parse-html", func(t *testing.T) {
		rMock.On("Do", &http.Request{}).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       &responseMock{},
		}, nil).Once()
		parseHTML = func(io.Reader) (*goquery.Document, error) {
			return nil, errDummy
		}
		lMock.On("Debug", "%s %v (%s)", "", http.StatusOK, time.Second.Truncate(time.Microsecond)).Once()
		lMock.On("Trace", "parsing %s", "").Once()
		lMock.On("Error", "failed parsing body: %s", errDummy.Error()).Once()
		parser := &Parser{logger: lMock, http: rMock}

		data, code, err := parser.getDoc("", "")
		assert.Nil(t, data)
		assert.Equal(t, code, http.StatusInternalServerError)
		assert.Error(t, err)
		assert.EqualError(t, err, errors.ErrParseBody.Error())
	})

	t.Run("ok", func(t *testing.T) {
		rMock.On("Do", &http.Request{}).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       &responseMock{},
		}, nil).Once()
		parseHTML = func(io.Reader) (*goquery.Document, error) {
			return goquery.NewDocumentFromReader(strings.NewReader("<html></html>"))
		}
		lMock.On("Debug", "%s %v (%s)", "", http.StatusOK, time.Second.Truncate(time.Microsecond)).Once()
		lMock.On("Trace", "parsing %s", "").Once()
		parser := &Parser{logger: lMock, http: rMock}

		data, code, err := parser.getDoc("", "")
		assert.NotNil(t, data)
		assert.Equal(t, code, http.StatusOK)
		assert.NoError(t, err)
	})
}

func TestQqueryToMap(t *testing.T) {
	parser := &Parser{}
	y, m, d := time.Now().Date()
	query := model.Query{
		Title:        "naruto",
		Page:         2,
		Type:         1,
		Score:        7,
		Status:       3,
		ProducerID:   4,
		MagazineID:   5,
		Rating:       6,
		ExcludeGenre: true,
		StartDate:    time.Now(),
		EndDate:      time.Now(),
		GenreIDs:     []int{1, 2, 3},
		FirstLetter:  "a",
	}
	result := map[string]interface{}{
		"c[0]":     "a",
		"c[1]":     "b",
		"c[2]":     "c",
		"c[3]":     "d",
		"c[4]":     "e",
		"c[5]":     "f",
		"c[6]":     "g",
		"q":        "naruto",
		"show":     50,
		"type":     1,
		"score":    7,
		"status":   3,
		"p":        4,
		"mid":      5,
		"r":        6,
		"gx":       1,
		"sd":       d,
		"sm":       int(m),
		"sy":       y,
		"ed":       d,
		"em":       int(m),
		"ey":       y,
		"genre[0]": 1,
		"genre[1]": 2,
		"genre[2]": 3,
		"letter":   "A",
	}
	assert.Equal(t, parser.queryToMap(query), result)
}
