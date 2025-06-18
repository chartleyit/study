package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// const apiKey = "2eec12665fdb482d492816d90b5a4220"
const apiKey = "4f9792bc0f55a12f74d2638a94b4473e"

func KtoF(kelvin float64) float64 {
	return (9/5)*(kelvin-273) + 32
}

func KtoC(kelvin float64) float64 {
	return kelvin - 273
}

func fetchWeather(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("\nError fetching weather for %s: %s\n", city, err)
		return data
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("\nError decoding weather data for %s: %s\n", city, err)
		return data
	}

	ch <- fmt.Sprintf("This is the %s, %.2f", city, KtoF(data.Main.Temp))

	return data
}

func main() {
	startNow := time.Now()
	cities := []string{"Miami", "Osaka", "London"}
	ch := make(chan string)

	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This operation took:", time.Since(startNow))
}
