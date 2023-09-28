package cmd

import (
	"github.com/spf13/cobra"
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
	//result := database.DB.Create(roleHasPermissions)
	//console.Success(cast.ToString(result.RowsAffected))
	//query := database.DB.Model(&user.User{})
	//rules := map[string]string{
	//	"name":  "like",
	//	"state": "=",
	//	"id":    "or,=",
	//	"phone": "and,=",
	//}
	//var sl []string
	//var w string
	//for k, v := range rules {
	//	//fmt.Printf("%s\t%s\n", k, v)
	//	//query += k + " " + v + "?" + "abc"
	//	wh := v
	//	if strings.Contains(v, ",") {
	//		rng := strings.Split(v, ",")
	//		wh = rng[1]
	//	}
	//
	//	switch wh {
	//	case "like":
	//		//query.Where(v+" LIKE ?", "ab")
	//		w = k + " LIKE ?"
	//	case "bet":
	//		//query.Where(k+"BETWEEN ? AND ?", "", "")
	//		w = k + " BETWEEN ? AND ?"
	//	case "in":
	//		//query.Where(k+"BETWEEN ? AND ?", "", "")
	//		w = k + " IN ?"
	//	default:
	//		w = k + wh + "?"
	//	}
	//
	//	//if strings.Contains(v, "or,") {
	//	//	query.Or(w)
	//	//} else {
	//	//	query.Where(w)
	//	//}
	//	sl = append(sl, w)
	//}
	//fmt.Println(sl)

	//time1 := time2.Now().Unix()
	//fmt.Println(reflect.ValueOf(time1).Kind())
	//fmt.Println(time2.Unix(time1, 0).Format("2006-01-02 15:04:05"))
	//a2 := "2022-1-2"
	//times := helpers.TimeUnix(a2)
	//str := "s"
	//fmt.Println(cast.ToUint64(str))
	//database.DB.First(user.User{}, cast.ToUint64(str))
}
