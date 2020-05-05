---
title: "Secure Cookie di Go"
date: 2020-04-29T02:37:53+07:00
draft: false
tags: ["cookie", "go", "gorilla", "hash", "hmac"]
categories: ["go"]
---
Seperti postingan tentang cookie sebelumnya, bahwa cookie adalah salah satu teknologi yang cukup rentan digunakan dalam hal komunikasi antara browser dan server, apalagi bila dipakai cookie tersebut dipakai untuk menyimpan data-data yang cukup sensitif seperti data otentikasi user atau yang lainnya. 

Meskipun begitu, tentu tidak bijak juga bila kita membiarkan cookie dengan tingkat keamanan yang apa adanya, meskipun tidak digunakan untuk menyimpan data sensitif. Salah satu cara untuk meningkatkan keamanan data di cookie adalah dengan menerapkan _secure cookie_.

Meskipun kita bisa membuat sendiri implementasi _secure cookie_ ini dengan pustaka bawaan dari Go, sudah ada pustaka yang cukup lazim digunakan secara umum untuk menerapkan _secure cookie_ ini, yaitu dengan menggunakan [Gorilla Secure Cookie](https://github.com/gorilla/securecookie).

Untuk menggunakannya cukup sederhana, silakan dibaca sendiri ya di [sini](http://www.gorillatoolkit.org/pkg/securecookie).

Salah satu yang pertimbangan penting menggunakan _secure cookie_ dari [Gorilla](http://www.gorillatoolkit.org/) ini adalah bagaimana dia menerapkan teknik otentikasi dan enkripsi di library ini.

Secara umum, library ini menggunakan [HMAC](https://tools.ietf.org/html/rfc2104) sebagai metode untuk mengamankan data cookie. Secara default library ini menggunakan SHA256 sebagai algoritma hashing dan menggunakan AES sebagaimana saat kita menginisiasi secureCookie 

    // Hash keys should be at least 32 bytes long
    var hashKey = []byte("very-secret")
    // Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
    // Shorter keys may weaken the encryption used.
    var blockKey = []byte("a-lot-secret")
    var s = securecookie.New(hashKey, blockKey)

ini bisa kita lihat di kode dari library ini, terutama di bagian fungsi constructornya

    func New(hashKey, blockKey []byte) *SecureCookie {
        s := &SecureCookie{
            hashKey:   hashKey,
            blockKey:  blockKey,
            hashFunc:  sha256.New,
            maxAge:    86400 * 30,
            maxLength: 4096,
            sz:        GobEncoder{},
        }
        if len(hashKey) == 0 {
            s.err = errHashKeyNotSet
        }
        if blockKey != nil {
            s.BlockFunc(aes.NewCipher)
        }
        return s
    }

di situ di set secara default bahwa fungsi hashing akan digunakan sha256, meskipun dimungkinkan untuk diganti sebagaimana dia sudah memberikan fungsi ini :

    // HashFunc sets the hash function used to create HMAC.
    // Default is crypto/sha256.New.
    func (s *SecureCookie) HashFunc(f func() hash.Hash) *SecureCookie {
        s.hashFunc = f
        return s
    }


Postingan ini hanya membahas tentang sebagian kecil, _di bawah kap_ dari library ini, khususnya tentang hashing karena berbicara tentang HMAC yang biasa dipakai untuk secure file transfer, HMAC menyelesaikan 2 hal yaitu :
1. Data integrity (yang disediakan oleh fungsi hashing, untuk membuktikan bahwa file/data yang dikirim tidak berubah)
2. Data Authenticity (bahwa yang mengirimkan file/data tersebut memang betul si pengirim)

Lain kali bila ada kesempatan kita bisa membicarakan point no 2 ya. Semoga.

Standard tentang HMAC sendiri bisa dibaca di dokumen [IETF](https://tools.ietf.org/html/rfc2104)


ps : Tetap saja, meskipun kita sudah menggunakan library ini, tidak disarankan untuk menyimpan data sensitif di cookie. Camkan itu anak muda !