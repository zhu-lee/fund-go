package model

import (
	"com.lee/config-web/biz"
	"com.lee/config-web/entity"
	"com.lee/fund/util/convert"
	"com.lee/fund/web"
	"net/url"
)

const (
	DefaultPageSize = 25
)

type BaseModel struct {
	*web.SM
	*web.PM
	*entity.MenuInfo
}

func NewBaseModel(ctx *web.Context) *BaseModel {
	bm := BaseModel{
		SM: web.NewSM(ctx),
		PM: new(web.PM),
	}

	//bm.AddCss(ctx, "")
	//bm.AddJs(ctx, "")
	bm.MenuInfo = biz.Menu.GetMenuInfo(ctx.Request.URL.Path)

	return &bm
}

type PageModel struct {
	PageSize   int   //每页数量
	PageIndex  int   //当前页码
	PageCount  int   //总页数
	TotalCount int   //总条数
	Pages      []int //显示的页码
	Query      url.Values
}

func NewPageModel(ctx *web.Context, pageIndex int, pageSize, totalCount int) *PageModel {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageIndex <= 0 {
		pageIndex = 1
	}

	m := &PageModel{
		PageIndex:  pageIndex,
		PageSize:   pageSize,
		TotalCount: totalCount,
	}

	m.PageCount = int(totalCount / pageSize)
	if totalCount%pageSize > 0 {
		m.PageCount++
	}

	// 最多显示7个页码
	var start, end int
	start = pageIndex - 3
	end = pageIndex + 3
	if m.PageCount <= 7 {
		start, end = 1, m.PageCount
	} else if pageIndex < 4 {
		start, end = 1, 7
	} else if pageIndex > m.PageCount-3 {
		end = m.PageCount
		start = end - 6
	}
	for i := start; i <= end; i++ {
		m.Pages = append(m.Pages, i)
	}

	u, _ := url.Parse(ctx.Request.RequestURI)
	m.Query = u.Query()
	return m
}

func (m *PageModel) Url(pi int) string {
	m.Query.Set("pi", convert.IntToString(pi))
	return "?" + m.Query.Encode()
}

func (m *PageModel) Page(inc int) int {
	return m.PageIndex + inc
}
