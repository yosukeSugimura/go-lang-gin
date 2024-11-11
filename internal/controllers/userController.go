package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	// ビジネスロジック
	c.JSON(http.StatusOK, gin.H{"message": "ユーザー情報取得成功"})
}
