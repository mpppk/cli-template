package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mpppk/cli-template/pkg/infra/handler"

	"github.com/mpppk/cli-template/pkg/infra"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	// Setup
	e := infra.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/api/sum?a=1&b=2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := handler.New()

	res := handler.SumResponse{Result: 3}
	resContents, err := json.Marshal(res)
	resJson := fmt.Sprintln(string(resContents))
	if err != nil {
		t.Fail()
	}

	// Assertions
	if assert.NoError(t, h.Sum(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, resJson, rec.Body.String())
	}
}
