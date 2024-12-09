package main

import (
	"fmt"
	"log"
	"net/http"
	"subscription-store/internal/api/routes"
	"subscription-store/internal/dal/infrastructure"
	"subscription-store/internal/service"

	//log "subscription-store/internal/utils"

	"github.com/spf13/viper"
)

func init() {
	//сделать init конфиг
	viper.SetConfigName("config")                                      // Имя файла без расширения
	viper.SetConfigType("yaml")                                        // Тип файла
	viper.AddConfigPath("/Users/v/Projects/subscription-store/config") // Каталог, где искать файл

	// Чтение файла конфигурации
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file: %w", err))
	}
}

func TestDaData() {
	token := viper.Get("daData.token")
	url := viper.Get("daData.url")

	// Инициализация клиента
	dadataClient := infrastructure.NewDaDataClient(token.(string), url.(string))

	// Инициализация бизнес-логики
	currencyUseCase := service.NewCurrency(dadataClient)

	// Выполнение запроса
	currencyInfo, err := currencyUseCase.GetCurrencyInfo("RUB")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Вывод результата
	fmt.Println("Currency Info:", currencyInfo)
}

func main() {

	TestDaData() //тестовый вызов http клиента и вывод на консоль

	//logger := log.

	r := routes.NewRouter()
	fmt.Println("Сервер запущен на порте :8080")
	log.Fatal(http.ListenAndServe(viper.Get("server.port").(string), r))
}
