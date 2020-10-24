package internal_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	a "soccer-api/internal"

	"github.com/stretchr/testify/require"
)

func TestIndexRoute(t *testing.T) {
	tests := map[string]struct {
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		"index route": {
			route:         "/",
			expectedError: false,
			expectedCode:  http.StatusOK,
			expectedBody:  `ok`,
		},
	}

	app := a.Setup()

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, test.route, nil)
			res, err := app.Test(req, -1)
			if err != nil {
				require.NoError(t, err)
				return
			}

			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)
			require.Equal(t, test.expectedCode, res.StatusCode)
			require.Equalf(t, test.expectedBody, string(body), "test index route")
		})
	}
}
