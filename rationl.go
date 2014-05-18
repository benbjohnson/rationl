package rationl

import (
	"fmt"
	"os"

	"code.google.com/p/gogoprotobuf/proto"
)

func (u *User) SetID(v int64) {
	u.ID = proto.Int64(v)
}

func (u *User) SetEmail(v string) {
	u.Email = proto.String(v)
}

func (u *User) SetAccessToken(v string) {
	u.AccessToken = proto.String(v)
}

func (i *Investigation) SetID(v string) {
	i.ID = proto.String(v)
}

func (i *Investigation) SetUserID(v int64) {
	i.UserID = proto.Int64(v)
}

func (i *Investigation) SetName(v string) {
	i.Name = proto.String(v)
}

func warn(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
}

func warnf(msg string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", v...)
}
