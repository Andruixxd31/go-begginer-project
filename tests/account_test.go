// go:build e2e
package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func createToken() string {
    token := jwt.New(jwt.SigningMethodHS256)
    tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))
    if err != nil {
        fmt.Println(err)
    }
    return tokenString
}


func TestPostComment(t *testing.T) {
    t.Run("can post comment", func(t *testing.T) {
        client := resty.New()
        bearer := fmt.Sprintf("bearer %v", createToken())
        fmt.Println("bearer: ", bearer)
        resp, err := client.R().
            SetHeader("Authorization", bearer).
            SetBody(`{"name": "Juan"}`).
            Post("http://localhost:8080/api/v1/account")
        assert.NoError(t, err)
        assert.Equal(t, 200, resp.StatusCode())
        
    })

    t.Run("cannot post comment without JWT", func(t *testing.T) {
        client := resty.New()
        resp, err := client.R().
            SetBody(`{"name": "Juan"}`).
            Post("http://localhost:8080/api/v1/account")
        assert.NoError(t, err)
        assert.Equal(t, 401, resp.StatusCode())
        
    })
}


