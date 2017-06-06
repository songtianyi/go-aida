package main
import (
	"github.com/gin-gonic/gin"
	"github.com/songtianyi/go-aida/restful/logic"
)

func main() {
	router := gin.Default()

	router.GET("/qrcode", logic.Qrcode)
	router.GET("/status", logic.Status)

	router.Run();
}
