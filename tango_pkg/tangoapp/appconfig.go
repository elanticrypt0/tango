package tangoapp

type AppConfig struct {
	Name             string `json:"name" toml:"name"`
	Version          string `json:"version" toml:"version"`
	Author           string `json:"author" toml:"author"`
	Contact          string `json:"contact" toml:"contact"`
	Repo             string `json:"repo" toml:"repo"`
	ServerProtocol   string `json:"server_protocol" toml:"server_protocol"`
	ServerHost       string `json:"server_host" toml:"server_host"`
	ServerPort       uint16 `json:"server_port" toml:"server_port"`
	Url              string `json:"url" toml:"url"`
	PublicPath       string `json:"public_path" toml:"public_path"`
	PublicAssetsPath string `json:"public_assets_path" toml:"public_assets_path"`
	UseAstroFrontend bool   `json:"use_astro_frontend" toml:"use_astro_frontend"`
	UseTempl         bool   `json:"use_templ" toml:"use_templ"`
	SetupEnabled     bool   `json:"setup_enabled" toml:"setup_enabled"`
	NotInProduction  bool   `json:"not_in_production" toml:"not_in_production"`
	OpenInBrowser    bool   `json:"open_in_browser" toml:"open_in_browser"`
	CORSOrigins      string `json:"cors_origins" toml:"cors_origins"`
	CORSHeaders      string `json:"cors_headers" toml:"cors_headers"`
}
