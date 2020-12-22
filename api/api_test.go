package api

import (
	"encoding/base64"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)



func Test_addNewArticle(t *testing.T) {

}

func Test_getAllArticles(t *testing.T) {
	go StartAPI(":8080")

	testCases := []struct{
		Method string
		ExpectedStatusCode int
	}{
		{
			Method: "GET",
			ExpectedStatusCode: http.StatusOK,
		},
	}

	for {
		_, err := net.Listen("tcp", ":8080")
		if err == nil {
			break
		}
	}

	for index, testCase := range testCases {
		request, err := http.NewRequest(testCase.Method, "http://localhost:8080/api/articles", nil)
		request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte("admin:admin")))

		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		getAllArticles(response, request)

		fmt.Println(response.Body)

		res := response.Result()

		if res.StatusCode != testCase.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, testCase.ExpectedStatusCode, res.Status)
		}
	}
}

func Test_getSingleArticle(t *testing.T) {

}

func Test_updateArticle(t *testing.T) {

}

func Test_deleteArticle(t *testing.T) {

}