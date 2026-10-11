package main

import (
	"fmt"
	"time"
)

// sadece veri gönderebilen bir kanala (chan<- string) sahip fonksiyon
func gonderen(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		mesaj := fmt.Sprintf("#%d. Mesaj", i)
		ch <- mesaj

		time.Sleep(100 * time.Millisecond) // Mesaj gönderme işlemi arasında kısa bir süre bekletiyoruz
	}
	close(ch)
}

func main() {
	// string tipinde veri taşıyan bir kanal oluşturuyoruz
	kanal := make(chan string)

	fmt.Println("Mesajlar bekleniyor...")

	// "gonderen" fonksiyonunu arka planda eşzamanlı(goroutine) olarak çalıştırıyoruz
	go gonderen(kanal)

	//kanaldan veri geldiği sürece çalışan döngü
	for alici := range kanal {
		fmt.Printf("Mesaj alındı: %s\n", alici)
	}

	fmt.Println("Tüm mesajlar alındı.")
}

// Çıktım:
// Mesajlar bekleniyor...
// Mesaj alındı: #1. Mesaj
// Mesaj alındı: #2. Mesaj
// Mesaj alındı: #3. Mesaj
// Mesaj alındı: #4. Mesaj
// Mesaj alındı: #5. Mesaj
// Tüm mesajlar alındı.
