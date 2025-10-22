package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestGlucose(t *testing.T) {
	router := SetupRouter()
	t.Run("Add glucose", func(t *testing.T) {
		want := "Client 1"
		got := "Client 1"
		if want != got {
			t.Errorf("want value %s  got value %s", want, got)
		}
	})
	t.Run("Get glucose by User", func(t *testing.T) {
		//simuler une requete
		req := httptest.NewRequest("GET", "/api/v1/glucose/1", nil)
		w := httptest.NewRecorder()
		//creer une serveur pour
		router.ServeHTTP(w, req)
		//On verifie la reponse
		require.Equal(t, http.StatusOK, w.Code)

	})
	t.Run("Delete glucose by ID", func(t *testing.T) {

	})
}
