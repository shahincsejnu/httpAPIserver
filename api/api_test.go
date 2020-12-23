package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)



func Test_addNewArticle(t *testing.T) {
	go StartAPI(":8080")

	testCases := []struct{
		Method string
		URL string
		Body io.Reader
		Username string
		Password string
		//Auth string
		ExpectedStatusCode int
	}{
		{
			Method: "POST",
			URL : "http://localhost:8080/api/article",
			Body : strings.NewReader(`{"ID" : "6", "Title" : "oka", "Body" : "hello", "Author" : {"ID" : "20", "Name" : "keu na", "Rating" : 20}}`),
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			ExpectedStatusCode: http.StatusCreated,
		},
		{
			Method : "POST",
			URL : "http://localhost:8080/api/article",
			Body : strings.NewReader(`{"ID" : "8", "Title" : "oka", "Body" : "hello", "Author" : {"ID" : "20", "Name" : "keu na", "Rating" : 20}}`),
			//Auth : "admin:noadmin",
			Username: "admin",
			Password: "noadmin",
			ExpectedStatusCode: http.StatusUnauthorized,
		},
	}

	for {
		_, err := net.Listen("tcp", ":8080")
		if err == nil {
			break
		}
	}

	for index, testCase := range testCases {
		request, err := http.NewRequest(testCase.Method, testCase.URL, testCase.Body)
		//request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(testCase.Auth)))
		request.SetBasicAuth(testCase.Username, testCase.Password)

		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		addNewArticle(response, request)

		fmt.Println(response.Body)

		res := response.Result()

		if res.StatusCode != testCase.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, testCase.ExpectedStatusCode, res.Status)
		}
	}
}

func Test_getAllArticles(t *testing.T) {
	go StartAPI(":8080")

	testCases := []struct{
		Method string
		URL string
		//Auth string
		Username string
		Password string
		ExpectedStatusCode int
	}{
		{
			Method: "GET",
			URL : "http://localhost:8080/api/articles",
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method: "GET",
			URL : "http://localhost:8080/api/articles",
			//Auth: "admin:1234",
			Username: "admin",
			Password: "1234",
			ExpectedStatusCode: http.StatusUnauthorized,
		},
	}

	for {
		_, err := net.Listen("tcp", ":8080")
		if err == nil {
			break
		}
	}

	for index, testCase := range testCases {
		request, err := http.NewRequest(testCase.Method, testCase.URL, nil)
		//request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(testCase.Auth)))

		request.SetBasicAuth(testCase.Username, testCase.Password)

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
	go StartAPI(":8080")

	testCases := []struct{
		Method string
		URL string
		//Auth string
		Username string
		Password string
		vars map[string]string
		ExpectedStatusCode int
	}{
		{
			Method: "GET",
			URL : "http://localhost:8080/api/article",
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			vars : map[string]string{
				"id" : "2",
			},
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method: "GET",
			URL : "http://localhost:8080/api/article",
			//Auth : "admin:noadmin",
			Username: "admin",
			Password: "noadmin",
			vars : map[string]string{
				"id" : "2",
			},
			ExpectedStatusCode: http.StatusUnauthorized,
		},
		{
			Method: "GET",
			URL : "http://localhost:8080/api/article",
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			vars : map[string]string{
				"id" : "200",
			},
			ExpectedStatusCode: http.StatusNoContent,
		},
	}

	for {
		_, err := net.Listen("tcp", ":8080")
		if err == nil {
			break
		}
	}

	for index, testCase := range testCases {
		request, err := http.NewRequest(testCase.Method, testCase.URL, nil)
		//request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(testCase.Auth)))
		request.SetBasicAuth(testCase.Username, testCase.Password)

		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		//vars := map[string]string {
		//	"id" : "2",
		//}

		request = mux.SetURLVars(request, testCase.vars)

		getSingleArticle(response, request)

		fmt.Println(response.Body)

		res := response.Result()

		if res.StatusCode != testCase.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, testCase.ExpectedStatusCode, res.Status)
		}
	}
}

