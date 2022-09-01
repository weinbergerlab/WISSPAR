package users

import (
	"context"
	"fmt"
	"github.com/adnaan/authn"
	"github.com/adnaan/authn/models"
	"github.com/adnaan/authn/models/account"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
	"net/http"
	"strings"
)

const defaultMaxAge = 60 * 60 * 24 * 30 // 30 days
const defaultPath = "/"
const sessionStoreKey = "auth"

type API struct {
	cfg          authn.Config
	sessionStore authn.SessionsStore

	// database client
	client *models.Client
}

// New authn client
func New(ctx context.Context, cfg authn.Config) *API {
	client, err := models.Open(cfg.Driver, cfg.Datasource)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}
	if cfg.SessionMaxAge == 0 {
		cfg.SessionMaxAge = defaultMaxAge
	}
	if cfg.SessionPath == "" {
		cfg.SessionPath = defaultPath
	}

	sessionStore := &defaultSessionStore{
		Ctx:    ctx,
		Client: client,
		Codecs: securecookie.CodecsFromPairs([]byte(cfg.SessionSecret)),
		SessionOpts: &sessions.Options{
			Path:   defaultPath,
			MaxAge: defaultMaxAge,
		}}

	return &API{
		cfg:          cfg,
		sessionStore: sessionStore,
		client:       client,
	}
}

//var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) error {
	if len(e) < 3 || len(e) > 254 {
		return fmt.Errorf("didn't match length needed: %d", len(e))
	}

	parts := strings.Split(e, "@")

	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return err
	}
	return nil
}

// hashPassword generates a hashed password from a plaintext string
func hashPassword(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pw), nil
}

// Signup a new account with email and password
func (a *API) Signup(ctx context.Context, email, password string, attributes map[string]interface{}) error {
	account, err := a.client.Account.Query().Where(account.Email(email)).First(ctx)
	if account == nil && err != nil && err.Error() == "models: account not found" {
		fmt.Println("Creating account...")

		hashedPassword, err := hashPassword(password)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		_, err = a.client.Account.Create().
			SetEmail(email).
			SetPassword(hashedPassword).
			SetProvider("email_signup").
			SetAttributes(attributes).
			SetConfirmed(true).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

// Login logs in the account with  email and password
func (a *API) Login(w http.ResponseWriter, r *http.Request, email, password string) error {
	if len(password) < 3 {
		return fmt.Errorf("ErrInvalidPassword")
	}

	if err := isEmailValid(email); err != nil {
		return err
	}

	acc, err := a.client.Account.Query().Where(account.Email(email)).First(r.Context())
	if err != nil {
		return err
	}

	if !acc.Confirmed {
		return fmt.Errorf("ErrEmailNotConfirmed")
	}

	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		return err
	}

	// set provider
	acc, err = acc.Update().SetProvider("email_signup").Save(r.Context())
	if err != nil {
		return err
	}

	session, err := a.sessionStore.Get(r, sessionStoreKey)
	if err != nil {
		return err
	}

	session.Values["id"] = acc.ID.String()

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}
