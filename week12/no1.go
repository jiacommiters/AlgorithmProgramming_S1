package main

import "fmt"

func main() {
    var username, password string
    var gagalLogin int = 0
    var loginBerhasil bool = false

    for !loginBerhasil {
        fmt.Scan(&username, &password)

        if username == "admin" && password == "admin" {
            loginBerhasil = true
        } else {
            gagalLogin = gagalLogin + 1
        }
    }	

    fmt.Println(gagalLogin)
    fmt.Println("Login berhasil")
}