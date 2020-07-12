package actions

import (
	"github.com/aofei/air"
)

type Ui struct {
}

func (a *Ui) Get(req *air.Request, res *air.Response) error {
	p := "./ui/dist" + req.HTTPRequest().URL.Path
	if "./ui/dist/" == p {
		p += "index.html"
	}
	return res.WriteFile(p)
}
