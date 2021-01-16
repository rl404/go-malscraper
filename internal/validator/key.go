package validator

import (
	"fmt"
	"strings"
)

const (
	keyEmptyAnime   = "mal:empty:anime"
	keyEmptyManga   = "mal:empty:manga"
	keyEmptyChar    = "mal:empty:character"
	keyEmptyPeople  = "mal:empty:people"
	keyEmptyArticle = "mal:empty:article"
	keyEmptyClub    = "mal:empty:club"
	keyEmptyNews    = "mal:empty:news"
	keyEmptyReview  = "mal:empty:review"
	keyEmptyUser    = "mal:empty:user"
	keyArticleTag   = "mal:article-tag"
	keyGenres       = "mal:genres"
	keyNewsTag      = "mal:news-tag"
	keyProducers    = "mal:producers"
	keyMagazines    = "mal:magazines"
)

func getKey(key string, params ...interface{}) string {
	strParams := []string{key}
	for _, p := range params {
		strParams = append(strParams, fmt.Sprintf("%v", p))
	}
	return strings.Join(strParams, ":")
}
