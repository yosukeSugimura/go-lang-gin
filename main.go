// main.go
package main

import (
	"gin_docker/service" // 正しいパスに置き換えてください
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// テンプレートディレクトリを設定
	router.LoadHTMLGlob("templates/*")

	// ルートページのエンドポイント
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// SeimeiServiceのインスタンスを作成
	seimeiService := service.NewSeimeiService()

	// 各文字の画数を取得するエンドポイント
	router.GET("/seimei/:name/:sei", func(c *gin.Context) {
		name := c.Param("name")
		sei := c.Param("sei")
		// 各文字の画数を取得
		strokeCounts, err := seimeiService.GetStrokesForEachCharacter(name, sei)
		if err != nil {
			log.Printf("エラー: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 成功時のレスポンス
		c.JSON(http.StatusOK, gin.H{"stroke_counts": strokeCounts})
	})

	// 五格を計算するエンドポイント
	router.GET("/seimei/:name/:sei/grids", func(c *gin.Context) {
		name := c.Param("name")
		sei := c.Param("sei")

		// 各文字の画数を取得
		strokeCounts, err := seimeiService.GetStrokesForEachCharacter(name, sei)
		if err != nil {
			log.Printf("エラー: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 五格の計算
		tenkaku, jinkaku, chikaku, gaikaku, sokaku, err := seimeiService.CalculateFiveGrids(strokeCounts)
		if err != nil {
			log.Printf("エラー: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 成功時のレスポンス
		c.JSON(http.StatusOK, gin.H{
			"tenkaku": tenkaku,
			"jinkaku": jinkaku,
			"chikaku": chikaku,
			"gaikaku": gaikaku,
			"sokaku":  sokaku,
		})
	})

	// サーバー起動
	router.Run(":8080")
}
