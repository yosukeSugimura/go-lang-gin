// routes/routes.go
package routes

import (
	"gin_docker/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// テンプレートディレクトリを設定
	router.LoadHTMLGlob("templates/*")

	// コントローラのインスタンスを作成
	seimeiController := controllers.NewSeimeiController()

	// ルートページのエンドポイント
	router.GET("/", seimeiController.ShowIndex)

	// 各文字の画数を取得するエンドポイント
	router.GET("/seimei/:name/:sei", seimeiController.GetStrokeCounts)

	// 五格を計算するエンドポイント
	router.GET("/seimei/:name/:sei/grids", seimeiController.CalculateGrids)

	return router
}
