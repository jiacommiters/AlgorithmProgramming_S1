package main

import "fmt"

func main() {

	var user, password string
	fmt.Scan(&user, &password)

	if user == "admin" && password == "12345" {

		fmt.Println("Login Berhasil")

	} else if user != "admin" && password != "12345" {

		fmt.Println("Username Tidak Ditemukan")
		fmt.Println("Password Salah")

	} else if user != "admin" {

		fmt.Println("Username Tidak Diptemukan")

	} else if password != "12345" {

		fmt.Println("Password Salah")

	} else {

	}

}
