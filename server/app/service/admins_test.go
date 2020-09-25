package service

import (
	"fmt"
	"server/app/model/admins"
	"server/library/global"
	"testing"

	"github.com/gogf/gf/database/gdb"

	"github.com/gogf/gf/frame/g"
)

func TestGetAdminList(t *testing.T) {
	var adminList []admins.AdminHasOneAuthority
	err := g.DB(global.Db).Table("admins").ScanList(&adminList, "AdminTest")
	if err != nil {
		panic(err)
	}
	err = g.DB(global.Db).Table("admins").
		Where("authority_id", gdb.ListItemValues(adminList, "AdminTest", "AuthorityId")).
		ScanList(&adminList, "Authority", "AdminTest", "authority_id:AuthorityId")
	if err != nil {
		panic(err)
	}
	fmt.Println(adminList)
}
