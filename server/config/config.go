package config

var Config *ServerConfig

type ServerConfig struct {
	Addr                        string
	Licensee                    string
	UserCertPath                string
	UserPrivateKeyPath          string
	LicenseServerCertPath       string
	LicenseServerPrivateKeyPath string
}

func InitServerConfig(c *ServerConfig) {
	Config = c
}
