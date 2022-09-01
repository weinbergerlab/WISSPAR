package app

import (
	"github.com/kelseyhightower/envconfig"
	"os"
	"strconv"
)

type Config struct {
	Name             string `json:"name" default:"Clinical Trial"`
	Domain           string `json:"domain" envconfig:"url" default:"http://localhost"`
	Port             int    `json:"port" default:"80"`
	HealthPath       string `json:"health_path" envconfig:"health_path" default:"/healthz"`
	ReadTimeoutSecs  int    `json:"read_timeout_secs" envconfig:"read_timeout_secs" default:"5"`
	WriteTimeoutSecs int    `json:"write_timeout_secs" envconfig:"write_timeout_secs" default:"10"`
	LogLevel         string `json:"log_level" envconfig:"log_level" default:"error"`
	LogFormatJSON    bool   `json:"log_format_json" envconfig:"log_format_json" default:"false"`
	Templates        string `json:"templates" envconfig:"templates" default:"templates"`
	SessionSecret    string `json:"session_secret" envconfig:"session_secret" default:"mysessionsecret"`
	APIMasterSecret  string `json:"api_master_secret" envconfig:"api_master_secret" default:"supersecretkeyyoushouldnotcommit"`

	// datasource
	Driver     string `json:"driver" envconfig:"driver" default:"sqlite3"`
	DataSource string `json:"datasource" envconfig:"datasource" default:"file:gomodest.db?mode=memory&cache=shared&_fk=1"`

	// smtp
	SMTPHost       string `json:"smtp_host" envconfig:"smtp_host" default:"mailhog"`
	SMTPPort       int    `json:"smtp_port,omitempty" envconfig:"smtp_port" default:"1025"`
	SMTPUser       string `json:"smtp_user" envconfig:"smtp_user" default:"myuser"`
	SMTPPass       string `json:"smtp_pass,omitempty" envconfig:"smtp_pass" default:"mypass"`
	SMTPAdminEmail string `json:"smtp_admin_email" envconfig:"smtp_admin_email" default:"noreply@gomodest.xyz"`
	SMTPDebug      bool   `json:"smtp_debug" envconfig:"smtp_debug" default:"true"`

	// goth
	GoogleClientID string `json:"google_client_id" envconfig:"google_client_id"`
	GoogleSecret   string `json:"google_secret" envconfig:"google_secret"`

	// Metabase information
	MetabaseURL string `json:"metabase_url" envconfig:"metabase_url"`
	MetabaseSecret   string `json:"metabase_secret" envconfig:"metabase_secret"`

	// OTP Secret
	OTPSecret   string `json:"otp_secret" envconfig:"otp_secret"`
}

func LoadConfig(envPrefix string) (Config, error) {
	var config Config
	if err := envconfig.Process(envPrefix, &config); err != nil {
		return config, err
	}

	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		return config, nil
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return config, err
	}
	config.Port = port

	return config, nil
}


