package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerServer(t *testing.T) {
	store := StubPlayerStore{scores: map[string]string{
		"Pepper": "20",
		"Floyd":  "10",
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
	store := StubPlayerStore{scores: map[string]string{}}
	server := &PlayerServer{&store}

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
	store := InMemoryPlayerStore{}
	server := PlayerServer{&store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertHttpStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")

}
