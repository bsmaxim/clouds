package main

import (
	"fmt"
	"net"
	"os"
)

func closeConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println("Ошибка закрытия:", err.Error())
		return
	}
	fmt.Println("Соединение закрыто (клиент)")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Введите номер клиента в качестве аргумента командной строки")
		return
	}

	var client_id string = os.Args[1]

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err.Error())
		return
	}
	defer conn.Close()

	// чтение
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Ошибка чтения:", err.Error())
		return
	}
	fmt.Println("С сервера получено сообщение:", string(buf[:n]))

	// Отправка данных серверу
	n, err = conn.Write([]byte("Hi, I am a client number " + client_id + "\r\n"))
	if err != nil {
		fmt.Println("Ошибка записи:", err.Error())
		return
	}
	fmt.Println("Серверу передано:", n, "символов")
}
