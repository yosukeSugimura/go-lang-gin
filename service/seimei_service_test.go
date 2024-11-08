package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockServerはAPIレスポンスを模擬するために使用
func MockServer(t *testing.T, stroke string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"stroke":"` + stroke + `"}]`))
	}))
}

// TestGetStrokeCountはGetStrokeCountメソッドをテスト
func TestGetStrokeCount(t *testing.T) {
	mockServer := MockServer(t, "5")
	defer mockServer.Close()

	service := NewSeimeiService()

	// モックサーバーのURLを設定
	apiUrl := mockServer.URL + "/kanji/v1/for-child/chars/山"
	strokeCount, err := service.GetStrokeCount(apiUrl)
	if err != nil {
		t.Fatalf("画数の取得に失敗しました: %v", err)
	}

	expected := 5
	if strokeCount != expected {
		t.Errorf("期待される画数: %d, 実際の画数: %d", expected, strokeCount)
	}
}

// TestGetStrokesForEachCharacterはGetStrokesForEachCharacterメソッドをテスト
func TestGetStrokesForEachCharacter(t *testing.T) {
	mockServer := MockServer(t, "3")
	defer mockServer.Close()

	service := NewSeimeiService()

	strokeCounts, err := service.GetStrokesForEachCharacter("太郎", "山田")
	if err != nil {
		t.Fatalf("各文字ごとの画数取得に失敗しました: %v", err)
	}

	expectedCounts := []int{3, 3, 3, 3}
	for i, count := range strokeCounts {
		if count != expectedCounts[i] {
			t.Errorf("期待される画数: %d, 実際の画数: %d", expectedCounts[i], count)
		}
	}
}

// TestCalculateFiveGridsはCalculateFiveGridsメソッドをテスト
func TestCalculateFiveGrids(t *testing.T) {
	service := NewSeimeiService()
	strokeCounts := []int{3, 3, 5, 3} // 山田太郎 の画数と仮定

	tenkaku, jinkaku, chikaku, gaikaku, sokaku, err := service.CalculateFiveGrids(strokeCounts)
	if err != nil {
		t.Fatalf("五格計算に失敗しました: %v", err)
	}

	// 期待される五格の計算結果
	expectedTenkaku := 6
	expectedJinkaku := 8
	expectedChikaku := 8
	expectedGaikaku := 6
	expectedSokaku := 14

	if tenkaku != expectedTenkaku {
		t.Errorf("期待される天格: %d, 実際の天格: %d", expectedTenkaku, tenkaku)
	}
	if jinkaku != expectedJinkaku {
		t.Errorf("期待される人格: %d, 実際の人格: %d", expectedJinkaku, jinkaku)
	}
	if chikaku != expectedChikaku {
		t.Errorf("期待される地格: %d, 実際の地格: %d", expectedChikaku, chikaku)
	}
	if gaikaku != expectedGaikaku {
		t.Errorf("期待される外格: %d, 実際の外格: %d", expectedGaikaku, gaikaku)
	}
	if sokaku != expectedSokaku {
		t.Errorf("期待される総格: %d, 実際の総格: %d", expectedSokaku, sokaku)
	}
}

// TestCalculateFiveGrids_ErrorはCalculateFiveGridsのエラーハンドリングをテスト
func TestCalculateFiveGrids_Error(t *testing.T) {
	service := NewSeimeiService()
	strokeCounts := []int{3} // 画数が不足している場合

	_, _, _, _, _, err := service.CalculateFiveGrids(strokeCounts)
	if err == nil {
		t.Fatal("エラーが発生するはずが、発生しませんでした")
	}
	expectedErr := "名前と姓の画数が不足しています"
	if err.Error() != expectedErr {
		t.Errorf("期待されるエラーメッセージ: %s, 実際のエラーメッセージ: %s", expectedErr, err.Error())
	}
}
