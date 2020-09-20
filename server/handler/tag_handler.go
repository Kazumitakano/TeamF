package handler

import (
	"fmt"
	"github.com/go-openapi/runtime/middleware"
	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/models"
	"github.com/shortintern2020-C-cryptograph/TeamF/server/gen/restapi/scenepicks"
	"log"
	"time"
)

func GetTag(p scenepicks.GetTagParams) middleware.Responder {
	//tagType := p.Type
	offset := p.Offset
	limit := p.Limit
	sort := p.Sort
	genre := p.Genre
	q := p.Q
	fmt.Printf("GET /tag offset: %d, limit: %d, sort: %v, genre: %s, q: %v\n", offset, limit, sort, genre, q)
	schema, err := getTag()
	if err != nil {
		log.Fatal(err)
	}

	//result := &models.Tag{
	//	ID:   1,
	//	Name: "ハウルの動く城",
	//	Type: "アニメ",
	//}
	//schema = append(schema, result)
	params := &scenepicks.GetTagOKBody{
		Message: "success",
		Schema:  schema,
	}
	return scenepicks.NewGetTagOK().WithPayload(params)
}

func PostTag(p scenepicks.PostTagParams) middleware.Responder {

	name := p.Tag.Name
	tagType := p.Tag.Type
	fmt.Printf("POST /tag name: %s, type: %s\n", name, tagType)

	// TODO: ここでfirebase認証

	// DBへ書き込み
	id, err := postTag(name, tagType)
	if err != nil {
		log.Println(err)
	}
	params := &scenepicks.PostTagOKBody{
		Message: "success",
		ID:      id,
	}
	return scenepicks.NewPostTagOK().WithPayload(params)
}

type tag struct {

	// id
	ID int64 `json:"id,omitempty" db:"id"`

	Name string `json:"id,omitempty" db:"name"`

	Type string `json:"type,omitempty" db:"type"`

	CTime time.Time `json:"ctime" db:"ctime"`

	UTime time.Time `json:"utime" db:"utime"`
}

func mapTag(t tag) models.Tag {
	res := models.Tag{
		ID:   t.ID,
		Name: t.Name,
		Type: t.Type,
	}
	return res
}
