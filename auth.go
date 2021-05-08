package mangadexv5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/pkg/errors"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginToken struct {
	Session   string    `json:"session"`
	Refresh   string    `json:"refresh"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	Result string       `json:"result"`
	Token  *LoginToken  `json:"token"`
	Errors APIErrorList `json:"errors"`
}

func (c *Client) Login(username, password string) error {
	result := &LoginResponse{}
	err := c.post("/auth/login", LoginRequest{Username: username, Password: password}, result)
	if err != nil {
		return err
	}

	if result.Result != "ok" {
		if result.Errors == nil {
			return fmt.Errorf("an unknown error occurred")
		}
		return result.Errors
	}

	result.Token.CreatedAt = time.Now()

	c.SetToken(result.Token)
	return nil
}

func (c *Client) Token() *LoginToken {
	return c.token
}

func (c *Client) SetToken(token *LoginToken) error {
	c.token = token
	if c.tokenFile != "" {
		b, err := json.Marshal(c.token)
		if err != nil {
			return errors.Wrap(err, "failed to encode token json")
		}
		err = ioutil.WriteFile(c.tokenFile, b, 0755)
		if err != nil {
			return errors.Wrap(err, "failed to write token to disk")
		}
	}
	return nil
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
}
type RefreshTokenResponse struct {
	Result  string      `json:"result"`
	Token   *LoginToken `json:"token"`
	Message string      `json:"message"`
}

func (c *Client) RefreshToken(refresh string) error {
	result := &LoginResponse{}
	err := c.post("/auth/refresh", RefreshTokenRequest{Token: refresh}, result)
	if err != nil {
		return err
	}

	if result.Result != "ok" {
		if result.Errors == nil {
			return fmt.Errorf("an unknown error occered")
		}
		return result.Errors
	}

	result.Token.CreatedAt = time.Now()

	c.SetToken(result.Token)
	return nil
}

func (c *Client) Authenticate(username, password string) error {
	b, err := ioutil.ReadFile(c.tokenFile)
	if err != nil {
		return errors.Wrap(err, "could not open token file")
	}

	token := &LoginToken{}

	err = json.Unmarshal(b, token)
	if err != nil {
		return errors.Wrap(err, "could not parse token json")
	}

	if token.CreatedAt.After(time.Now().Add(-10 * time.Minute)) {
		log.Print("use existing token")
		c.token = token
	} else if token.CreatedAt.After(time.Now().Add(-4 * time.Hour)) {
		log.Print("refresh token")
		return c.RefreshToken(token.Refresh)
	} else {
		log.Print("login")
		return c.Login(username, password)
	}

	return nil
}
