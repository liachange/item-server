package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"item-server/app/models/role_has_permission"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	// 存进去 redis 中
	//redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	// 从 redis 里取出
	//console.Success(redis.Redis.Get("hello"))
	permKey := []uint64{2, 3, 4, 5}
	var roleHasPermissions []*role_has_permission.RoleHasPermission
	for _, v := range permKey {
		row := &role_has_permission.RoleHasPermission{
			PermissionID: v,
			RoleID:       3,
		}
		roleHasPermissions = append(roleHasPermissions, row)
	}
	fmt.Println(&roleHasPermissions)
	//result := database.DB.Create(roleHasPermissions)
	//console.Success(cast.ToString(result.RowsAffected))

}
