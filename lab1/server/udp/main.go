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

func handleUDPConnection(conn *net.UDPConn) {
    defer conn.Close()

    buf := make([]byte, 1024)
    n, addr, err := conn.ReadFromUDP(buf)
    if err != nil {
        fmt.Println("Ошибка чтения:", err.Error())
        return
    }
    fmt.Println("Получено от клиента:", string(buf[:n]))

    // отправка сообщения клиенту
    message := "Hi client, I'm server\r\n"
    n, err = conn.WriteToUDP([]byte(message), addr)
    if err != nil {
        fmt.Println("Ошибка записи", err.Error())
        return
    }
    fmt.Println("Клиенту передано: ", n, " символов")
}

func main() {
    addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
    if err != nil {
        fmt.Println("Ошибка разрешения адреса: ", err.Error())
        return
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("Ошибка запуска сервера: ", err.Error())
        return
    }

    defer conn.Close()

    fmt.Println("UDP сервер слушает адрес localhost:8080")

    for {
        var buf [1024]byte
        _, addr, err := conn.ReadFromUDP(buf[0:])
        if err != nil {
            fmt.Println("Ошибка чтения:", err.Error())
            continue
        }

        fmt.Println("Получено от клиента:", string(buf[:]))

        conn.WriteToUDP([]byte("Hi client, I'm server\r\n"), addr)
    }
}
