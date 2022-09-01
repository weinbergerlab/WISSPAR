package app

import (
	"context"
	"fmt"
	"github.com/hako/branca"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/gen/models"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/sqlx"
	"gitlab.com/weinbergerlab/immunogenicity-data-project/app/users"
	"log"
	"net/http"

	"github.com/adnaan/authn"

	"github.com/go-playground/form"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	"github.com/go-chi/httplog"

	rl "github.com/adnaan/renderlayout"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

type Context struct {
	authn       *authn.API
	users       *users.API
	cfg         Config
	formDecoder *form.Decoder
	db          *models.Client
	branca      *branca.Branca
}

type APIRoute struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func Router(ctx context.Context, cfg Config) chi.Router {
	fmt.Println("opening models")
	db, err := sqlx.Open(cfg.Driver, cfg.DataSource)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("creating schema")
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("created schema")

	appCtx := Context{
		db:          db,
		cfg:         cfg,
		formDecoder: form.NewDecoder(),
		branca:      branca.NewBranca(cfg.APIMasterSecret),
	}

	authnConfig := authn.Config{
		Driver:        cfg.Driver,
		Datasource:    cfg.DataSource,
		SessionSecret: cfg.SessionSecret,
		SendMail:      sendEmailFunc(cfg),
		GothProviders: []goth.Provider{
			google.New(
				cfg.GoogleClientID,
				cfg.GoogleSecret,
				fmt.Sprintf("%s/auth/callback?provider=google", cfg.Domain),
				"email", "profile",
			),
		},
	}

	appCtx.authn = authn.New(ctx, authnConfig)

	appCtx.users = users.New(ctx, authnConfig)

	// logger
	logger := httplog.NewLogger(cfg.Name,
		httplog.Options{
			JSON:     cfg.LogFormatJSON,
			LogLevel: cfg.LogLevel,
		})

	index, err := rl.New(
		rl.Layout("index"),
		rl.DisableCache(true),
		rl.Debug(true),
		rl.DefaultData(defaultPageHandler(appCtx)),
	)

	if err != nil {
		log.Fatal(err)
	}

	// middlewares
	r := chi.NewRouter()
	r.Use(middleware.Compress(5))
	r.Use(middleware.Heartbeat(cfg.HealthPath))
	r.Use(middleware.Recoverer)
	r.Use(httplog.RequestLogger(logger))
	r.NotFound(index("404"))
	// public
	r.Route("/", func(r chi.Router) {
		r.Get("/", index("home", homePage()))
		r.Get("/overview", index("overview", overviewPage(appCtx)))
		r.Get("/export", index("dataexport", dataExportPage()))
		r.Get("/trial-dashboard", index("clinicaltrialdashboard", clinicalTrialsOverviewPage()))
		r.Get("/trial-new", index("newtrial", clinicalTrialsOverviewPage()))
		//r.Get("/trial", index("trial", clinicalTrialComparisonPage(appCtx)))
		r.Get("/pcv_antibodies", index("pcv", pcvPage()))
		r.Get("/signup", index("account/signup"))
		r.Post("/signup", index("account/signup", signupPage(appCtx)))
		r.Get("/confirm/{token}", index("account/confirmed", confirmEmailPage(appCtx)))
		r.Get("/login", index("account/login", loginPage(appCtx)))
		r.Post("/login", index("account/login", loginPageSubmit(appCtx)))
		r.Get("/auth/callback", index("account/login", loginProviderCallbackPage(appCtx)))
		r.Get("/auth", index("account/login", loginProviderPage(appCtx)))
		r.Get("/magic-link-sent", index("account/magic"))
		r.Get("/magic-login/{otp}", index("account/login", magicLinkLoginConfirm(appCtx)))
		r.Get("/forgot", index("account/forgot"))
		r.Post("/forgot", index("account/forgot", forgotPage(appCtx)))
		r.Get("/reset/{token}", index("account/reset"))
		r.Post("/reset/{token}", index("account/reset", resetPage(appCtx)))
		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			acc, err := appCtx.authn.CurrentAccount(r)
			if err != nil {
				log.Println("err logging out ", err)
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
			acc.Logout(w, r)
		})
		r.Get("/change/{token}", index("account/changed", confirmEmailChangePage(appCtx)))
	})

	// authenticated
	r.Route("/account", func(r chi.Router) {
		r.Use(appCtx.authn.IsAuthenticated)
		r.Get("/", index("account/main", accountPage(appCtx)))
		r.Post("/", index("account/main", accountPageSubmit(appCtx)))
		r.Post("/delete", index("account/main", deleteAccount(appCtx)))
	})

	r.Route("/app", func(r chi.Router) {
		r.Use(appCtx.authn.IsAuthenticated)
		r.Get("/", index("insert", insertClinicalTrial()))
		r.Get("/sp", index("supportingdata", getSupportingContent()))
	})

	r.Route("/view", func(r chi.Router) {
		r.Get("/", index("view", lookupTrialPage()))
	})

	r.Route("/dashboard", func(r chi.Router) {
		r.Route("/clinical-trial", func(r chi.Router) {
			r.Post("/", GetVsTrials(appCtx))
		})
		r.Route("/manufacturer", func(r chi.Router) {
			r.Get("/", GetDistinctManufacturers(appCtx))
		})
		r.Route("/schedule", func(r chi.Router) {
			r.Get("/", GetSchedule(appCtx))
		})
	})

	r.Route("/export-options", func(r chi.Router) {
		r.Route("/usecase", func(r chi.Router) {
			r.Get("/", GetUseCase(appCtx))
		})
		r.Route("/responsible-parties", func(r chi.Router) {
			r.Get("/", GetResponsibleParties(appCtx))
		})
		r.Route("/vaccine", func(r chi.Router) {
			r.Get("/", GetVaccine(appCtx))
		})
		r.Route("/icg", func(r chi.Router) {
			r.Get("/", GetICG(appCtx))
		})
		r.Route("/data-export", func(r chi.Router) {
			r.Post("/", PostCSVExport(appCtx))
			r.Get("/", GetCSVExport(appCtx))
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(appCtx.authn.IsAuthenticated)
		r.Use(middleware.AllowContentType("application/json"))
		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", list(appCtx))
			r.Post("/", create(appCtx))
		})
		r.Route("/tasks/{id}", func(r chi.Router) {
			r.Put("/status", updateStatus(appCtx))
			r.Put("/text", updateText(appCtx))
			r.Delete("/", delete(appCtx))
		})
		r.Route("/trials", func(r chi.Router) {
			r.Post("/", insertStudy(appCtx))

		})
		r.Route("/responsible-parties", func(r chi.Router) {
			r.Get("/", GetResponsibleParties(appCtx))
		})
		r.Route("/trials/{id}", func(r chi.Router) {
			r.Get("/", Get(appCtx))
			r.Get("/outcome-overview-list/", GetOutcomeOverviewList(appCtx))
			r.Get("/metadata/", GetMetadata(appCtx))
			r.Get("/outcome-measures/", GetOutcomeMeasures(appCtx))
			r.Route("/custom-outcome-measure", func(r chi.Router) {
				r.Post("/", InsertCustomOutcomeMeasure(appCtx))
				r.Delete("/{omid}", DeleteOutcomeMeasure(appCtx))
			})
		})
		r.Route("/data-export", func(r chi.Router) {
			r.Post("/", PostCSVExport(appCtx))
		})
		r.Route("/vaccine", func(r chi.Router) {
			r.Get("/", GetVaccine(appCtx))
			r.Post("/", CreateVaccine(appCtx))
			r.Put("/{id}", UpdateVaccine(appCtx))
			r.Delete("/{id}", DeleteVaccine(appCtx))
		})
		r.Route("/dose-description", func(r chi.Router) {
			r.Get("/", GetDoseDescription(appCtx))
			r.Post("/", CreateDoseDescription(appCtx))
			r.Put("/{id}", UpdateDoseDescription(appCtx))
			r.Delete("/{id}", DeleteDoseDescription(appCtx))
		})
		r.Route("/schedule", func(r chi.Router) {
			r.Get("/", GetSchedule(appCtx))
			r.Post("/", CreateSchedule(appCtx))
			r.Put("/{id}", UpdateSchedule(appCtx))
			r.Delete("/{id}", DeleteSchedule(appCtx))
		})
		r.Route("/usecase", func(r chi.Router) {
			r.Get("/", GetUseCase(appCtx))
			r.Post("/", CreateUseCase(appCtx))
			r.Put("/{id}", UpdateUseCase(appCtx))
			r.Delete("/{id}", DeleteUseCase(appCtx))
		})
		r.Route("/manufacturer", func(r chi.Router) {
			r.Get("/", GetManufacturer(appCtx))
			r.Post("/", CreateManufacturer(appCtx))
			r.Put("/{id}", UpdateManufacturer(appCtx))
			r.Delete("/{id}", DeleteManufacturer(appCtx))
		})
		r.Route("/icg", func(r chi.Router) {
			r.Get("/", GetICG(appCtx))
			r.Post("/", CreateICG(appCtx))
			r.Put("/{id}", UpdateICG(appCtx))
			r.Delete("/{id}", DeleteICG(appCtx))
		})
	})

	return r
}
