package entity


type Config struct {
	ServerPort					string	`mapstructure:"SERVER_PORT"`

	DbRootPassword				string	`mapstructure:"DB_ROOT_PASSWORD"`
	DbName						string	`mapstructure:"DB_NAME"`
	DbUser						string	`mapstructure:"DB_USER"`
	DbPassword					string	`mapstructure:"DB_PASSWORD"`
	DbHost						string	`mapstructure:"DB_HOST"`

	AccessTokenHourLifespan		int		`mapstructure:"ACCESS_TOKEN_HOUR_LIFESPAN"`
	RefreshTokenHourLifespan	int		`mapstructure:"REFRESH_TOKEN_HOUR_LIFESPAN"`
	AccessTokenSecret			string	`mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret			string	`mapstructure:"REFRESH_TOKEN_SECRET"`
	TokenPrefix					string	`mapstructure:"TOKEN_PREFIX"`

}