# Concurrency (Eşzamanlılık)

**Concurrency (Eşzamanlılık)**, birden fazla görevin aynı zaman diliminde yönetilmesi ve ilerletilmesini sağlayan bir tasarım yaklaşımıdır.

---

## Concurrency ve Parallelism (Paralellik) Farkı

* **Parallelism (Paralellik)** ile farklı olarak birden fazla işlemin gerçek anlamda aynı anda farklı çekirdeklerle çalışması gerekmez.
* Tek bir çekirdek, görevlerden biri beklerken diğer göreve geçer.

---

## Örnekler

* Python'daki `asyncio` veya Node.js'deki `event loop` ile tek bir thread'in binlerce isteği yönettiği asenkron yapılar örnek verilebilir.

---

## Go Dilinde Concurrency

Go, eşzamanlılığı (concurrency) dilin merkezine alan bir tasarıma sahiptir. Bunu gerçekleştirmek için **Goroutine** ve **Channel (Kanal)** adını verdiği iki güçlü mekanizma kullanır.

Go'da goroutine oluşturmak için fonksiyonun başına sadece `go` anahtar kelimesini eklemeniz yeterlidir.

---

## Race Condition (Yarış Durumu)

Eşzamanlılık kavramında en temel problemlerden biri olan **Race Condition (Yarış Durumu)** şöyle açıklanabilir:

> Aynı anda aynı kaynağa veya değişkene erişmeye çalışan kodların çıktısının belirsiz hale gelmesidir.