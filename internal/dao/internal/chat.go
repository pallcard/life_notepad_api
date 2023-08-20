// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatDao is the data access object for table t_chat.
type ChatDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns ChatColumns // columns contains all the column names of Table for convenient usage.
}

// ChatColumns defines and stores column names for table t_chat.
type ChatColumns struct {
	Id        string //
	UserId    string // 用户ID，逗号分割
	SenderId  string // 发送者ID,最新的一条
	Content   string // 内容，最新的一条
	IsLiked   string // 是否链接 1是 2不是
	Unread    string // 未读 1未读 2已读
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// chatColumns holds the columns for table t_chat.
var chatColumns = ChatColumns{
	Id:        "id",
	UserId:    "user_id",
	SenderId:  "sender_id",
	Content:   "content",
	IsLiked:   "is_liked",
	Unread:    "unread",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewChatDao creates and returns a new DAO object for table data access.
func NewChatDao() *ChatDao {
	return &ChatDao{
		group:   "default",
		table:   "t_chat",
		columns: chatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ChatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ChatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ChatDao) Columns() ChatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ChatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ChatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ChatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
