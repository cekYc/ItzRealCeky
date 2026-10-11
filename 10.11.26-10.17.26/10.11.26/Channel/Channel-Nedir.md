# Channel (Kanal) Nedir?

Channel (Kanallar), iki veya daha fazla goroutine'in birbirlerine veri yollayıp senkronize bir şekilde devam etmelerini sağlayan özel bir borudur. Eşzamanlı kodun çok daha okunaklı ve deadlock'a daha az açık olmasını sağlar.

---

## Temel İşlemler

* **Veri gönderme:** `ch <- değer`
* **Veri alma:** `değer := <- ch`
* **Kapatma:** `close(ch)`

---

## Unbuffered vs. Buffered

* **Unbuffered (`make(chan int)`):** 
  * Senkron iletişimdir.
  * Gönderen ve alan, karşısındaki hazır olana kadar bekler (*"El ele ver birlikte devam edelim"* tarzıdır).
  * Kesin senkronizasyon ve handshake (el sıkışma) içindir.

* **Buffered (`make(chan int, 5)`):** 
  * Asenkron kuyruktur.
  * Gönderen buffer dolana kadar durmaz, yazıp geçer. Buffer boşalınca alan bekler (*"Şuraya bırak ben sonra bakarım"* tarzıdır).
  * Hız farkı olan işlerde queue (kuyruk) oluşturmak içindir.

---

## Basit Örnek

```go
func main() {
    ch := make(chan string) // unbuffered

    go func() {
        ch <- "selam" // veri gönder, karşı taraf alana kadar bekle
    }()

    mesaj := <- ch // veri al
    fmt.Println(mesaj)
}
```

Bu kodda `main` goroutine, arka plandaki goroutine’in mesajı göndermesini bekler.

---

## Channel Olmasaydı: Shared Variable + Mutex

Channel olmasa shared variable + mutex şöyle yazılırdı:

```go
var (
    mesaj string
    mu    sync.Mutex
    ready bool
)

func main() {
    go func() {
        mu.Lock()
        mesaj = "selam kanka"
        ready = true
        mu.Unlock()
    }()

    // Hazır olana kadar dön (işlemciyi meşgul eder)
    for {
        mu.Lock()
        if ready {
            fmt.Println(mesaj)
            mu.Unlock()
            return
        }
        mu.Unlock()
        time.Sleep(time.Millisecond) // biraz nefes al
    }
}
```

---

## Tip Güvenliği İçin Yönlü Kanallar

```go
func gönder(ch chan<- string) { // sadece gönder
    ch <- "geldi"
}

func al(ch <-chan string) { // sadece al
    fmt.Println(<-ch)
}
```

---

## Ne Zaman Kullanılır?

* İş dağıtmak için
* Sonuçları toplamak için
* "Bu iş bitsin sonra devam edelim" demek için
* `select` ile birden fazla kanalı aynı anda dinlemek için
