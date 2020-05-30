package shorten

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"onursume.com/shortener/consts"
	"onursume.com/shortener/db"
)

func ShortenUrl(c echo.Context) error {
	url := c.Param("url")
	unescapedUrl, check := urlValidator{path: url}.check()

	if check != nil {
		return c.String(http.StatusBadRequest, check.Error())
	}

	id, errId := GenerateId()
	if errId != nil {
		panic(errId)
	}

	document := Url{UniqueId: id, UrlPath:unescapedUrl}
	couchbase := &db.Couchbase{BucketName: consts.UrlBucket}
	_, err := couchbase.GetBucket().DefaultCollection().Insert(id, &document, nil)
	if err != nil {
		panic(err)
	}

	shortedUrl := "Shortened Link: \nlocalhost:" + consts.ServerPort + "/" + id
	return c.String(http.StatusOK, shortedUrl)
}

func RedirectToUrl(c echo.Context) error {
	id := c.Param("id")
	couchbase := &db.Couchbase{BucketName: consts.UrlBucket}
	result, err := couchbase.GetBucket().DefaultCollection().Get(id, nil)

	if err != nil{
		return err
	}

	var entity Url
	err = result.Content(&entity)
	if err != nil {
		panic(err)
	}

	return c.Redirect(http.StatusMovedPermanently, entity.UrlPath)
}