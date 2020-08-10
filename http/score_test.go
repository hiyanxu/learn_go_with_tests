package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const jsonContentType = "application/json"

func TestPlayerServer(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}}
	server := &PlayerServer{store: &store}

	t.Run("Pepper", func(t *testing.T) {
		// 创建一个request method:get 请求路径：/player/Pepper
		//request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil
		request := newGetScoreRequest("Pepper")
		// 返回ResponseRecorder，该struct实现了http.ResponseWriter，可以用来记录response
		response := httptest.NewRecorder()

		// 创建一个请求并调用到handler中
		//PlayerServer(response, request)
		server.ServeHTTP(response, request)

		assertHttpStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
		//got := response.Body.String()
		//want := "20"
		//
		//if got != want {
		//	t.Errorf("got '%s', want '%s'", got, want)
		//}
	})

	t.Run("Floyd", func(t *testing.T) {
		//request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		//PlayerServer(response, request)
		server.ServeHTTP(response, request)

		assertHttpStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
		//got := response.Body.String()
		//want := "10"
		//
		//if got != want {
		//	t.Errorf("got '%s', want '%s'", got, want)
		//}
	})

	t.Run("none", func(t *testing.T) {
		request := newGetScoreRequest("none")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		//got := response.Code
		//want := http.StatusNotFound
		//if got != want {
		//	t.Errorf("got status %d want %d", got, want)
		//}
		assertHttpStatus(t, response.Code, http.StatusNotFound)
		//assertResponseBody(t, response.Body.String(), "")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertHttpStatus(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{scores: map[string]int{}}
	//server := &PlayerServer{&store}
	server := NewPlayerServer(&store)

	t.Run("post", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertHttpStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != "Pepper" {
			t.Errorf("did not store correct winner got '%s' want '%s'", store.winCalls[0], player)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	player := "Pepper"
	store := InMemoryPlayerStore{store: map[string]int{
		player: 0,
	}}
	//server := PlayerServer{&store}
	server := NewPlayerServer(&store)

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertHttpStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")

}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayerStore{nil, nil, wantedLeague}
	//server := &PlayerServer{store: &store}
	server := NewPlayerServer(&store)

	t.Run("/league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertContentType(t, response, jsonContentType)

		var got []Player = getLeagueFromResponse(t, response.Body)
		assertHttpStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
	})
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()

	// 将返回的json数据解析为特定的类型
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()

	if response.Header().Get("content-type") != "application/json" {
		t.Errorf("response did not have content-type of application/json, got %v", response.HeaderMap)
	}
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertHttpStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertHttpStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})
}

func newLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}
