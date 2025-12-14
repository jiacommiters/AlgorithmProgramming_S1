package main
import "fmt"

func main() {
    var day string
    var date, month, year, daysInMonth, waitingDays int

    fmt.Scan(&day, &date, &month, &year)

    for day != "Exit" {
        waitingDays = 2
        if month == 4 || month == 6 || month == 9 || month == 11 {
            daysInMonth = 30
        } else if month == 2 {
            daysInMonth = 28
        } else {
            daysInMonth = 31
        }
        if (year % 4 == 0 || year % 400 == 0) && year % 100 != 0 && month == 2 {
            daysInMonth = 29
        }
        if day == "Kamis" || day == "Jumat" {
            waitingDays = 4
        }
        if (date + waitingDays) > daysInMonth {
            month = month + 1
            date = (date + waitingDays) - daysInMonth
            if month > 12 {
                year = year + 1
                month = month - 12
            }
        } else {
            date = date + waitingDays
        }
        fmt.Printf("Passpor bisa diambil pada tanggal %d bulan %d tahun %d \n", date, month, year)
        fmt.Scan(&day, &date, &month, &year)
    }
}