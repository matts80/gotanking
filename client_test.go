package gotanking

import (
	"testing"
)

const (
	appID      = "dummy"
	realm      = "na"
	playerName = "dummy"
)

func TestClientSetup(t *testing.T) {

	// client sets realm appropriately
	t.Run("client sets realm appropriately", func(t *testing.T) {

		realmTests := []struct {
			realm string
			want  string
		}{
			{realm: "na", want: "https://api.worldoftanks.com/wot/"},
			{realm: "eu", want: "https://api.worldoftanks.eu/wot/"},
			{realm: "ru", want: "https://api.worldoftanks.ru/wot/"},
			{realm: "asia", want: "https://api.worldoftanks.asia/wot/"},
			{realm: "moon", want: "https://api.worldoftanks.com/wot/"},
		}

		for _, tt := range realmTests {
			got, _ := NewClient(appID, SetRealm(tt.realm))
			if got.baseURL != tt.want {
				t.Errorf("got %q want %q", got.baseURL, tt.want)
			}
		}
	})

	// default realm is NA
	t.Run("default realm is NA", func(t *testing.T) {
		got, _ := NewClient(appID)
		want := "https://api.worldoftanks.com/wot/"

		if got.baseURL != want {
			t.Errorf("got %q want %q", got.baseURL, want)
		}
	})

	// base URL can be changed
	t.Run("base URL can be changed", func(t *testing.T) {
		url := "http://localhost:8080/api/"
		got, _ := NewClient(appID, SetBaseURL("http://localhost:8080/api/"))

		if got.baseURL != url {
			t.Errorf("got %q want %q", got.baseURL, url)
		}

	})
}

// when we expect an error
func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error here")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
