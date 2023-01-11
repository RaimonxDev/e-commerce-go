package postgres

import (
	"github.com/RaimonxDev/e-commerce-go.git/model"
)

func DefaultPagination(p *model.Pagination) {
	// default limit
	if p.Limit == 0 {
		p.Limit = 10
	}
	// Default page
	if p.Page == 0 {
		p.Page = 1
	}
	// Max limit
	if p.Limit > 100 {
		p.Limit = 100
	}
}
