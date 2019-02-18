package mainpage

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

var testFixture = []struct {
	url string
	method string
	httpStatus int
}{
	{"/", "GET", http.StatusOK},
	{"/", "POST", http.StatusMethodNotAllowed},
}

func TestGetMainPage(t *testing.T) {

	for key, fixture := range testFixture {

		req, err := http.NewRequest(fixture.method, fixture.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetForm)
		handler.ServeHTTP(rr, req)

		assert.Equalf(
			t,
			fixture.httpStatus,
			rr.Code,
			"Handler returned wrong status code: got %v want %v, test_case: %v",
			rr.Code,
			fixture.httpStatus,
			key,
		)

	}
}