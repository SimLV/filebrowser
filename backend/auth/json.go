package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gtsteffaniak/filebrowser/backend/common/logger"
	"github.com/gtsteffaniak/filebrowser/backend/database/users"
)

type jsonCred struct {
	Password  string `json:"password"`
	Username  string `json:"username"`
	ReCaptcha string `json:"recaptcha"`
}

// JSONAuth is a json implementation of an Auther.
type JSONAuth struct {
	ReCaptcha *ReCaptcha `json:"recaptcha" yaml:"recaptcha"`
}

// Auth authenticates the user via a json in content body.
func (a JSONAuth) Auth(r *http.Request, userStore *users.Storage) (*users.User, error) {
	var cred jsonCred

	if r.Body == nil {
		logger.Error("nil body error")
		return nil, os.ErrPermission
	}
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		logger.Error("decode body error")
		return nil, os.ErrPermission
	}

	// If ReCaptcha is enabled, check the code.
	if a.ReCaptcha != nil && len(a.ReCaptcha.Secret) > 0 {
		ok, err := a.ReCaptcha.Ok(cred.ReCaptcha) //nolint:govet

		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, os.ErrPermission
		}
	}
	u, err := userStore.Get(cred.Username)
	if err != nil {
		return nil, fmt.Errorf("unable to get user from store: %v", err)
	}
	err = users.CheckPwd(cred.Password, u.Password)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// LoginPage tells that json auth doesn't require a login page.
func (a JSONAuth) LoginPage() bool {
	return true
}

const reCaptchaAPI = "/recaptcha/api/siteverify"

// ReCaptcha identifies a recaptcha connection.
type ReCaptcha struct {
	Host   string `json:"host"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// Ok checks if a reCaptcha responde is correct.
func (r *ReCaptcha) Ok(response string) (bool, error) {
	body := url.Values{}
	body.Set("secret", r.Secret)
	body.Add("response", response)

	client := &http.Client{}

	resp, err := client.Post(
		r.Host+reCaptchaAPI,
		"application/x-www-form-urlencoded",
		strings.NewReader(body.Encode()),
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	var data struct {
		Success bool `json:"success"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return false, err
	}

	return data.Success, nil
}
