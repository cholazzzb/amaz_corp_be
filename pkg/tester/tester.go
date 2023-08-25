package tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cholazzzb/amaz_corp_be/internal/app"
	"github.com/cholazzzb/amaz_corp_be/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type MockTest struct {
	Description     string
	method          string
	route           string
	body            map[string]interface{}
	ExpectedCode    int
	ExpectedMessage string
	ExpectedData    interface{}
	withAuth        bool
	Request         *http.Request
}

func NewMockTest() *MockTest {
	return &MockTest{}
}

func (b *MockTest) Desc(desc string) *MockTest {
	b.Description = desc
	return b
}

func (b *MockTest) GET() *MockTest {
	b.method = http.MethodGet
	return b
}

func (b *MockTest) POST() *MockTest {
	b.method = http.MethodPost
	return b
}

func (b *MockTest) PUT() *MockTest {
	b.method = http.MethodPut
	return b
}

func (b *MockTest) Route(route string) *MockTest {
	b.route = route
	return b
}

func (b *MockTest) Body(body map[string]interface{}) *MockTest {
	b.body = body
	return b
}

func (b *MockTest) Expected(
	expectedCode int,
	expectedMessage string,
	expectedData interface{},
) *MockTest {
	b.ExpectedCode = expectedCode
	b.ExpectedMessage = expectedMessage
	b.ExpectedData = expectedData
	return b
}

func (b *MockTest) WithAuth() *MockTest {
	b.withAuth = true
	return b
}

func (b *MockTest) BuildRequest() *http.Request {
	switch {
	case b.method == http.MethodGet:
		return httptest.NewRequest(b.method, b.route, nil)
	default:
		body, err := json.Marshal(b.body)
		if err != nil {
			panic(fmt.Errorf("failed to marshal %v", b.body))
		}
		req := httptest.NewRequest(b.method, b.route, bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		return req
	}
}

type MockTester struct {
	Tests       []MockTest
	app         *fiber.App
	withAuth    bool
	bearerToken string
}

func (mc *MockTester) Setup(envPath string) {
	config.GetEnv(envPath)
	mc.app = app.GetApp()
}

func (mc *MockTester) Test(t *testing.T) {
	if mc.withAuth {
		// TODO: Make sure the user is registered
		body, err := json.Marshal(map[string]interface{}{
			"username": "testing1",
			"password": "testing1",
		})
		if err != nil {
			panic(fmt.Errorf("failed to marshal %v", body))
		}
		req := httptest.NewRequest(http.MethodPost, "http://localhost:8000/api/v1/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := mc.app.Test(req, 10000)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println("respond body:", bodyString)
		fmt.Println()

		type BearerToken struct {
			Token string `json:"token"`
		}
		bearer := BearerToken{}
		json.Unmarshal(bodyBytes, &bearer)
		mc.bearerToken = bearer.Token
	}

	for _, test := range mc.Tests {
		req := test.BuildRequest()

		if test.withAuth {
			req.Header.Add("Authorization", "Bearer "+mc.bearerToken)
		}
		resp, err := mc.app.Test(req, 10000)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println("respond body:", bodyString)
		fmt.Println()

		assert.Equalf(t, test.ExpectedCode, resp.StatusCode, test.Description)
	}
}

func (mc *MockTester) AddTest(test MockTest) {
	mc.Tests = append(mc.Tests, test)
	if test.withAuth {
		mc.withAuth = true
	}
}
