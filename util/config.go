package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type database struct {
	Password string
	Usuario  string
}

//Config é a estrutura que representa toda a configuração da API
type Config struct {
	Database database
}

//Configuration variável global das configurações da api
var configuration *Config

//LoadConfig método responsável por carregar as variáveis de ambiente
func LoadConfig() error {

	err := godotenv.Load(".env-development")
	if err != nil {
		fmt.Println("[Método: Config.LoadConfig()] Arquivo não encontrado: .env-development. Favor configure o arquivo antes de iniciar a API")
	}
	fmt.Println("Arquivo .env-development carregado com sucesso")

	configuration = &Config{
		Database: database{
			Usuario:  os.Getenv("USUARIO"),
			Password: os.Getenv("PASSWORD"),
		},
	}
	return nil
}

//Get returns a Config Structure
func Get() *Config {
	return configuration
}
