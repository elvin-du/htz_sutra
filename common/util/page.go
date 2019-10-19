package util

type Page struct {
	PageIndex int32
	PageSize  int32
}

type PageInfo struct {
	PageIndex  int32       `json:"page_index"`
	PageSize   int32       `json:"page_size"`
	TotalPages int64       `json:"total_pages"`
	TotalItems int64       `json:"total_items"`
	Items      interface{} `json:"items"`
}

func NewPageInfo(page Page, total int64, items interface{}) *PageInfo {
	totalPages := total / int64(page.PageSize)
	return &PageInfo{page.PageIndex, page.PageSize, totalPages, total, items}
}
