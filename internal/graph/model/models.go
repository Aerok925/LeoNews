package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
	"time"
)

type News struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Img         *string   `json:"img"`
	CreateAt    time.Time `json:"create_at"`
}

func MarshalID(id int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

func UnmarshalID(v interface{}) (int, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be strings")
	}
	i, e := strconv.Atoi(id)
	return int(i), e
}
