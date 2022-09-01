package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"

	rl "github.com/adnaan/renderlayout"
)

func insertClinicalTrial() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		//userID := authn.AccountIDFromContext(r)
		return rl.D{}, nil
	}
}

func viewClinicalTrial() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		//userID := authn.AccountIDFromContext(r)
		return rl.D{}, nil
	}
}

func getSupportingContent() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		//userID := authn.AccountIDFromContext(r)
		return rl.D{}, nil
	}
}

func pcvPage() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		return rl.D{
			"title": "Graphical View",
		}, nil
	}
}

func homePage() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		return rl.D{
			"title": "WISSPAR",
		}, nil
	}
}

func clinicalTrialsOverviewPage() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		return rl.D{
			"title": "Clinical Trials",
		}, nil
	}
}

func lookupTrialPage() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		return rl.D{
			"title": "Look Up A Trial",
		}, nil
	}
}

func dataExportPage() rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		return rl.D{
			"title": "Immunogenicity Data Export",
		}, nil
	}
}

func overviewPage(appCtx Context) rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		signingToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"resource": struct {
				Dashboard int `json:"dashboard"`
			}{Dashboard: 1},
			"params": struct {}{},
			"exp": time.Now().Add(time.Minute * 10).Unix(), // 10 minute expiration
		})

		// Sign and get the complete encoded token as a string using the secret
		token, err := signingToken.SignedString([]byte(appCtx.cfg.MetabaseSecret))
		if err != nil {
			return nil, fmt.Errorf("failed at sign the token err: %w", err)
		}

		iframeUrl := fmt.Sprintf("%s/embed/dashboard/%s#bordered=true&titled=true", appCtx.cfg.MetabaseURL, token)
		return rl.D{
			"title": "",
			"iframeURL": iframeUrl,
		}, nil
	}
}

func clinicalTrialComparisonPage(appCtx Context) rl.Data {
	return func(w http.ResponseWriter, r *http.Request) (rl.D, error) {
		signingToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"resource": struct {
				Dashboard int `json:"dashboard"`
			}{Dashboard: 2},
			"params": struct {}{},
			"exp": time.Now().Add(time.Minute * 10).Unix(), // 10 minute expiration
		})

		// Sign and get the complete encoded token as a string using the secret
		token, err := signingToken.SignedString([]byte(appCtx.cfg.MetabaseSecret))
		if err != nil {
			return nil, fmt.Errorf("failed at sign the token err: %w", err)
		}

		iframeUrl := fmt.Sprintf("%s/embed/dashboard/%s#bordered=true&titled=true", appCtx.cfg.MetabaseURL, token)
		return rl.D{
			"iframeURL": iframeUrl,
		}, nil
	}
}
