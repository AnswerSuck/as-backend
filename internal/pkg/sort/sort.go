package sort

import (
	"strings"

	"github.com/Masterminds/squirrel"
)

const (
	OrderASC  = "ASC"
	OrderDESC = "DESC"
)

type Sort struct {
	col   string
	order string
}

func NewSort(col, order string) Sort {
	order = strings.ToUpper(order)
	if order != OrderASC && order != OrderDESC {
		order = OrderASC
	}

	return Sort{
		col:   col,
		order: order,
	}
}

// Attach attaches sort to builder.
func (s Sort) Attach(b squirrel.SelectBuilder) squirrel.SelectBuilder {
	return b.OrderBy(s.col + " " + s.order)
}