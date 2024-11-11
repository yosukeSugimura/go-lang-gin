package main

import (
	"encoding/json"
	"gin_docker/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// テスト用のエンジンとルートをセットアップ
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// テンプレートディレクトリ設定（必要であれば）
	router.LoadHTMLGlob("templates/*")

	// SeimeiServiceのモックインスタンス
	seimeiService := service.NewSeimeiService()

	// ルートとエンドポイント設定
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/seimei/:name/:sei", func(c *gin.Context) {
		name := c.Param("name")
		sei := c.Param("sei")
		strokeCounts, err := seimeiService.GetStrokesForEachCharacter(name, sei)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"stroke_counts": strokeCounts})
	})

	router.GET("/seimei/:name/:sei/grids", func(c *gin.Context) {
		name := c.Param("name")
		sei := c.Param("sei")
		strokeCounts, err := seimeiService.GetStrokesForEachCharacter(name, sei)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tenkaku, jinkaku, chikaku, gaikaku, sokaku, err := seimeiService.CalculateFiveGrids(strokeCounts)
		if err != nil {
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
	})

	return router
}

// ルートエンドポイントのテスト
func TestRootEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// 各文字の画数を取得するエンドポイントのテスト
func TestGetStrokesEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/seimei/太郎/山田", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string][]int
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expected := []int{3, 5, 4, 9} // 仮の画数リスト
	assert.Equal(t, expected, response["stroke_counts"])
}

// 五格を計算するエンドポイントのテスト
func TestCalculateFiveGridsEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/seimei/太郎/山田/grids", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]int
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 期待される五格の値を仮定
	expected := map[string]int{
		"tenkaku": 8,
		"jinkaku": 9,
		"chikaku": 13,
		"gaikaku": 12,
		"sokaku":  21,
	}

	for key, expectedValue := range expected {
		assert.Equal(t, expectedValue, response[key])
	}
}
