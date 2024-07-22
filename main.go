package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func getHref(t html.Token)(ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return
}

func crawl(url string, ch chan string, chFinished chan bool) {
	resp, err := http.Get(url) // получить ответ

	// функция с defer вызывается в конце этой функции, отложенная функция
	defer func() {
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("Не удалось выполнить сканирование", url)
		return
	}

	b := resp.Body // получить тело ответа
	defer b.Close() // закрыть тело ответа после выполнения функции

	z := html.NewTokenizer(b) // токенизатор html

	for {
		tt := z.Next() // получить следующий токен

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

		isAnchor := t.Data == "a"

		if !isAnchor {
			continue
		}

		ok, url := getHref(t)

		if !ok {
			continue
		}

		hasProto := strings.Index(url, "http") == 0 // проверка на протокол

		if hasProto {
			ch <- url // публикация на сайте уникальных url-адресов
		}
	}
	}
}

func main() {
	foundUrls := make(map[string]bool) // найденные url-адреса
	seedUrls := os.Args[1:] // Получение аргументов командной строки

	chUrls := make(chan string) // канал url
	chFinished := make(chan bool) // канал завершения работы

	// разбор всех url адресов
	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinished)
	}

	for c := 0; c < len(seedUrls); {
		select {
		case url := <- chUrls:
			foundUrls[url] = true
		case <- chFinished:
			c++
		}
	}

	fmt.Println("\nFound", len(foundUrls), "unique links:\n")

	// распечатка всех адресов
	for url, _ := range foundUrls {
		fmt.Println("- " + url)
	}

	// закрыть канал
	close(chUrls)
}