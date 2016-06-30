package main 
//tes
import (
	"fmt"
	"github.com/dieuhd/blog/app"
)

func main () {
	fmt.Println(app.GetInstance().Config.Database.Host)
}