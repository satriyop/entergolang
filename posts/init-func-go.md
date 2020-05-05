---
title: "Init Func di Go"
date: 2020-04-30T02:10:49+07:00
draft: false
tags: ["init", "go","package"]
categories: ["go"]
---

Salah satu nama fungsi yang spesial di Go adalah init. Seperti namanya, kita bisa menebak bahwa fungsi dari `func init` ini adalah untuk melakukan sebuah inisiasi dari value-value tertentu yang dibutuhkan oleh package main.

Sebelum membahas tentang detil func init, alur dari eksekusi sebuah package main.
1. Inisiasi dari package main
2. Pemanggilan func main

Pada tahap pertama, yaitu iniasi package main hal-hal berikut secara berurutan akan dilakukan :
- Apabila ada package lain yang di-import, maka package tersebut akan diinisiasi terlebih dahulu.
- Package-package tersebut (bila ada lebih dari satu package yang di-import) diinisiasi satu demi satu.
- Variabel-variabel dari package-level pertama diinisiasikan dalam urutan deklarasi.
- Pemanggilan func init

atau secara singkat seperti ini :

`import --> const --> var --> init()`

Pada tahap kedua, ketika func main _return_, yang perlu diperhatikan adalah dia tidak menunggu (bila ada) Go Routine untuk selesai. Ingat, func main akan segera mengakhiri aplikasi/program apabila dia mencapai return.

Kembali ke func init. Fungsi ini bisa dideklarasikan dengan cara seperti ini :

    func init() { … }

Contoh dari implementasi func ini adalah sebagai berikut 

    var y int
    func main() {
        x := 10
        fmt.Println(x + y)
    }

    func init() {
        fmt.Println("Init 1 dipanggil")
        y = 20
    }

Meskipun variabel y tidak diinisiasi di func main, tapi di func init.

Kita bisa mempunyai lebih dari satu func init , misalnya kita bisa menambahan func init lain sebagai tambahan dari kode di atas.

    func init() {
        fmt.Println("Init 2 dipanggil")
        y++
    }

    func init() {
        fmt.Println("Init 3 dipanggil")
        y--
    }

Urutan dari func init ini dipanggil akan selalu urut dalam hal ini. Tapi tidak selalu ketika ketika mengimpor package/library lain (yang mana mereka punya func init sendiri-sendiri).



Detail dari init, bisa dibaca kembali di [effective Go](https://golang.org/doc/effective_go.html#init).