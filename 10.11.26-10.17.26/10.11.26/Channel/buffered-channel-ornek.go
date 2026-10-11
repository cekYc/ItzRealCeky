package main

import (
	"fmt"
	"time"
)

func main() {

	// buffered channel oluşturuyoruz
	oge := make(chan string, 3)

	// buffered channel'a veri gönderiyoruz
	oge <- "1. veri"
	oge <- "2. veri"
	oge <- "3. veri"

	fmt.Println("Buffered channel'a 3 veri gönderildi (Kanal doldu).")

	// kanalın mevcut kapasitesini ve doluluk durumunu kontrol ediyoruz
	fmt.Printf("Kanalin kapasitesi: %d, Kanalin doluluk durumu: %d\n", cap(oge), len(oge))

	// eğer 4. mesaj gönderilmeye çalışırsak:
	// oge <- "4. mesaj" // Bu satır çalıştırılırsa program kilitlenir (deadlock) çünkü kanal doludur ve alıcı yoktur.

	fmt.Println("ogeler dagitiliyor...")
	time.Sleep(500 * time.Millisecond)

	// buffered channel'dan veri alıyoruz
	x := len(oge) // kanalın doluluk durumunu alıyoruz
	for i := 1; i <= x; i++ {
		mesaj := <-oge
		fmt.Printf("Alinan mesaj: %s\n", mesaj)
	}

	fmt.Println("Buffered channel'dan tüm mesajlar alindi.")
}

// Çıktım:
// Buffered channel'a 3 veri gönderildi (Kanal doldu).
// Kanalin kapasitesi: 3, Kanalin doluluk durumu: 3
// ogeler dagitiliyor...
// Alinan mesaj: 1. veri
// Alinan mesaj: 2. veri
// Alinan mesaj: 3. veri
// Buffered channel'dan tüm mesajlar alindi.
