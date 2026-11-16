package main

import "fmt"

func main() {
    var karakter string
    fmt.Scan(&karakter)

    if len(karakter) == 1 &&
        (karakter[0] == 'a' || karakter[0] == 'i' || karakter[0] == 'u' ||
            karakter[0] == 'e' || karakter[0] == 'o' ||
            karakter[0] == 'A' || karakter[0] == 'I' || karakter[0] == 'U' ||
            karakter[0] == 'E' || karakter[0] == 'O') {
        fmt.Println("Vokal")
    } else if len(karakter) == 1 &&
        ((karakter[0] >= 'a' && karakter[0] <= 'z') ||
            (karakter[0] >= 'A' && karakter[0] <= 'Z')) {
        fmt.Println("Konsonan")
    } else {
        fmt.Println("Bukan huruf")
    }
}