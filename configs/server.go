package configs

import "github.com/spf13/viper"

func SetServerConfigs() {
	viper.SetDefault("NAME", "dcard-2024-backend-intern-assignment")
	viper.SetDefault("VERSION", "1.0.0")
	viper.SetDefault("PORT", "5000")
}
