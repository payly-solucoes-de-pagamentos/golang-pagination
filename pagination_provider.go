package pagination

import (
	"github.com/gin-gonic/gin"
)

type IPaginationProvider interface {
	GetPaginationQuery(ctx *gin.Context, fallback ...FallBackPaginationFunc) *PaginationQuery
	GetPaginationResult(ctx *gin.Context) *PaginationResult
	SetPaginationData(paginationData *PaginationData)
}