func Test_updateArticle(t *testing.T) {
	go StartAPI(":8080")

	testCases := []struct{
		Method string
		URL string
		Body io.Reader
		vars map[string]string
		//Auth string
		Username string
		Password string
		ExpectedStatusCode int
	}{
		{
			Method: "PUT",
			URL : "http://localhost:8080/api/article",
			Body : strings.NewReader(`{"ID" : "6", "Title" : "oka", "Body" : "hello", "Author" : {"ID" : "20", "Name" : "keu na", "Rating" : 20}}`),
			vars : map[string]string {
				"id" : "1",
			},
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			ExpectedStatusCode: http.StatusCreated,
		},
		{
			Method: "PUT",
			URL : "http://localhost:8080/api/article",
			Body : strings.NewReader(`{"ID" : "6", "Title" : "oka", "Body" : "hello", "Author" : {"ID" : "20", "Name" : "keu na", "Rating" : 20}}`),
			vars : map[string]string {
				"id" : "202",
			},
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			ExpectedStatusCode: http.StatusNoContent,
		},
		{
			Method: "PUT",
			URL : "http://localhost:8080/api/article",
			Body : strings.NewReader(`{"ID" : "6", "Title" : "oka", "Body" : "hello", "Author" : {"ID" : "20", "Name" : "keu na", "Rating" : 20}}`),
			vars : map[string]string {
				"id" : "1",
			},
			//Auth : "admin:noadmin",
			Username: "admin",
			Password: "noadmin",
			ExpectedStatusCode: http.StatusUnauthorized,
		},
	}

	for {
		_, err := net.Listen("tcp", ":8080")
		if err == nil {
			break
		}
	}

	for index, testCase := range testCases {
		request, err := http.NewRequest(testCase.Method, testCase.URL, testCase.Body)
		//request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(testCase.Auth)))
		request.SetBasicAuth(testCase.Username, testCase.Password)

		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		request = mux.SetURLVars(request, testCase.vars)

		updateArticle(response, request)

		fmt.Println(response.Body)

		res := response.Result()

		if res.StatusCode != testCase.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, testCase.ExpectedStatusCode, res.Status)
		}
	}
}

func Test_deleteArticle(t *testing.T) {
	go StartAPI(":8080")

	testCases := []struct{
		Method string
		URL string
		vars map[string]string
		//Auth string
		Username string
		Password string
		ExpectedStatusCode int
	}{
		{
			Method: "DELETE",
			URL : "http://localhost:8080/api/article",
			vars : map[string]string{
				"id" : "1",
			},
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			ExpectedStatusCode: http.StatusOK,
		},
		{
			Method: "DELETE",
			URL : "http://localhost:8080/api/article",
			vars : map[string]string{
				"id" : "101",
			},
			//Auth : "admin:admin",
			Username: "admin",
			Password: "admin",
			ExpectedStatusCode: http.StatusNoContent,
		},
		{
			Method: "DELETE",
			URL : "http://localhost:8080/api/article",
			vars : map[string]string{
				"id" : "1",
			},
			//Auth : "admin:noadmin",
			Username: "admin",
			Password: "noadmin",
			ExpectedStatusCode: http.StatusUnauthorized,
		},
	}

	for {
		_, err := net.Listen("tcp", ":8080")
		if err == nil {
			break
		}
	}

	for index, testCase := range testCases {
		request, err := http.NewRequest(testCase.Method, testCase.URL, nil)
		//request.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(testCase.Auth)))
		request.SetBasicAuth(testCase.Username, testCase.Password)

		if err != nil {
			t.Fatal(err)
		}

		response := httptest.NewRecorder()

		request = mux.SetURLVars(request, testCase.vars)

		deleteArticle(response, request)

		fmt.Println(response.Body)

		res := response.Result()

		if res.StatusCode != testCase.ExpectedStatusCode {
			t.Errorf("Case %v: expected %v got %v", index, testCase.ExpectedStatusCode, res.Status)
		}
	}
}