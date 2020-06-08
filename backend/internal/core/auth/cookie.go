package auth

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

type CookieManager struct {
	instance *securecookie.SecureCookie
}

func newCookieManager(secretKey, blockKey string) *CookieManager {
	return &CookieManager{
		instance: securecookie.New([]byte(secretKey), []byte(blockKey)),
	}
}

func (c *CookieManager) Put(w http.ResponseWriter, name string, value map[string]string) error {
	encoded, err := c.instance.Encode(name, value)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:  name,
		Value: encoded,
		Path:  "/",
		// TODO: make this secure on production
		// Secure:  true,
		HttpOnly: true,
		Expires:  time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, cookie)
	return nil
}

func (c *CookieManager) Fetch(r *http.Request, name string) (map[string]string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}

	value := make(map[string]string)
	err = c.instance.Decode(name, cookie.Value, &value)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *CookieManager) Destroy(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	})
}
