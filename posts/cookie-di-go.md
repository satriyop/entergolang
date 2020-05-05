---
title: "Cookie Di Go"
date: 2020-04-26T01:10:38+07:00
draft: false
tags: ["cookie", "go","http"]
categories: ["go"]
---
Sering kita mendengar tentang _cookie_ ketika membaca artikel tentang pemrograman web. Bila kita mengetikkan cookie saja di mesin pencari google, maka yang didapatkan tentu saja informasi tentang dunia per-kue-an yang mengundang rasa lapar. Tapi hasil pencarian ini akan sama sekali berbeda ketika kita ganti kata kunci pencarian kita menjadi "cookie http". Salah satu referensi dari cookie yang akan kita bahas adalah dari mozilla, yang bisa dibaca di
[sini](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies).


Terlepas dari kontroversi dari cookie, terutama yang menyangkut privacy yang bisa dibaca di [sini](https://www.cookiebot.com/en/gdpr-cookies/), cookie tetaplah menjadi elemen yang fundamental dalam membangun aplikasi berbasis web/http. Hal ini dikarenakan sifat dari protokol HTTP sendiri yang [stateless](https://en.wikipedia.org/wiki/Stateless_protocol) sehingga dibutuhkanlah sebuah metodologi untuk membuat aplikasi mampu mengenali satu pengguna dari pengguna lainnya tanpa
harus melakukan proses otentikasi berulang-ulang.

Pada bahasa pemrograman Go, disediakan fasilitas untuk mengatur cookie yang akan diterima/akan dikirmkan ke client (dalam hal ini browser, karena yang akan menyimpan cookie adalah browser yang digunakan oleh pengguna aplikasi). Cookie di Go, ada di pustaka standar [`net/http`](https://golang.org/pkg/net/http/#SetCookie) yang tentu saja sangat masuk akal, karena cookie erat hubungannya dengan http.

Untuk mengeset sebuah cookie, kita bisa membuat sebuah _handler_ khusus, misalnya seperti dibawah ini :

    func createCookie(w http.ResponseWriter, r *http.Request) {
	    cookieName := "session"
	    c := &http.Cookie{}

	    if c.Value == "" {
		    c.Name = cookieName
		    c.Value = "qwerty"
		    c.Expires = time.Now().Add(2 * time.Minute)
		    c.Path = "/"
	    }

	    http.SetCookie(w, c)

    }

bila kita akses url path yang akan dihandle oleh _handler_ ini misalnya /cookie

    mux.HandleFunc("/cookie", createCookie)

maka kita bisa mengetesnya dengan menggunakan curl dengan command seperti ini :
    
    curl -X GET -i localhost:3000/cookie

dan kalau kita amati 

    HTTP/1.1 200 OK
    Set-Cookie: session=qwerty; Path=/; Expires=Sat, 25 Apr 2020 18:43:10 GMT
    Date: Sat, 25 Apr 2020 18:41:10 GMT
    Content-Length: 0


maka kita bisa melihat _response header_ yang dikirimkan dari server sudah menyisipkan cookie dengan jangka waktu tertentu.


Cookie memang terlihat sederhana, dia tidak lain adalah sebuah _text_ biasa yang dikirim berulang dari client ke server, dalam hal ini kita mengeset cookie dari sisi server, tapi cookie juga bisa diset dari sisi client (dengan menggunakan Javascript), tapi hal ini kita tidak bahas di artikel ini.

Sebagai catatan yang menarik, cookie ini diciptakan oleh [Lou Montulli](https://en.wikipedia.org/wiki/Lou_Montulli) saat dia bekerja di Netscape pada tahun 1994. Teknologi yang sudah lumayan tua inipun masih bertahan  (dan khusus untuk server-side rendering, mungkin satu-satunya metode yang bisa digunakan sampai saat ini).


Berbagai isu keamanan yang muncul dari cookie ini antara lain XSS, XRSF, XST sehingga programmer diharapkan sangat berhati-hati saat melakukan logika pengesetan cookie. Beberapa faktor yang harus diperhatikan antara lain :
1. Expiration dari cookie
2. Set Secure Cookie
3. Http-only Cookie
4. Same-site Cookie

Dalam hal privacy, bila web kita akan melayani pengguna di kawasan Eropa, maka harus menaati peraturan GDPR (General Data Protection Regulation). 

Demikian singkat bincang tentang cookie, tentunya tidak lengkap kalau tidak ditemani segelas kopi.

