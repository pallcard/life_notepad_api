// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Note is the golang structure for table note.
type Note struct {
	Id        int         `json:"id"        description:""`
	UserId    int         `json:"userId"    description:"用户ID"`
	Content   string      `json:"content"   description:"内容"`
	Images    string      `json:"images"    description:"图片"`
	Location  string      `json:"location"  description:"头像"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" description:""`
}
