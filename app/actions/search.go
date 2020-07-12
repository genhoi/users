package actions

import (
	"github.com/aofei/air"
	"github.com/genhoi/users/module/user"
)

type Search struct {
	userRepo *user.Repository
}

func NewSearch(userRepo *user.Repository) *Search {
	return &Search{userRepo: userRepo}
}

func (a *Search) Head(req *air.Request, res *air.Response) error {
	res.Header.Add("Access-Control-Allow-Origin", "*")
	return nil
}
func (a *Search) Get(req *air.Request, res *air.Response) error {
	res.Header.Add("Access-Control-Allow-Origin", "*")

	query := ""
	if queryParam := req.Param("query"); queryParam != nil {
		query = queryParam.Value().String()
	}

	users, err := a.userRepo.FindAllByQuery(query, 5)
	if err != nil {
		return err
	}

	return res.WriteJSON(users)
}
