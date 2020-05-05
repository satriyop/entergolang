---
title: "OOP : Message Passing di Go"
date: 2020-05-02T16:55:13+07:00
draft: false
tags: ["go", "oop", "message passing"]
categories: ["go"]
---

Definisi dari secara singkat dari _Message Passing_ adalah mengirimkan message ke object, dan menyerahkan kendalinya ke object tersebut. 
Dalam pemgrograman berorientasi Object hal ini bisa diartikan juga, bahwa consumer hanya bisa melakukan request terhadap sebuah service, tapi bagaimana service tersebut diimplementasikan akan diserahkan kepada object itu sendiri.

Pendektatan yang bisa digunakan dalam melakukan _message passing_ adalah dengan menggunakan :
1. Interface
2. Channel


### Interface

    type PaymentOption interface {
        ProcessPayment(float32) bool
    }


    type CashAccount struct {}
    func (c *CashAccount) ProcessPayment(amount float32) bool {...}

    type CreditAccount struct {}
    func (c *CreditAccount) ProcessPayment(amount float32) bool {...}

Bila kita melakukan ini

    ca := &CashAccount{}
    ca.ProcessPayment(1000)

ini berarti kita _direct invocation_ yang mana bukan message passing.

Kita bisa melakukan message passing seperti ini 

    var paymentOption PaymentOption
    pO = &CashAccount{}

    ok := pO.ProcessPayment(1000)

bedanya adalah ketika kita menggunakan ca.ProcessPayment berarti kita sudah menspesifikasikan servicenya, tapi dengan menggunakan pO.ProcessPayment kita melakukan pemanggilan dari interface, jadi kita tidak tahu service apa yang akan dipanggil di CashAccount (tapi kita cukup yakin dan tahu bahwa ada service tersebut karena CashAccount mengimplementasikan interface paymentOption).


### Channel

    type CreditAccount struct {}
    func (c *CreditAccount) processPayment(amount float32) bool {...}

harap diingat bahwa CreditAccount ini ada dalam scope package (huruf kecil pada processPayment), maka package lain(bisa dikatakan consumer dari package ini) tidak bisa mengakses method ini.

lalu kita membuat fungsi constructor (yang diekspos ke package lain dengan huruf kapital) seperti ini

    func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
        ca := &CreditAccount{}
        
        go func(chargeCh chan float32) {
            for amount := range chargeCh {
                ca.processPayment(amount)
            }
        }(chargeCh)

        return ca
    }


Kita membuat go routine yang akan menerima channel yang akan diterima . Di dalam go routine ini kita akan mendengarkan message yang ada dalam channel chargeCh, ketika message diterima, maka akan mentrigger invokasi processPayment dengan amount yang ada dalam message tersebut.

Dalam hal ini kita menggunakan channel untuk consumer (package lain misalnya) untuk mempunyai akses ke CreditAccount dan memanggil method / service.

Dari sisi consumer/package lain akan seperti ini

    chargeCh = make(chan float32)
    account := CreateCreditAccount(chargeCh)
    chargeCh <- 1000

Kita akan memberikan mengirimkan message ke channel chargeCh dan penerimanya yaitu CreditAccount yang akan mengambil alih proses selanjutnya.

Hal ini memberikan abstraksi yang terbaik antara message yang dikirimkan dan service/method yang dipanggil. Karena dari sisi customer tidak mengetahui bahwa ada method/service processPayment di sisi CreditAccount.