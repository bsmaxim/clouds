// задание: «Доска объявлений». Если клиент передает команду «LIST», то ему выводятся все объявления (текстовые строки), добавленные пользователями на доску объявлений через «;». Если клиент передает другую текстовую строку S, она сохраняется во внешнем файле (или базе данных) и пользователю возвращается подтверждение формата: Message added: “S”. Если клиент вводит пустую строку, то соединение разрывается.

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ln, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Printf("server started on: %v\n", ln.Addr())

	createDb("database")
	initCollection("items")

	bs := newBulletinServer()
	s := &http.Server{
		Handler: bs,
	}

	errc := make(chan error, 1)
	go func() {
		s.Serve(ln)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		log.Printf("не удалось обработать: %v", err)
	case sig := <-sigs:
		log.Printf("завершение: %v", sig)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return s.Shutdown(ctx) // закрывает сервер после закрытия всех соединений или отмены контекста
}
