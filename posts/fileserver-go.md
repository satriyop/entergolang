---
title: "Fileserver Go"
date: 2020-04-28T00:54:33+07:00
draft: false
tags: ["http", "go","fileserver"]
categories: ["go"]
---

Membuat fileserver dengan Go, adalah urusan _sepotong kue_.
Pustaka standar dari Go, yaitu `net/http` menyediakan fasilitas ini sebagai salah satu `Handler` yang bisa kita gunakan untuk melayani _request_ http dan melayani konten dari _file system_ OS yang dipakai.

Berikut di bawah adalah contoh sangat sederhana dari membuat file server di Go

    func main() {
        fmt.Println("File server on Go")

        // directory yang akan dipakai fileserver
        dir := http.Dir("/")

        // buat handler untuk serving konten di file system
        fileServer := http.FileServer(dir)

        // register handler fileServer untuk menangani request di path "/"
        http.Handle("/", fileServer)

        // start server di port 8000, dengan default ServeMux
        log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
    }

Perlu diperhatikan bahwa kode ini bukan kode production, perlu diperhatikan untuk filter jenis konten yang disediakan dan faktor keamanan lainnya.


ps : semua kode tersebut di atas bisa sangat dimampatkan lagi seperti ini :
    
    log.Fatal(http.ListenAndServe(":8000", http.FileServer(http.Dir("/"))))


