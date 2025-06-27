package config

import (
	"log" // Pour logger les informations ou erreurs de chargement de config

	"github.com/spf13/viper" // La bibliothèque pour la gestion de configuration
)

// Config est la structure principale qui mappe l'intégralité de la configuration de l'application.
// Les tags `mapstructure` sont utilisés par Viper pour mapper les clés du fichier de config
// (ou des variables d'environnement) aux champs de la structure Go.
type Config struct {
	Server    Server    `mapstructure:"server"`
	Database  Database  `mapstructure:"database"`
	Analytics Analytics `mapstructure:"analytics"`
	Monitor   Monitor   `mapstructure:"monitor"`
}

// ServerConfig contient la configuration du serveur web
type Server struct {
	Port    int    `mapstructure:"port"`
	BaseURL string `mapstructure:"base_url"`
}

// DatabaseConfig contient la configuration de la base de données
type Database struct {
	Name string `mapstructure:"name"`
}

// AnalyticsConfig contient la configuration des analytics asynchrones
type Analytics struct {
	BufferSize  int `mapstructure:"buffer_size"`
	WorkerCount int `mapstructure:"worker_count"`
}

// MonitorConfig contient la configuration du moniteur d'URLs
type Monitor struct {
	IntervalMinutes int `mapstructure:"interval_minutes"`
}

// LoadConfig charge la configuration de l'application en utilisant Viper.
// Elle recherche un fichier 'config.yaml' dans le dossier 'configs/'.
// Elle définit également des valeurs par défaut si le fichier de config est absent ou incomplet.
func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Définir les valeurs par défaut pour toutes les options de configuration.
	// Ces valeurs seront utilisées si les clés correspondantes ne sont pas trouvées dans le fichier de config
	// ou si le fichier n'existe pas.

	// Valeurs par défaut pour le serveur
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.base_url", "http://localhost:8080")

	// Valeurs par défaut pour la base de données
	viper.SetDefault("database.name", "url_shortener.db")

	// Valeurs par défaut pour les analytics
	viper.SetDefault("analytics.buffer_size", 1000)
	viper.SetDefault("analytics.worker_count", 5)

	// Valeurs par défaut pour le moniteur
	viper.SetDefault("monitor.interval_minutes", 5)

	// Lire le fichier de configuration.
	if err := viper.ReadInConfig(); err != nil {
		// Si le fichier n'existe pas, on utilise les valeurs par défaut
		log.Printf("Config file not found, using default values: %v", err)
	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}

	// Démapper (unmarshal) la configuration lue (ou les valeurs par défaut) dans la structure Config.
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	log.Printf("Configuration loaded: Server Port=%d, DB Name=%s, Analytics Buffer=%d, Monitor Interval=%dmin",
		cfg.Server.Port, cfg.Database.Name, cfg.Analytics.BufferSize, cfg.Monitor.IntervalMinutes)

	return &cfg, nil
}
