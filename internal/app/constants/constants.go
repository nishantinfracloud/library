package constants

const (
	// Service Name
	ServiceName = "library"

	//Viper Keys
	AllowedOrigins = "AllowedOrigins"
	TrustedProxies = "TrustedProxies"
	Environment    = "Environment"

	// URL contexts
	Library    = "/library"
	API        = "/api"
	Version_V1 = "/v1"
	FetchBooks = "/books"
	AddBook    = "/addBook"

	//Config File Path
	ConfigFilePath = "./internal/app"

	//Error Message
	InvalidUrlShortenerNameError = "Invalid UrlShortener name provided"

	// Environment
	DevEnvironment = "dev"

	LibraryBaseUrl = "LibraryBaseUrl"
)
