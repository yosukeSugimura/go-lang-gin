// service/seimei_service.go
package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// KanjiResponse 構造体は、APIのレスポンスの構造に合わせて定義
type KanjiResponse struct {
	Stroke string `json:"stroke"`
}

// SeimeiService 構造体はAPI呼び出しを行うためのクラス
type SeimeiService struct{}

// NewSeimeiService 関数はSeimeiServiceのインスタンスを作成
func NewSeimeiService() *SeimeiService {
	return &SeimeiService{}
}

// GetStrokeCount メソッドは名前を引数に取り、APIから画数を取得して返す
func (s *SeimeiService) GetStrokeCount(name string) (int, error) {
	apiUrl := fmt.Sprintf("https://apino.yukiyuriweb.com/api/kanji/v1/for-child/chars/%s", name)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Printf("リクエストの作成に失敗しました: %v", err)
		return 0, fmt.Errorf("failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("APIリクエストが失敗しました: %v", err)
		return 0, fmt.Errorf("API request failed")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("APIエラー: ステータスコード %d, ステータスメッセージ %s", resp.StatusCode, resp.Status)
		return 0, fmt.Errorf("API request returned an error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("レスポンスボディの読み込みに失敗しました: %v", err)
		return 0, fmt.Errorf("failed to read response body")
	}

	var kanjiData []KanjiResponse
	if err := json.Unmarshal(body, &kanjiData); err != nil {
		log.Printf("レスポンスのデコードに失敗しました: %v", err)
		return 0, fmt.Errorf("failed to decode response")
	}

	if len(kanjiData) > 0 {
		strokeCount, err := strconv.Atoi(kanjiData[0].Stroke)
		if err != nil {
			log.Printf("画数の変換に失敗しました: %v", err)
			return 0, fmt.Errorf("failed to convert stroke count")
		}
		return strokeCount, nil
	} else {
		return 0, fmt.Errorf("stroke count not found")
	}
}

// GetStrokesForEachCharacter メソッドは文字列を受け取り、各文字ごとに画数を取得してリストを返す
func (s *SeimeiService) GetStrokesForEachCharacter(name string, sei string) ([]int, error) {
	strokeCounts := []int{}
	seimei := sei + name
	// 文字列を1文字ずつループ
	for _, char := range seimei {
		strokeCount, err := s.GetStrokeCount(string(char))
		if err != nil {
			log.Printf("画数の取得に失敗しました: %v", err)
			return nil, fmt.Errorf("failed to get stroke count for character: %c", char)
		}
		strokeCounts = append(strokeCounts, strokeCount)
	}

	return strokeCounts, nil
}

// CalculateFiveGrids メソッドは、五格（天格、人格、地格、外格、総格）を計算する
func (s *SeimeiService) CalculateFiveGrids(strokeCounts []int) (int, int, int, int, int, error) {
	if len(strokeCounts) < 2 {
		return 0, 0, 0, 0, 0, fmt.Errorf("名前と姓の画数が不足しています")
	}

	// 姓と名に分ける（前半が姓、後半が名と仮定）
	surnameStrokes := strokeCounts[:len(strokeCounts)/2]
	givenNameStrokes := strokeCounts[len(strokeCounts)/2:]

	// 天格: 姓の合計画数
	tenkaku := sumStrokes(surnameStrokes)

	// 人格: 姓の最後の文字と名の最初の文字の画数の合計
	jinkaku := surnameStrokes[len(surnameStrokes)-1] + givenNameStrokes[0]

	// 地格: 名の合計画数
	chikaku := sumStrokes(givenNameStrokes)

	// 総格: 姓と名の全ての画数の合計
	sokaku := sumStrokes(strokeCounts)

	// 外格: 総格から人格を引いた画数
	gaikaku := sokaku - jinkaku

	return tenkaku, jinkaku, chikaku, gaikaku, sokaku, nil
}

// sumStrokes は画数の合計を求めるヘルパー関数
func sumStrokes(strokes []int) int {
	total := 0
	for _, stroke := range strokes {
		total += stroke
	}
	return total
}
