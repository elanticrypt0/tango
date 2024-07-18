package pagination

import "strconv"

type Pagination struct {
	PageStart    int
	PageEnd      int
	PageCurrent  int
	PageNext     int
	PagePrev     int
	ItemsPerPage int
}

func NewPagination(currentPage, itemsPerPage, totalItems int) *Pagination {
	p := &Pagination{
		PageStart:    1,
		ItemsPerPage: itemsPerPage,
		PageCurrent:  currentPage,
	}
	p.calculatePages(totalItems)
	if currentPage+1 <= p.PageEnd {
		p.PageNext++
	}

	if currentPage-1 > p.PageStart {
		p.PagePrev--
	}
	return p
}

func (p *Pagination) calculatePages(totalItems int) {
	if totalItems >= p.ItemsPerPage {
		pages := totalItems / p.ItemsPerPage
		pagesMod := totalItems % p.ItemsPerPage
		if pagesMod > p.ItemsPerPage {
			pages++
		}
		p.PageEnd = pages
	} else {
		p.PageEnd = 1
	}

}

func (p *Pagination) ToString(elem string) string {
	var val string
	switch elem {
	case "end":
		val = strconv.Itoa(p.PageEnd)
	case "start":
		val = strconv.Itoa(p.PageStart)
	case "next":
		val = strconv.Itoa(p.PageNext)
	case "prev":
		val = strconv.Itoa(p.PagePrev)
	case "current":
		val = strconv.Itoa(p.PagePrev)
	}
	return val
}
