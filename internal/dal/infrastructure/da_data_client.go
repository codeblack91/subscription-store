package infrastructure

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CurrencyClient interface {
	FindCurrency(query string) (string, error)
}

type DaDataClient struct {
	Token string
	URL   string
}

func NewDaDataClient(token, url string) *DaDataClient {
	return &DaDataClient{
		Token: token,
		URL:   url,
	}
}

// FindCurrency — метод для запроса данных о валюте
func (c *DaDataClient) FindCurrency(query string) (string, error) {
	// Формируем тело запроса
	body := map[string]string{
		"query": query,
	}

	// Преобразуем тело запроса в JSON
	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("error marshalling json: %v", err)
	}

	// Создаем новый POST запрос
	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Token "+c.Token)

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Возвращаем тело ответа
	return string(bodyResp), nil
}
