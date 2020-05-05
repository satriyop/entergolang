---
title: "OOP : Polimorphism di Go"
date: 2020-05-05T16:35:29+07:00
draft: false
tags: ["go", "oop", "polimorphism"]
categories: ["go"]
---

Bila kita mempunyai _package_ dengan beberapa fungsionalitas/service yang sudah didefinisikan sedangkan kita mempunyai satu type atau beberapa type yang akan menyediakan fungsionalitas tersebut ke consumer (kode yang akan berkomunikasi dengan Package/Type kita tersebut) maka desain yang baik adalah consumer tidak tahu, type mana yang menyediakan fungsionalitas/service yang diconsume-nya. Consumer hanya tahu bahwa mereka berinteraksi secara _polimorfik_ dengan sebuah type, dan fungsionalitas/service yang akan disediakan oleh type tersebut, tapi implementasi detilnya tidak akan kelihatan. Hal ini disebut _polimorphism_.

Interface **Reader** di pustaka standard mempunyai satu method yaitu

    Read([]byte) (int, error)

Tapi kita bisa melakukan Read dari hal-hal bermacam seperti dari :
- File
- TCP
- WebSocket
Bila Ketiganya mempunyai method _Read_ juga, artinya mereka mengimplementasikan interface _Reader_.

Maka Bila bagian kode kita ingin menconsume sebuah _Reader_ maka kita bisa menggunakan object dari File, TCP, atau WebSocket, karena ketiganya mengimplementasikan interface _Reader_.

Di Go, kita **hanya** menggunakan Interface sebagai cara melakukan polimorphism atau bila kita memakai contoh seperti artikel sebelumnya maka akan seperti ini:

    type PaymentOption interface {
        ProcessPayment(float32) bool
    }

    type CreditAccount struct {...}
    func (c *CreditAccount) ProcessPayment (amount float32) bool {...}

    type CashAccount struct {...}
    func (c *CashAccount) ProcessPayment (amount float32) bool {...}

Di kode di atas terlihat bahwa type CreditAccount maupun CashAccount mengimplementasikan interface PaymentOption. Sehingga bila kita ingin berkomunikasi di antara keduanya kita bisa menggunakan cukup meng-consume interface PaymentOption saja. 

    func pay(a PaymentOption) {
        paid := a.ProcessPayment(...)
    }

Tentu saja kta bisa melakukan _Type Assertion_ bila ingin berkomunikasi dengan nilai konkret dari interface tersebut.