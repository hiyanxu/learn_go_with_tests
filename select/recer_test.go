package _select

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowUrl := "http://www.facebook.com"
	quickUrl := "http://www.quii.co.uk"

	want := quickUrl
	got := Racer(slowUrl, quickUrl)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestRacer2(t *testing.T) {
	// 实现func(w http.ResponseWriter, r *http.Request){}方法，即可以作为http handler
	//slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	time.Sleep(20 * time.Millisecond)
	//	w.WriteHeader(http.StatusOK)
	//}))
	//quickServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//}))
	slowServer := makeDelayedServer(20 * time.Millisecond)
	quickServer := makeDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer quickServer.Close()

	slowUrl := slowServer.URL
	quickUrl := quickServer.URL
	fmt.Println(slowUrl, quickUrl)

	want := quickUrl
	got := Racer(slowUrl, quickUrl)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

	// 采用defer，在程序结束时进行close
	//slowServer.Close()
	//quickServer.Close()
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer22(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		slowServer := makeDelayedServer(5 * time.Millisecond)
		quickServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer quickServer.Close()

		slowUrl := slowServer.URL
		quickUrl := quickServer.URL

		want := quickUrl
		got, _ := Racer3(slowUrl, quickUrl)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server does not respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 3)
		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}
