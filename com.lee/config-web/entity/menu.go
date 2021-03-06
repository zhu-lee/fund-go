package entity

import "fmt"

type MenuNode struct {
	ID           string
	Level        int
	Parent       *MenuNode
	VisibleItems []*MenuNode
	Name         string      `json:"name" xml:"name,attr"`
	Icon         string      `json:"icon" xml:"icon,attr"`
	Url          string      `json:"url" xml:"url,attr"`
	Hidden       bool        `json:"hidden" xml:"hidden,attr"`
	Items        []*MenuNode `json:"items" xml:"item"`
}

func (n *MenuNode) SetID(k int) {
	n.ID = fmt.Sprintf("%v.%v", n.Level, n)
}

type MenuInfo struct {
	Current    *MenuNode
	LeftMenus  []*MenuNode
	Breadcrumb []*MenuNode
}
