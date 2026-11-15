package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.0" )

func zD8okUOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2XwWq2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZogJSK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zI3Ru3rcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s69PNge3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5KA4qwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ry9ZNcLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThF4rqt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kd8qdhtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXqwuhNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHbG2HajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8UP5sG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnx2pr4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cmtc4lu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUltmQIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHwVWidoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SgSXCmKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NLo25eAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xhSSNAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TeC3OEzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFsUffsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SElW73rlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1jyEOeR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e80H5SZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func klmABK0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvdIrOuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRMoPwDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8bFIsGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8k7Mi0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fex3y0CeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8bhhzfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1JeBzs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPEjA94FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZijm9nbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98Dl88mqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBENjrKUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dq4TODVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQ6ARKeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caLt5pGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwDUpMcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxhCHzSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gm4WOL4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbnxnIWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SDenJsGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qi7flhoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJx4oVCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AXdbmr5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lu5yIzi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOMpTdjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sqgx5MDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7arQwwGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nw11Trr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wy8FIV9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSdMAznmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8O6hm3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blhBwWanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zcI7v5CfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFbTfA1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsXjrke0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F46bAiQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYCl2QULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e158JAFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFtxeI2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCNmCDdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fS3c4lUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJbsjSQlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrAPKWCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cb3BELs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyoBlLkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeiPeH9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeq8Zs98Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NoFD8e13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMOygsBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9rBLlqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtyDdvJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ipk8pK5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDGHnVD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PNlChzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQ6Ix1JzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMLAJCGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3Gzh66vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyBEoaO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pRLF1w85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBJlQH9FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0TvmG7xTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MyRpb9ykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8OrR0CUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BI4Ns48sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iE9E2xSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfAItJVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uE6HZNpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHX62UocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u5Du89QOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sh1Mwf4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxgsArQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKGlNc0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0u1VclYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xCATeJPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4KZdolnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIu0nLhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RNwuccyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2hNjHL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func guTyp4eCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuFOhMDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5m5H45bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgcfKcZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIKCtTlmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Myvzu9gpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1yTACW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QInri5fTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3z3EJrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGOTqGSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bQiykJPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FX2MdH3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUmQZPksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfOXIhjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nl56zA85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8XdVdi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jh9RYSk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DEMx0afJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcfCj5UKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLUtIdrrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2qoUx89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tz3MXdRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSqCAcJKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtiB6IS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhExVVfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNZKPkXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZuI5kd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wLQsO6mtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsLg7XfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efKcJmOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2WqPnBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJz20rkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYcAEYENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNZHkEvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0olXQUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZXO2XNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJKaBoi9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rbIuPL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ge0soj4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xG4eA3swWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29Yn8WptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8wGoW1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qc5APFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJxJu35FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8fl78AnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfhOWTBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbTS7a18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AKKQkNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7W6jq7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JRWPwvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynxLSXoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kpx2Qj8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wtwPt3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QuOMcA8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CS2tIknHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VaKnZdwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3p7xrF4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ytqlth8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paKEFzrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MP2jupUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBo415bAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y08fImq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsBEprELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPjrRSMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T7ljzDiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tU1RkndMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLyeUCODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func As2qLvSpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylHsLC8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOT8JvoHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykxabwRrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0FKZPJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRpetgCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moZQffeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjL3w3fTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YsoY7637Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5G78Y5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIUnH3e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3tT4HpwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHA244cgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8NXzkgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwVhF3aGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6muQPu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I36pR7CBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGL0dwxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func see8HHBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kTELsCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17EOZIc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ew1nvKcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7FT9UBqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IHVWci1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vALAlAe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6Nw2uOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fs0uU7sNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jW3hhl0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zx2kRY1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kx9DQVTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gl8uwvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XtqHKIr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vj2tRC6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AYgD1LpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SD0UXXzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4CSwzzEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IDe9dMVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDTPgBFtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NHGsI4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWHOXEpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4MgQU6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAmX5O4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOl1TeqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dyg1tuyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BKbXNHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLpL0XPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsADSDtqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tORJdDrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etPzUOpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M3Gf13f4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x59ePejUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ebh9grtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJVGkYb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuFPZK94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNrUjHYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyudbuRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g988WjiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Agkhl6fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqX8lO6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QmhwD0FgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mVl5NtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffH5tKPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXF5dWJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHNCxDLYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3xCFUsmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnUAHrH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASlN1jBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByzLQEBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kOCEbSUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oC3mIOxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OYSYXpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7YbabZgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yGyLVWrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzMHc9axWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhB14rm2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKH3hH0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2f3YeZFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZiFn2ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToZqjB2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ty1vTRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pdVXrBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Di3XGSC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56qnrYDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmeaHDtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4ccmCZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lc1NPe0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2nBUuMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKy4uyGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbaEbgB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJWnDR4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fISWXoTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SQZiZj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pgqqp446Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZgG8TabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oItyaGAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQsA9tuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3C48qKIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aspsibRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXgy3r3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 651QM6zKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9g5rc4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiOO5gI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y20iCNs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sp994FApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VnhkgeXjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYaams9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPMw7IdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38X09yEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxQ6P8d8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vM5b8MocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sa6XYlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FdUYAMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbfKVpOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Du1CV42OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nAdpg1fjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func As0YWsfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXkMtx6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N56b7YKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mykllHm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MoQLZXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOBRffq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hOPUmcYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RCMic8g7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLkaGlMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugs1EC2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4bOwqJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKuHFcybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2tnpIvIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXG7DyJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6kXfdGXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5swRLPLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PGwv0BvOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqujfomlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmyqXegUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cbm57iPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vXcaNpbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kk40CdJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tl7upJtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsKHGBe6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g87U6QpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2mrP6BsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7RdhhilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRAo84gjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoxExsEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQpB6dACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsDAQwL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqG7wfj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXfs5D9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnZRLdf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IrXF0LbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lew8XX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dph3Tsr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T12mubirWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYGklIW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T17RJN3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0qCeHc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bI9VRtioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbBzTGLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Olyv8arjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFPllEeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GAc3uCgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqfYrKqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dev4NuyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USOtTsYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFC48j1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbmqZMkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EIdYBFMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPD6p6BXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfUjjpxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJwGjKvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYRdtPYWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wPnWrzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGW129ImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8J7ilp22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDEbGKaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3fZUvhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFrKB8NEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2JWtUqBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyDvZyZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8ij1AkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrwlA9NvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func By5gEcPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URhNWBNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HVBRYB3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OyoTCSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdBcZCLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsD9YJ1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWtnUgdTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func faNQ5dGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DuwLWxcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkslTMHvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6B4xVF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9BvBuNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fAsUUnQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gMZZTHxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erw7viLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLXD8vzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBTLRI7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0ygpeB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30dNleokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVzhoWHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QvyseLbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1qml7J1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpmyKvwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jfc9eSLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QmIXkUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5stQIvYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqMs6mzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpHZR3TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pqg7fy0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HGHnxW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCiTGRLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4rqUVtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CdVAcvR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruCjY0ZuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dV31aAQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tt3JXWUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAWNmRMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHEWRv3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lU8LiIwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLmRNShSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m0R7f1pJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hb0pJERCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unUS20DfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jyeXftEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zV2aM6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lah5G7sGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63S0ESrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzFEPteXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FwfMnvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E4kQkOoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func La1Ql0zNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDpO3FYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5huSCedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scoBlyfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEehhKAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yovm788KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fdBaTnQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iThxIK4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prXFpA2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnkPfG0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNcj2y0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdS4otLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nylakgFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lm1EjYlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPQxNEbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTTyirGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcDb5WOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WW8nD5xrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dF6uX1zkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AxOWSNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xR19i1prWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZAl4OP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JihbsTVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gq80SWh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2eKTIu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51VQVsnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YX8ZzlUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpyimuBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RO1G1rOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MvwtHN5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoQ6DWpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1oVBoeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgFgb58hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jPBWLsWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O7sinIGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WYFnW322Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUMQibvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cm1WAtbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c7jsRzrUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxg0oCTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfOeEqHTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EPgCaPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TO9iv6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BThDPWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kduudKjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qiJBzhrSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 039hCCiHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func epGzRk0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3I5g1gzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHPZioEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dxvD7JnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsIB6wu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nHxOcSSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zN5UV6pLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IF6ye43RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3gEwcmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNhscWdCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Q6ByRUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhB4sGinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBNaAdMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OpZl944zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uemlVcdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkW9k9leWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgf0j46IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnqikCMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DwvYy3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrCscHpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XUMId8iSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpQkKLEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qjtYcuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgDs7Py7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pd4ze9H2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDPYNDCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fehPYt5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlVT7uyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qazvywZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZpDGYSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NhUr95pUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 68fjWJGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhvfdM5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKRkhmvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7HyC0EUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHvOFuCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2nlgVbYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8eDlgUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNqfZsx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2A8mrbmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjuCYRxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqAZVgYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ut4M82RqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0I0ptsf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KL2wogwZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aERDvIK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zaDtBEYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYVUj955Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WphInsdqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ogu7Se9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RwMNlKLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4UYEXgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQ6yhr3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7mjGlHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cn7Lq5SZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nd6ORlXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RexzukiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qafAtQotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhElL1ZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cRR4TZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vD2PXc3AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QUxkq9xYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jENh8jpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyoJnDXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsHCV6KpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9in0RIVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YNYjFtaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eaAXcgQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBHAXIOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPoXWsgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YuldG9lBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVMHnbazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDw49xIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPwSzHxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzF1gJz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfigSOzxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWrWhu8oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXkzESV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mh3opLBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aVv9o8DSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxFFbXaCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i9eODMAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47QgVgWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7jlVICFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7t9fVW2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Pw00CKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFN6321oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfzNQcUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aZWqA4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvZtEBMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBqHc6PNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SW64RX47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPbyU7xvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wAu434OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbdrmEZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYVifqZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxeofgbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nENnke09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yIiStk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cg7wzcihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXlOAh7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYh4lnL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yYLXC27SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mX7qkB3WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ddh36IGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2rt6m78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moOZc441Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoJt35XgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxmeXgshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUCJcwQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iM51jJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhW7fsDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YUs743gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X37TWrxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnM0FFFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HkfIQYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZjHTE2fFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaJ8MUdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r4Nh9R1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5BpX1nPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5Bh0Gp6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vd1hQL7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qU1d8JMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Njlx2oNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJLkfOdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtrteBIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yc3ABpDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Usj6wFZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRgOXjP5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPMXcECNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JialMyP0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ahZuUKogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBLjArZ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJwiACyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eDG3cZQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1oH6uuDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cHvY9Ur1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Achsqi4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hd1hX7w9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xelSZzfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mIscRpvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xp3JA6kPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f4h16GAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXzl0DyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9C3AzSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjcYeg60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q8AZyhMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CBvYzS4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IE6UqjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qxNI5qQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oj1cDNuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztrhMbsmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCW63xCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWPKsdA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sxa1Lc9TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F64GcPVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYYD2QhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqC1eEYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnvsXH4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOyPeNENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NcjGC0CgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9WKXZZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYqwA7PfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpScChQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGaoknBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6B5sqGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JMfQwBdyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QEDXzosNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Emie7NaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGcW5qdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DyrlnNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSFBNw5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aWvDIl8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nisHES3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3xn4HWwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aW88wH7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpZ9qNloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9MikzZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbdGOw2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPBpn1o9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TncQo9abWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1A64nUcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTh6iHUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6Pfmd6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z1OT6r7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQwLqh6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpUAIhtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4jADjwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3PybPJEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXJzhKHRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjL59qGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwlOPUYvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pL4NLDjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func edEb2m19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CXUHy2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJqeVLUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NwZXQ4ZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJHEAZmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1OZC5AUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSLx7LyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOuWd53XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRntzriYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxGTGM64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xp5nYIUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwBk4s9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJPTj4llWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNUlqLreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCBhgeNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoBBfUhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S6BAGV9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyP84N0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkbDWlVOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5awMTlHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U38bHZpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOpsFW1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHzntErZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDZ6Vca2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLSYt5BaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYupXhBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSqNrHl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pBjZ0gkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mL7g8cNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uWUFnFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zyuILzw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxKL3k2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2vZU3oqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDBVIT9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xESAl9WMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVx1mm7wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtwJeR8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tswz0Xh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxLaSQPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23tLrnnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAlJQudAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lni6cwfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Himdeo59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npBjus7pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UIhQiMCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vARnz0DuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqGK5eOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqLX8LZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pA9IGHjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func irrKB6xxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C8Ix2GynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zo826WiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rr955zPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhMAxy7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func op7PvwelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4HPY9vdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tRO537uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzM2K8ImWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kTuJHnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PE8LSSn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obOS8QFYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPNTGWb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWJioUbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7UV6DPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ho1QtJ18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWJI5M3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCpvnkiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZVn1Z4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jq3C9IfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvQPViGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1srYjVjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCnPV23lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGZggjccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWm5ZaSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJw6gA3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRTZJN6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xqh0owqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tB1uocb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oJZgKoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7oXYYw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dp7xwjvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuIvfaX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGgJgPTDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGKGL7S6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVynvvWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yr4QFt4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8IAq3lqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhiY6qyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxjzpqVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TeRJYSIsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJDZug0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjmJSwkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TpY8GmrFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fg4TNR4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFIOgD2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUpSd3DKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QTXhVmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qencGcmuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XokhMPqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H37TfijdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ognmmZX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izM1r86VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgHzgqXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqyEKMEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cAYF9FHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EMCHcaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJElunxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgBkes4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcIdwUdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PcKpBJfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ArwHWj5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3Obvm8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7exC4DwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCEYLwbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFzHO3EaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func veYw2zd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

