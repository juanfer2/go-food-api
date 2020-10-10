package constants

import "github.com/juanfer2/api-rest-go/utils"

func ApiUrl() string {
	return utils.GetEnvVariable("API_URL")
}

func Port() string {
	return utils.GetEnvVariable("PORT")
}

var SecretKey = "jdnfksdmfksd"
