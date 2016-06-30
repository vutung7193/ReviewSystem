package main 
//tes
import (
	"fmt"
	"bitbucket.org/hblabvn/go_base/app"
)

func main () {
	fmt.Println(app.GetInstance().Config.Database.Host)
}