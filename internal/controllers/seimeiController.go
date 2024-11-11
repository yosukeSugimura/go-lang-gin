// controllers/seimeiController.go
package controllers

import (
	"gin_docker/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SeimeiController struct {
	seimeiService *service.SeimeiService
}

func NewSeimeiController() *SeimeiController {
	return &SeimeiController{
		seimeiService: service.NewSeimeiService(),
	}
}

// ルートページ表示
func (s *SeimeiController) ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 各文字の画数を取得
func (s *SeimeiController) GetStrokeCounts(c *gin.Context) {
	name := c.Param("name")
	sei := c.Param("sei")
	strokeCounts, err := s.seimeiService.GetStrokesForEachCharacter(name, sei)
	if err != nil {
		log.Printf("エラー: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stroke_counts": strokeCounts})
}

// 五格を計算
func (s *SeimeiController) CalculateGrids(c *gin.Context) {
	name := c.Param("name")
	sei := c.Param("sei")

	strokeCounts, err := s.seimeiService.GetStrokesForEachCharacter(name, sei)
	if err != nil {
		log.Printf("エラー: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tenkaku, jinkaku, chikaku, gaikaku, sokaku, err := s.seimeiService.CalculateFiveGrids(strokeCounts)
	if err != nil {
		log.Printf("エラー: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tenkaku": tenkaku,
		"jinkaku": jinkaku,
		"chikaku": chikaku,
		"gaikaku": gaikaku,
		"sokaku":  sokaku,
	})
}
