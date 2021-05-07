package mangadexv5

import "fmt"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginToken struct {
	Session string `json:"session"`
	Refresh string `json:"refresh"`
}

type LoginResult struct {
	Result string       `json:"result"`
	Token  *LoginToken  `json:"token"`
	Errors APIErrorList `json:"errors"`
}

func (c *Client) Login(username, password string) error {
	result := &LoginResult{}
	err := c.post("/auth/login", LoginRequest{Username: username, Password: password}, result)
	if err != nil {
		return err
	}

	if result.Result != "ok" {
		if result.Errors == nil {
			return fmt.Errorf("an unknown error occered")
		}
		return result.Errors
	}

	c.token = result.Token.Session
	return nil
}

func (c *Client) Token() string {
	return c.token
}

func (c *Client) SetToken(token string) {
	c.token = token
}
