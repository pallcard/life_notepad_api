// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package note

import (
	"context"
	v1 "life_notepad_api/api/note/v1"
)

type INoteV1 interface {
	NoteList(ctx context.Context, req *v1.NoteListReq) (res *v1.NoteListRes, err error)
	AddNote(ctx context.Context, req *v1.AddNoteReq) (res *v1.AddNoteRes, err error)
}
