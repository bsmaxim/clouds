package main

import (
	"fmt"
	"net"
)

func closeListener(ln net.Listener) {
    err := ln.Close()
    if err != nil {
        fmt.Print("Ошибка закрытия сервера: ", err.Error())
        return
    }
}

func handleTCPConnection(conn net.Conn) {
    defer conn.Close()

    // отправка сообщения клиенту
    message := "Hi client, I'm server\r\n"
    n, err := conn.Write([]byte(message));
    if err != nil {
        fmt.Println("Ошибка записи", err.Error())
        return
    }
    fmt.Println("Клиенту передано: ", n, " символов")

    buf := make([]byte, 1024)
    n, err = conn.Read(buf)
    if err != nil {
        fmt.Println("Ошибка чтения:", err.Error())
        return
    }
    fmt.Println("Получено от клиента:", string(buf[:n]))
}

func main() {
    ln, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Ошибка запуска сервера: ", err.Error())
        return
    }

    defer closeListener(ln)

    fmt.Println("TCP сервер слушает адрес localhost:8080")

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Ошибка подключения к клиенту: ", err.Error())
            return
        }
        go handleTCPConnection(conn)
    }
}
