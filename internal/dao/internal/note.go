// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NoteDao is the data access object for table t_note.
type NoteDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns NoteColumns // columns contains all the column names of Table for convenient usage.
}

// NoteColumns defines and stores column names for table t_note.
type NoteColumns struct {
	Id        string //
	UserId    string // 用户ID
	Content   string // 内容
	Images    string // 图片
	Location  string // 位置
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// noteColumns holds the columns for table t_note.
var noteColumns = NoteColumns{
	Id:        "id",
	UserId:    "user_id",
	Content:   "content",
	Images:    "images",
	Location:  "location",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewNoteDao creates and returns a new DAO object for table data access.
func NewNoteDao() *NoteDao {
	return &NoteDao{
		group:   "default",
		table:   "t_note",
		columns: noteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NoteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NoteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NoteDao) Columns() NoteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NoteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NoteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NoteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
