---
title: "OOP : Encapsulation di Go"
date: 2020-04-30T16:09:30+07:00
draft: false
tags: ["oop", "go","encapsulation"]
categories: ["go"]
---

Salah satu filosofi dalam pemrograman berorientasi object adalah bahwa _consumer_ (dalam Go, bisa diartikan package lain yang ingin berinteraksi dengan package kita) tidak boleh terlalu banyak tahu tentang detil object yang akan di-_consume_/dipanggil.

Hal ini bisa diselesaikan dengan 2 hal secara umum yaitu _Encapsulation_ dan _Message Passing_, kita bahas dahulu tentang _encapsulation_.

Dalam hal _encapsulation_ berarti _consumer_  dapat meminta/memanggil sebuah _service_ (biasanya dalam bentuk method) pada object yang di-_consume_-nya, tanpa perlu tahu detail bagaimana _service_ tersebut diimplementasikan (logika dan struktur datanya).

Go tidak punya _Class_ dan tidak punya _access modifier_ seperti _public/protected/private_ sehingga dalam hal encapsulation kita punya strategi yang berbeda yaitu :
1. Package Oriented Design
2. Interface

## Package Oriented Design
Dalam bahasa pemrograma berbasis object seperti Java/C# melakukan desain object dengan menggunakan Class (Beberapa hal yang bisa dilakukan di Class bisa dilakukan juga dengan Type kalau di Go contohnya dalam hal definisi struktur data dari object), tapi Type tidak bisa disamakan dengan class secara harfiah!.

Kita bisa membuat method-method yang bertugas sebagai _accessor/mutator_ (getter/setter) terhadap object (dalam hal ini pada level _package_)

    package payment

    type BankAccount struct {
        accountNumber   string
        accountOwner    string
    }

    func(b BankAccount) AccountNumber() string
    func(b BankAccount) AccountOwner() string


Bila kita punya struktur data lain, misalnya CreditAccount pada satu package (dalam hal ini payment), maka CreditAccount akan punya akses yang penuh terhadap field-field dari BankAccount. Tapi bila kita punya package lain misalnya, dan berusaha mengakses BankAccount ini maka aksesnya akan terbatas (tergantung apakah method tersebut di-expose public - dengan huruf kapital atau tidak)


Bila suatu struktur data harus berkolaborasi/berintegrasi dengan type BankACcount ini, maka BankAccount harus didesain supaya bisa diakses (lagi lagi, hanya cukup dengan memberikan huruf kapital pada field yang mau di-expose ke dunia luar).

_Package_ di Go bisa dianggap sebagai unit terkecil dalam hal mendesain sehingga kita akan mendesain aplikasi kita dari dari sisi package, bukan Type (lagi lagi juga, Type tidak bisa disamakan dengan Class meskipun ada sedikit kemiripan).


## Interface

    package payment
    type bankAccount struct {...}
    func(b BankAccount) AccountNumber() string {...}
    func(b BankAccount) AvailableAccount() float32 {...}

perhatikan sekarang bankAccount menggunakan huruf kecil, yang artinya datanya tidak akan bisa diakses dari luar package payment.

Untuk menjadikannya available maka bisa menggunakan interface seperti ini

    type PaymentOption interface {
        AccountNumber() string
        AvailableAccount()  float32
    }

Interface dalam hal ini digunakan untuk menyembunyikan data dari consumer, karena consumer (kode/package lain) tidak bisa mengakses secara langsung type BankAccount.