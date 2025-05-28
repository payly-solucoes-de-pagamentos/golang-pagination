package pagination

import (
	"errors"

	paginate "github.com/gobeam/mongo-go-pagination"
)

type SortOrientation int

const (
	ASC  SortOrientation = 1
	DESC SortOrientation = -1
)

func (sort SortOrientation) Int() int {
	return int(sort)
}

func ToSortOrientation(sort string) SortOrientation {
	if sort == "ASC" {
		return ASC
	} else if sort == "DESC" {
		return DESC
	}
	err := errors.New("could not parse sort")
	panic(err)
}

func (sort SortOrientation) String() string {
	if sort == ASC {
		return "ASC"
	} else {
		return "DESC"
	}
}

type FallBackPaginationFunc func(config *PaginationQuery)

const defaultPage int64 = 1
const defaultSize int64 = 50
const defaultSort SortOrientation = ASC

type PaginationQuery struct {
	Size   int64           `json:"size"`
	Page   int64           `json:"page"`
	Sort   SortOrientation `json:"sort"`
	SortBy string          `json:"sortBy"`
}

type PaginationData struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}

func ToPaginationData(paginationData paginate.PaginationData) *PaginationData {
	return &PaginationData{
		Total:     paginationData.Total,
		Page:      paginationData.Page,
		PerPage:   paginationData.PerPage,
		Prev:      paginationData.Prev,
		Next:      paginationData.Next,
		TotalPage: paginationData.TotalPage,
	}
}

type PaginationResult struct {
	BaseURL    string `json:"baseURL"`
	PageURL    string `json:"pageURL"`
	PrevURL    string `json:"prevURL"`
	NextURL    string `json:"nextURL"`
	Total      int64  `json:"total"`
	Page       int64  `json:"page"`
	PerPage    int64  `json:"perPage"`
	Prev       int64  `json:"prev"`
	Next       int64  `json:"next"`
	TotalPages int64  `json:"totalPages"`
}
