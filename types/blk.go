package types

type (
	// DataSearch 搜索展示的详细信息
	DataSearch struct {
		Data []DataSearchItem `json:"data"` // 信息
		Page int              `json:"page"` // 页码
		Size int              `json:"size"` // 数量
	}

	DataSearchItem struct {
		URI        string   `json:"uri"`         // 地址
		Title      string   `json:"title"`       // 标题
		Author     string   `json:"author"`      // 作者
		Category   []string `json:"category"`    // 分类
		Status     string   `json:"status"`      // 状态
		UpdateTime *Time    `json:"update_time"` // 更新时间
	}
)

type (
	// DataInfo 数据详细信息
	DataInfo struct {
		Title      string         `json:"title"`       // 标题
		Author     string         `json:"author"`      // 作者
		Pic        string         `json:"pic"`         // 封面
		Desc       string         `json:"desc"`        // 简介
		Category   []string       `json:"category"`    // 分类
		Status     string         `json:"status"`      // 状态
		UpdateTime *Time          `json:"update_time"` // 更新时间
		List       []DataListItem `json:"list"`        // 章节列表
	}

	DataListItem struct {
		Title string `json:"title"` // 标题
		URI   string `json:"uri"`   // 地址
	}
)

type (
	// DataItemInfo 获取数据项的信息
	DataItemInfo struct {
		Title      string     `json:"title"`      // 标题
		Content    []string   `json:"content"`    // 内容
		Pagination Pagination `json:"pagination"` // 页码
	}

	Pagination struct {
		Prev string `json:"prev"` // 上一章
		Next string `json:"next"` // 下一章
		Info string `json:"info"` // 详细信息
	}
)
