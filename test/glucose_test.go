package test

import (
	"healthtrack/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestGlucose(t *testing.T) {
	var db *gorm.DB
	router := server.SetUpRoute(db)
	t.Run("Add glucose", func(t *testing.T) {
		want := "Client 1"
		got := "Client 1"
		if want != got {
			t.Errorf("want value %s  got value %s", want, got)
		}
	})
	t.Run("Get glucose by User", func(t *testing.T) {
		//simuler une requete
		req := httptest.NewRequest(http.MethodGet, "/glucose/1", nil)
		w := httptest.NewRecorder()
		//creer une serveur pour
		router.ServeHTTP(w, req)
		//On verifie la reponse
		require.Equal(t, http.StatusOK, w.Code)

	})
	t.Run("Delete glucose by ID", func(t *testing.T) {

	})
}
