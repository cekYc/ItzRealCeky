package main

import (
	"fmt"
	"time"
)

// belirli bir aralıkta çıktı üreten fonksiyon
func islemSayar(id int, sure time.Duration) {
	for i := 1; i <= 4; i++ {
		fmt.Printf("Goroutine %d: %d. işlem yapılıyor...\n", id, i)
		time.Sleep(sure)
	}
}

func main() {

	// "go" anahtar kelimesi ile fonksiyonları arka planda (goroutine olarak) çalıştırıyoruz
	go islemSayar(1, 100*time.Millisecond)
	go islemSayar(2, 150*time.Millisecond)
	go islemSayar(3, 250*time.Millisecond)

	// main fonksiyonu bitmeden önce tüm goroutine'ların tamamlanmasını beklemek için bir süre bekletiyoruz
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("Tüm goroutine'lar tamamlandı.")
}

// Çıktım:
// Goroutine 1: 1. işlem yapılıyor...
// Goroutine 2: 1. işlem yapılıyor...
// Goroutine 3: 1. işlem yapılıyor...
// Goroutine 1: 2. işlem yapılıyor...
// Goroutine 2: 2. işlem yapılıyor...
// Goroutine 1: 3. işlem yapılıyor...
// Goroutine 3: 2. işlem yapılıyor...
// Goroutine 1: 4. işlem yapılıyor...
// Goroutine 2: 3. işlem yapılıyor...
// Goroutine 2: 4. işlem yapılıyor...
// Goroutine 3: 3. işlem yapılıyor...
// Goroutine 3: 4. işlem yapılıyor...
// Tüm goroutine'lar tamamlandı.
