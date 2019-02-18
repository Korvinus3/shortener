package shortener

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"net/url"
	"strings"
	"errors"
	"github.com/gorilla/mux"
)

var testProcessLinkFixture = []struct {
	link string
	storageKey string
	storageErr error
	httpStatus int
}{
	{"http://test_url", "testKey", nil,http.StatusOK},
	{"http://test_url",  "", nil,http.StatusBadRequest},
	{"", "test_key", nil,http.StatusBadRequest},
	{"http://test_url", "test_key", errors.New("error message"),http.StatusBadRequest},
}


func TestProcessLink(t *testing.T) {

	oldMethod := writeData
	defer func() {writeData = oldMethod}()

	for key, fixture := range testProcessLinkFixture {

		writeData = func(value string) (string, error) {
			return fixture.storageKey, fixture.storageErr
		}

		form := url.Values{}
		form.Add("link", fixture.link)

		req, err := http.NewRequest("POST", "/short-link/", strings.NewReader(form.Encode()))

		req.PostForm = form

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProcessLink)
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

var testProcessRedirectFixture = []struct {
	url string
	originalURLFromStorage string
	httpStatus int
}{
	{"test", "testKey", http.StatusFound},
	{"test",  "",http.StatusNotFound},
	{"", "test_key",http.StatusBadRequest},
}


func TestProcessRedirect(t *testing.T) {

	oldMethod := getURL
	defer func() {getURL = oldMethod}()

	for key, fixture := range testProcessRedirectFixture {

		getURL = func(key string) string {
			return fixture.originalURLFromStorage
		}

		req, err := http.NewRequest("Get", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		req = mux.SetURLVars(req,map[string]string{
			"key" : fixture.url,
		})

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ProcessRedirect)
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