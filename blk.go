package blk

import (
	"context"
	"net/url"

	"github.com/mousemin/blk/types"
)

type Blk interface {
	// Name 名称
	Name() string
	// Author 作者
	Author() types.Author
	// DataType 当前数据类型
	DataType() types.DataType
	// IsByURI 判断当前的url是否是当前blk
	IsByURI(uri *url.URL) bool
	// InitConf 初始化配置文件
	InitConf(conf []byte) error
	// SetLicense 设置license
	SetLicense(license []byte) error
	// Search 搜索
	Search(ctx context.Context, search string, page, size int) (*types.DataSearch, error)
	// DataInfo 获取数据信息
	DataInfo(ctx context.Context, uri string) (*types.DataInfo, error)
	// DataItemInfo 数据项信息
	DataItemInfo(ctx context.Context, uri string) (*types.DataItemInfo, error)
}
