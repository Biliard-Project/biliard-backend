package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Biliard-Project/biliard-backend/controllers"
	"github.com/Biliard-Project/biliard-backend/migrations"
	"github.com/Biliard-Project/biliard-backend/models"
	"github.com/Biliard-Project/biliard-backend/templates"
	"github.com/Biliard-Project/biliard-backend/views"
	"github.com/joho/godotenv"

	"github.com/go-chi/chi/v5"
)

type config struct {
	PSQL models.PostgresConfig
	SMTP models.SMTPConfig
	CSRF struct {
		Key    string
		Secure bool
	}
	Server struct {
		Address string
	}
}

func loadEnvConfig() (config, error) {
	var cfg config
	err := godotenv.Load()
	if err != nil {
		return cfg, err
	}
	// TODO: PSQL
	cfg.PSQL = models.DefaultPostgresConfig()

	// TODO: SMTP
	cfg.SMTP.Host = os.Getenv("SMTP_HOST")
	cfg.SMTP.Port, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	cfg.SMTP.Username = os.Getenv("SMTP_USERNAME")
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")

	cfg.CSRF.Key = "zTRUrqhAFWSH0NR6SsGpFRQn7KqLEvvh"
	cfg.CSRF.Secure = false

	cfg.Server.Address = "0.0.0.0:3000"

	return cfg, nil
}

func webserver() {
	cfg, err := loadEnvConfig()
	if err != nil {
		panic(err)
	}

	// SET UP DATABASE
	db, err := models.Open(cfg.PSQL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// SETUP SERVICES
	userService := &models.UserService{
		DB: db,
	}
	sessionService := &models.SessionService{
		DB: db,
	}
	pwResetService := &models.PasswordResetService{
		DB: db,
	}
	patientService := &models.PatientService{
		DB: db,
	}
	recordService := &models.RecordService{
		DB: db,
	}
	emailService := models.NewEmailService(cfg.SMTP)

	// SETUP MIDDLEWARE
	umw := controllers.UserMiddleware{
		SessionService: sessionService,
	}

	// SETUP CONTROLLERS
	userC := controllers.Users{
		UserService:          userService,
		SessionService:       sessionService,
		PasswordResetService: pwResetService,
		EmailService:         emailService,
	}
	userC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))
	userC.Templates.SignIn = views.Must(views.ParseFS(templates.FS, "signin.html", "tailwind.html"))
	userC.Templates.ForgotPassword = views.Must(
		views.ParseFS(templates.FS, "forgot-pw.html", "tailwind.html"),
	)
	userC.Templates.CheckYourEmail = views.Must(
		views.ParseFS(templates.FS, "check-your-email.html", "tailwind.html"),
	)
	userC.Templates.ResetPassword = views.Must(
		views.ParseFS(templates.FS, "reset-pw.html", "tailwind.html"),
	)

	patientC := controllers.Patients{
		PatientService: patientService,
	}
	recordC := controllers.Records{
		RecordService: recordService,
	}

	// SETUP ROUTER AND ROUTES
	r := chi.NewRouter()
	r.Use(setCors)
	r.Use(umw.SetUser)
	r.Get(
		"/",
		controllers.StaticHandler(
			views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html")),
		),
	)
	r.Get(
		"/contact",
		controllers.StaticHandler(
			views.Must(views.ParseFS(templates.FS, "contact.html", "tailwind.html")),
		),
	)
	r.Get(
		"/faq",
		controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))),
	)
	r.Get("/signup", userC.New)
	r.Post("/signup", userC.Create)
	r.Get("/signin", userC.SignIn)
	r.Post("/signin", userC.ProcessSignIn)
	r.Post("/signout", userC.ProcessSignOut)
	r.Get("/forgot-pw", userC.ForgotPassword)
	r.Post("/forgot-pw", userC.ProcessForgotPassword)
	r.Get("/reset-pw", userC.ResetPassword)
	r.Post("/reset-pw", userC.ProcessResetPassword)
	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", userC.CurrentUser)
	})
	r.Route("/patients", func(r chi.Router) {
		r.Get("/", patientC.ProcessGetPatients)
		r.Get("/{patientID}", patientC.ProcessGetPatientByID)
		r.Post("/", patientC.Create)
	})
	r.Route("/records", func(r chi.Router) {
		r.Get("/", recordC.GetAllPatientRecords)
		r.Get("/patient/{patientID}", recordC.GetRecordsByPatientID)
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// START THE SERVER
	fmt.Printf("starting server at %s...\n", cfg.Server.Address)
	err = http.ListenAndServe(cfg.Server.Address, r)
	if err != nil {
		panic(err)
	}
}

func main() {
	go webserver()
	select {}
}

func setCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
