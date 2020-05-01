package util

import (
	"math"
	"net/url"

	"github.com/gorilla/schema"
)

type Pager struct {
	PageSize   int `json:"pageSize", schema:"pageSize"`
	PageNumber int `json:"pageNumber", schema:"pageNumber`
	TotalPage  int `json:"totalPage", schema:"-"`
}

func (p *Pager) GetOffset() int {
	return p.PageSize * (p.PageNumber - 1)
}

func (p *Pager) UpdateTotalPage(total int) {
	if total > 0 {
		p.TotalPage = int(math.Ceil(float64(total) / float64(p.PageSize)))
	} else {
		p.TotalPage = 0
	}
}

var decoder = schema.NewDecoder()

func InitPager(requestParams url.Values) *Pager {
	pager := &Pager{}
	decoder.Decode(pager, requestParams)

	if pager.PageSize == 0 {
		pager.PageSize = 10
	}

	if pager.PageNumber == 0 {
		pager.PageNumber = 1
	}

	return pager
}
