package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.1" )

func foRlAoJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDlQTyX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MtRTfI8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpzJpqpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KE4LXEZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McsFrlWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxM4iZJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHH99sFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBVRmxCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCClPWNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9X0CCruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z1bi9NBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qvUwKcQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gmfNFVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHUQrQ5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNCy4EkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leeMaNTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Opp3cXpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bl0VjUXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUDsJaHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wD2AqUjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lejZ7ig6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GycLac8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOyUfT5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfc01j9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A04ouZOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGdMgxl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIOwyjG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XidebAV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOqK0JdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRCYRPBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zRTgRMWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aE38ljweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdLjhZDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ocjIW3yPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iQJAreEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W59n55GkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1hTdbOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHJYl3eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vy1P4MeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqAxPalPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f2n0N3hDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IEb9UaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q26RWm5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gcJ6dY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsNpLRazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xb9q1O1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H52ugWg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUvoQBxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1E9AHKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5Hq3xPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hS2puUJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FmnCXcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ubCzVSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZeY7ZrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxw7JjIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwSwfUWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TBttGxFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzRIjqnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nL5iW2HCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func claZpdvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igM23hfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mC9WzSgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiHQjwQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWpDtJ8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DHIEWjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPugZFVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qeC5oTCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qPbz4YqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uClCVjHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urSJnuaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bowSXDSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IV0lrOjHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajucdO6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d3osMSQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eNPcnMMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func veSLDqsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbgilblZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5YVADYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMG2cP7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G0prT3hbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kU7Lp84WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dY6YhTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRKiK5rjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbCaBXvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndy7r7LyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Er9yPb8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMmtHHvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrRqbXweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21VSenPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhCaj6seWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILhaOc0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Txv1AxomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsGjqUp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXeMWcRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSkTTCDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aw8PTW9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hl8T5nhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcGRfbD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cWiXk1BYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YqPKuciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnwo0l7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUZOSnh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKfVUpw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Atklx2SUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYMSxR5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qpx4mWdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJt3gBk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUdwo13IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7dTlpZKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HawKDhjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mujE84MPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tl4VIObkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kjGkDjeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mubTslHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHteI55fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3UMwFySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5mhCij0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qP7ileqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJpQ4fiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nNZcN7ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6WMfx7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOQBUbfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K27zcs0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQfuV1WyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcl0XJiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYK1TTqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgnmRMnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EseHSfIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5xxwrERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVlPVg3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3Q1gKiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgzNOJXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0hgKDyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOlmwDq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPP3bELdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpNM9PF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3XsJKn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZMbjISXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wefopAO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjSml1tGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2a54KhgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func feWdGBVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ourxBkEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zdUIxwlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URtnTs3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUqGmcwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xv1pZzwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDljlgSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LD6lEgAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sSUwSyT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Puy7w9pyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mnwmvPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WP5RJcLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRbc4I9ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bSAZp8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TV7NpEzdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvElwRobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsELy5n6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fh1PvjU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func py28eUIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zj0OVJ05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3uKrfqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcxBGO5UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1vhzYVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r1iTkBtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O46c5PC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOKj5yOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRQ6VHEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kPptSUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9e1XxyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0m99e9xkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aa5MjxINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjdbAWKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Coyc5D47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgUZTzlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POi8K2aeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUsh4lNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hz9zQvCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PxadDoj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UVvcbPJ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7m42qH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZo4ENlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YVoJNIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ng3qnEb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ko2it7tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7tv9MyErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czmq8O3NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSKw00boWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UM9gFfu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95fRGvTUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ob5cD8TFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSNpIdk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHxlsOBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQTk4zN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K45TqlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7tIKcQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6IPAwr7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bbzSCYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIAgjYmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmPMqWLWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6L56EHzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pudB6TmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6TTyR2eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WhPHfgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9F1yPFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xe0bSjkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECm6whCdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mi2qAd3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7U6XXgUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CuU8dQtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpuUUwvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCH7KByJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1l6srHltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eh3HK2FBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryM2FMhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KT3T74BjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tFnQBnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8DjbOCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHqQvccOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HF78en8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tytCZLA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfaPbfo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IolHZmW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gBYXgviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmCvD3AiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func deSxD9iMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JyJ4JTEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCc7wZUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGoP7cl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iAfU17lqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8kEVK0LCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WU3azgBLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QfUzZEj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmR4lfYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khHf5ZmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jb8tu0FMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoEBbzBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzgjzwMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTBTUgLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3fLFwf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EBDJzbOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqeQBHVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkBxpoBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPQOy03mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEB5CaKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWsIU6jXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3F7BxSxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDOFXXr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PQjXypEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pvp7JuZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vY9dUEckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 391OvP1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGQcbPjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rL4Q6w1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzTkKIQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8qn3ay7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gCItPKMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWQaoKvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l3exxvcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dm8EhV2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N43shU8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NkQZQDWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wacIa9uTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJJLJR17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiR1HfqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZGt3PbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0QtZ6iPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Lck2WnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkG6l63ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kXJgr1L5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 248Sze8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xObt9QqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWYluR0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFP8aGOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U93Hyhd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KabrzuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcL4GcVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMpT18o6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APKbnrZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoKHvZMvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJu9hYghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCpzN1lzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJ05gkWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52q2q2AvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X13mdpAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjFwHvMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wi5zB45PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooFzsvLJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZ9XqAn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2QdIWWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kT0CFEAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2BkcpzzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GNfakpYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8BMooVhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVQxKYpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func droqzMk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpML0q7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tOP2WIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjIGns2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkzMSpShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQXBhxlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NG4IbvXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFEtMLpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBMtNKBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPoFvPwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVGLvclUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWhTh67JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgKq28mpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71PAjhJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmuvkZAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJdG4dYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xrMlSjG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuIQOicBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ejJRjUSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMH74RjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACx7AYbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8PTzZvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T0bKzMEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4yRvcqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtiQ0alQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KdOWh8cFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kucCUXfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buryr4I3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7E02dndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAo1lTOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mKP7U9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FH4BCSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bR5r6GbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIr54arWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQ70ZJj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAiUU0GdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddfol2D1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26IPTpuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGKAeMytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjcBSEYwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftmNaAZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WLIF0W4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Qh4H15vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzohrlhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpTHhLlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hMtnYm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7DGVwiQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeNFny7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhssb2iwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5dwVnbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNMgSewbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBQBZspNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjPS9IaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBPSGcAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wswdrjFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QsoJv3eiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func af824NjWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQEvqayAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J99jIZDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4c5KadjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQR6sluwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9SYQmjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZ11DTp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5IAKnklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 64fIA48SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Am7cfHZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxSIwRq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KwcU192Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EdiQXrP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwTRH1YbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTzVSwAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwnhB2QfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xobOpY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utMi4svyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HyrMPGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zdxJ1Q4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMxQ3h5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TI21b5WiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jA3XuRetWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YthUx2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stexnPT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhGk1CDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rj6hOEfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEQxufLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v651fYQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8zFSVRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gQ6elAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OolwFp1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R37lwZVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrAIggdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oazCfx6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OR0E23kBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3fmgwtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgMdVMs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9L5RaqvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UJmuChsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyS9cdxHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ECptOoPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bk9EyrL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmOjBtr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EyrjGP71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBO9nPZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvN6vSQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MM7nwLjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSjz9cc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcSJOGYCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRhisKk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tMbQMvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZFV7LWFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yz4uGqD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJIf4aVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWfu9f1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6QE4QteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wk5kyOyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mh3UELnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GouSQXNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JzfDe9uNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1x11xigGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmQDwipgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2PyacHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ivEFczZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujuI5hUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHTYTtW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7OfPFnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3Ngy8r1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCdvQ8AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XWaVS1rRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rn7yKyuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVk4lJA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDJV8oxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMM1uXqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fy3qJqwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 08Lix4zEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4pwgQLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3LJUz9S3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEtiylstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c24gimmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZP6vtNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6rqnN3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hg2KgqGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4G7F3rQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s55CV8FQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0WrZBZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Emw0mD40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gPjE5Vu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7J38P3pWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfIn2un7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmr1kHaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQfvGvPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o780r5X3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilDW1z7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMU5nzY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8VEPYkZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H13j2Y8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvvsxoOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjSaY5h2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzbFI32FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grQTNeMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 91m61zPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbBFsNd0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R8pUt9XSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLWAJTkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9kX5qhkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWTVkSTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAinCWbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMMe0KLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hhgkP8VTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvEjOwJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXlbqwv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zFzDvDj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3BNwD83jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbxjlkdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQBQbJJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3SAWOV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaI31rpfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQpG6MdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rw7bqamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1NJE5duWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WycZF5ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJrAdOTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0WsvcqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fjUN7GtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndMqHhjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYzcmqe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvaRrn6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func txEQaIEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYcUMcssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGTsF7I6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qv6z7VHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgDjSUe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0CPpT7OzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dXVCkBphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDugLhHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lB6hD3K1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehCWfzrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvSkMf7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcCzm65bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDe0FdEyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9IPxiYYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J280gXvzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQSSMICsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ff2vZyUkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tV7FI0uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eyVZp9ixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3OTFiNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LI4eZannWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwVodieiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQCMA7CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDbEIhVLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aqogqjENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGB8tfD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXFzgEr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThguG8mTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDU903BKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vflgVueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMYUiSqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRcSBkH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueJR2wXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uvwFKcGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEJL8duoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RbwNDj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1rmWos0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HwOpiTIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFxYBz8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTUlKhRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h2eu1sIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhYSHfsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEFEseMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AE1afbPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TK824ymZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7Fzz7btWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpHKFrSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2F0Z3lbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWHlWcRDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BK06TeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pvWvzVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9hwAcenWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNcNbXR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9z8wv6UaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hwtrh3pZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozaGvQF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5DhltYLdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOPjAZL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TUHEtK0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCAJ1eiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NprJRwtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJSW8A6sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJAfkvBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9TLwy7YrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jXDPo3umWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMMjAgrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9afaLNMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3WsS8cYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N952CmVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VeP1aVgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipztq2WOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AygYVEcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yr5xMIpXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivJCx1ubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVcTdwu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEriRv38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRWK4U7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bh7EGEvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7A2lskJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9tGd0L5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fB5YwMHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dSs0eyqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRYLUZkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6J4tvaiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eF4bZUIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BK4zfNCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQncMsBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func muX1CNcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oaNyIPWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H65DOjSvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIBGHIR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFqcGgJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kb6FEiAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xi9LKNn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XxZoGIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElNVnhauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YAwbYvvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAhN8i02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GogyX7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZs9sGynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFN9R1hnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvXI9QUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66Qt7f9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSuHY22dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgr55DcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgB48tCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQBFHwVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o5zvUh7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgeaXP1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3jevXVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psyMYc51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pVDm4cSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggLWhbsVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOsfILdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5KR37RxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFMpd3JiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2iKOCyZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5PQnTxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJlgyTECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8coaAIuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkx4ugmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a31NULPbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnhZXQ1QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arIHXW4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WHXjnDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXVOvCLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7oKK5onLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLY1hsDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8AwXPztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfczN4NtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acsYSkLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFxxpFPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ovGkSYm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWL2b7gWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XlVLzUgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yWZ8XzPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lO224u1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySFdUJ8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7c9CpLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQVrlECoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V17yXoZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9e1xObKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0zJBGnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soHrs1DnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3aOZvXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZjSLXWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6psd0VVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5BfBX0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RRn1YnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D77HwkQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIOu2iHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MX6V4NNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1p3k1b1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSKPE7P3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kpP9ytDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLEJ2c6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3LiyY5dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ePclcranWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fA0MhrRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F0XyZutaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qx7AdToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNiK7hoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjRPS4cMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJiZxbvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8I3vnN5OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRdsEFppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvum2TDVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aFmGjfWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Tli304XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZEyA9hkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJJy9Js4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80w7lIVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyr4lOuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsf0pFg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlHJeBTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyFiSSTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c90ufmmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLjmBgtyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4OPWRQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HmQ75PXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cI62kDGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdjgKoImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func civK5CUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nptnFqoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MumJxwFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J6zpEavTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJTtkVIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DS0RVGxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D8hICJzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcGA9PlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bCOBkIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQ3TQPbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jl9wcax1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3B2gAvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYPfu0qeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gbjoe5JGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtJIeFTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89U1siDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fccRjxGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwP3Kc5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pM2f9NRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18Imdz3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wq19fv4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBVjVeYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLOP92KgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtZ0JlKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryisJn2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3Z1VD5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOGLLmfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bp4lo3uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jh64WhJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgLmFNR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yppo6EmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fy7mwJ4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zUPehzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyxgG9hgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKkwmKeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KLXasqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKAJZ3GhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9yXk6hbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6ZGfmpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQQPgPNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUfa2HoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BI0uTYGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkSlzqk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZ8Ftz50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kVpOnE4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkacieM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vDV9UrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2mLz2nAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGYyvD8hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEYHYbMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ViP4BimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8uJ94tXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dqlx6tFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HaffgRA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 77FNTl4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaWIqrSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQxeqIDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3zYM2XqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEHvxTgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMXEFBItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSL3i3gWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K145zlThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQCHatU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2itrsNZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjAINTO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qj2qbqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wygAnguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldhUPYJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fHq2O00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ypAI9MsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ah5REAPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3J3D0RoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7AMbsrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uySgHSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UdQH9fTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bn4K7b6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PID74OhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c29nTrUeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26YkQHg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5hBbZMaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caXsdM4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ammqV81ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPO221vaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLEQnkpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QKfitQfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ju3JmfhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uxtx5Tj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHKvxyjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ItvPbwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGwx7Cc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dUQbjI24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgsiToz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GQK8GyLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func COdBfiKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0jaIxv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLoHs6NRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdXs3BKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6lBJjRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDrpKUdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCdTBSVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOXT03y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXcwygY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJd8oCEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func StxnZ26dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7AyFomhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJgFXs15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpC0NpP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dg3SPNjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c1PkFqtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7OZlt1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9oeNffJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W28l2z00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwpZMrfnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YU5JRaj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func saR0FnlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IL4EcJvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bt7NTqsKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BmjI2bJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPMCLaY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ocTH18XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wKFKigrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sf7Ail9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wzWVWxOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yhPslrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3n31MhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CT2ZeIhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJApAVkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmV5mSXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYnZ9d1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvYVegxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMQ54LpmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbGCB2QkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mpNzYFCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y55GqeWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3jj7TUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QugiCp1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6brRyKpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pjeyXNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0jIjU3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func waOs9zzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U08ZxbdEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6zb6WX9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZuZAEAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OckSxmPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fml2uRwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3l8HqmcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgbGdNoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrdKuIOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVLYcp8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3TeNlCI0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZd61HFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgHIaR1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EAC97B22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pmm1MOCpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxD9mS1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ct2cMZQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 49BYYGO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxOu8KBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7NmORRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOVXm0RXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sO9tCd6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEa4jc4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RV2ZAmTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GobGh7JdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVFMf1jaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func StqbCv89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqen3yvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fScDUdDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlyWkWPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qH2xcdDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESIh8VFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWSeIYy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5m2SqXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nhzgiz4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wZDj4fcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJT28dqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Dv51PXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bP8twLNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydk9xgDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XhwZHFj3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nDHmSPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGTHMtEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afazHLA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFT15GlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVKY1g46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glWDOXeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKC7AW07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HCmRV7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VW0SgqeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfii3dTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPfR9nzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzewyDaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGLWECzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNtwgIZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTn134aPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYzyFNmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RScDWfT5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUpXbQgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2v0SjCeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXYex5FjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCoQGARZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzafDSm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CU0Xd3BWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26nfKsGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7bmfOnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fahGFBEmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUqFFPdNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWmbNQqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jh7LTDXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueaQasfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqXfQj4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOuL3G1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBWQ3oAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjFcTFGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyvT9IIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBQd0jtXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCGl14yAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxwvrQV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5aVnkH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVQTCs23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeCHbq8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpLioCbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6xPeoGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBlbVhbWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VtPBDosWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLDSgoGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KizTEifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tagmj7ssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARkgNXwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6CBgguvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzW0WdnkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cNpIbQqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWTArKDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkWm6qE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjrjhlJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gY1LYGCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3u9MLRWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71RfpzCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIRx45owWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZJ2ehVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFhMcqxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5mpcECwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWcvcbI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grvKZRpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5OTeQCW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yP9LXjPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ablqkQuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZj2MfixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRmnGX43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y6eG50bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dxXIKPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GmHuyIAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBvWr45BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdL3WKmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xZnJ9T7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2hQ8idWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBRn9lyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIH4UTmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EvHOEsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AezZWYnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8n8Gol3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oz9kYq9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10xheSS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lz8cheGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPRE7yxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7GaevfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 864dxtuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1h0T5yrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1jQGQfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HltxrYhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYVtPNo8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6REfVX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlbVsvbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqLjo05WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ujsLeZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6oRdm0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHcJ0VZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0uzqwYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7M328L6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEiB7S6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OzjUTixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WMphaWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kfdN9d8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJCUbGktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1bLRUmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8Haj6J4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TH6MhJD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjaLaCoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mA5ge4C0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1T3qKMB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvS53jptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxgcVaxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9lJl43WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vC1RhiCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivsx5685Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O0r1jxvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9dfhgHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIq5eWs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdPAtTHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywtHFLDPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBw5KOX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lf5TpFeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcO8jX6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5kjM2goWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pONUOfYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onkvjKfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKjuNZtLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SucVVlfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6iJi4YnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qA6rOj5IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IR3wlhxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPplu9dwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2bZd4AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omvMLNFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8JJ7h8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4a7UT1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcnWgpfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2Vxi8feWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epfUMpXlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhOLXRDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNis9UfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfoaD2FaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3JR61KoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACtpq43HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHACwM3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3K0aYXvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5c0LCevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dS65VxTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AmUmMNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZxKqcSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EaCNPeL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2y4bbvGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SeDFMeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urcKGXaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Nc0ps2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKGNV7RGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eGCw3JAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYnPpDdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xelvsElXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6WsYhrY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CH9mJAITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9QioPTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tboOn1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zFxH70qBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

