---
title: "Perkenalan Pustaka *net/http* di Go"
date: 2020-04-24T22:49:28+07:00
draft: false
tags: ["http", "go", "api"]
categories: ["go"]
---

Bahasa pemrograman modern biasanya sudah menyediakan fitur server (biasanya web/http server) sebagai pustaka standar, begitupun dengan Go.

Untuk membangun sebuah server dengan Go, yang akan menangani (handle) komunikasi dengan protokol HTTP, kita bisa menggunakan pustaka **net/http**.
Sebelum berusaha membuat program dengan basis komunikasi HTTP, ada baiknya pembaca mempelajari dahulu informasi detil tentang protokol ini di [sini](https://developer.mozilla.org/en-US/docs/Web/HTTP/Overview).

Sebagai sebuah penyederhanaan, pada dasarnya komunikasi HTTP akan melibatkan dua jenis _message_ sebagai sarana komunikasi, yaitu **request** dan **response**.

Saat kita mengakses sebuah website dengan menggunakan browser (firefox, chrome, Internet explorer), maka bisa dianggap browser dianggap sebagai client yang akan mengirimkan _request_ ke sebuah web server (misalnya : google.com). Web server dari google.com, kemudian memproses _request message_ tersebut berdasarkan bermacam kriteria dan logika, kemudian mengirimkan _response_ kembali ke browser and menampilkan hasilnya ke pengguna.

Balik ke Go, dengan menggunakan pusataka _net/http_ ini, kita bisa menerima _request_ dari client , melakukan _handling_ terhadap _request_ tersebut dan mengirimkan kembali _response_ ke client. Dengan kata lain dalam melakukan pemrograman web dengan Go, kita hanya fokus pada bagian _handling_ tersebut yang mencakup 2 hal yaitu : 
1. Menangani _request_ , misalnya melakukan pencocokan url yang diminta di _request_ dan menentukan _handler_ mana yang akan menangani.
2. Menulis dan mengirimkan _response_ kembali ke client oleh **Handler** tertentu dari langkah pertama. 

Dalam hal menangani _request_ , di langkah pertama, pustaka net/http menyertakan [**ServeMux**](https://golang.org/pkg/net/http/#ServeMux), sebagai _multiplexor_ dari _request_ yang diterima, sehingga bisa menentukan handler mana yang bertanggung jawab menangani _request_ tersebut.

Selanjutnya, _handler_ yang ditunjuk akan menangani lebih lanjut lalu  menuliskan _response_ .  _handler_ tidak hanya menuliskan di _body_, tapi juga di _header_ . Lagi-lagi detil tentang masing-masing **HTTP Message** bisa dipelajari dari artikel Mozilla di atas.

Pustaka net/http sendiri sudah menyertakan beberapa fungsi yang akan menghasilkan beberapa handler yang umum seperti _FileServer_, _NotFoundHandler_ dan _RedirectHandler_.

Semua Type bisa menjadi sebuah _handler_ , asal mempunyai sebuah method ServeHTTP dengan _signature_ seperti ini :
    ServeHTTP(http.ResponseWriter, *http.Request)
atau dengan kata lain, semua object yang menerapkan _interface_ dari [**http.Handler**](https://golang.org/pkg/net/http/#Handler).

Sehingga, dengan memahami ini kita bisa membuat sebuah _handler_ yang kita sesuaikan dengan kebutuhan kita (_custom handler_).
Contoh umum dari _custom handler_ ini sebagai berikut template nya :

    type customHandler struct {
        name string
    }

    func (cH *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        // contoh kode atau logika
        cH.name = "Ini adalah custom handler"
        // Kirim response body
        w.Write([]byte("Halo " + cH.name))
    }

Okay, kita sudah mempunyai sebuah _handler_, mari kita masukkan ke program utama kita seperti ini :

    func main() {
        mux := http.NewServeMux()
        ch := &customHandler{}
        mux.Handle("/customHandler", ch)
        log.Println("Listening....")
        http.ListenAndServe(":3000", mux)
    }

Kita bisa tes program kita, yang bisa dibilang sebuah "mini-API" ini dengan menggunakan curl / PostMan / Insomnia sebagai tools.
Bila menggunakan curl, yang mana tool ini sangat sederhana tapi _powerful_ , kita bisa menuliskan ini di terminal (dengan catatan curl sudah terinstall).

    curl -X GET localhost:3000/customHandler

Maka kita akan melihat _string_ yang sudah dituliskan di _body response_ oleh _handler_. Tentu saja, _handler_ yang kita tuliskan ini masih sangatlah dasar, karena dia akan menanggapi segala jenis _HTTP Verb_ (Tidak hanya **GET** seperti contoh di atas). Sehingga kalau kita coba mengganti _HTTP Verb_ menjadi **POST**, kita akan mendapatkan hasil yang sama.

    curl -X POST localhost:3000/customHandler

Sampai di sini dulu, untuk lebih mendetilkan HTTP Verb apa saja yang bisa ditangani oleh Handler, bisa dilakukan pembaca sebagai latihan/riset mandiri. Petunjuknya adalah, di bagian Handler, kita bisa mengecek HTTP Verb dari _request_ sehingga bisa bisa menentukan _response_ apa yang akan dikirimkan. Ayo, kamu bisa !



ps : Oiya, alih-alih kita membuat type baru (dalam artikel ini type customHandler) untuk hal yang sederhana seperti contoh di atas, kita bisa juga menggunakan fungsi biasa, dengan catatan fungsi itu mempunyai signature seperti di bawah :
>func(http.ResponseWriter, *http.Request)

Tentu saja ini juga menjadi hal yang bisa dieksplorasi lebih lanjut oleh pembaca.