---
title: "OOP : Inheritance Composition di Go"
date: 2020-05-04T13:16:17+07:00
draft: false
tags: ["go", "oop", "inheritance", "composition"]
categories: ["go"]
---

Bagaimana caranya untuk _reuse_ dari _behavior_ sebuah object yang sudah ada ? secara umum ada dua hal yang bisa dilakukan :
- Inheritance
- Composition

Meskipun begitu di Go, memilih _composition over inheritance_, meski begitu mari kita singgung sedikit tentang _inheritance_ supaya kita mendapatkan gambaran utuhnya.

Inheritance adalah strategi _reuse_ dimana sebuah type adalah berdasarkan type yang lainnya (base type), sehingga meng-_inherit_ atau mewarisi fungsionalitas dari base type tersebut. Seringkali digambarkan sebagai hubungan _parent-child_.

    Account 
        AvailableFund()
        ProcessPayment()
    CreditAccount
        RequestCreditIncrease()


Bila kita ingin supaya CreditAccount meng-inherit dua behavior yaitu` AvailableFund()` dan ProcessPayment dari Account, kita bisa melakukannya dengan meng-extend Account dari CreditAccount.

      Account (parent)
         |
        /|\ (extends)
         |
    CreditAccount (child)

Hal ini punya tantangan sendiri yaitu _parent-child_ ini sifatnya _tightly coupled_ (berkaitan erat), bahwa CreditCard (child) harus tahu detil tentang Account(parent), dan mengganti sesuatu di Account(parent) menjadi sangat sulit karena akan ada akibatnya bagi CreditCard.

Selain itu juga akan lebih susah di-debug dan maintain, karena inheritance ini bisa beberapa tingkat dan biasanya di sisi child akan meng-overwrite behavior dari parent, maka akan lebih sulit menentukan saat debug darimana behavior ini berasal. 

Selain itu child akan mewarisi semua behavior dari parent, baik dibutuhkan maupun tidak, ada yang menyebut juga masalah ini sebagai problem [gorilla-banana-jungle](https://softwareengineering.stackexchange.com/questions/368797/sample-code-to-explain-banana-monkey-jungle-problem-by-joe-armstrong)  pada pemgrograman berorientasi object.


Go tidak mempunyai fitur yang mendukung inheritance secara default, tapi memilih menggunakan Composition.

#### Composition

Composition adalah strategi _reuse_ dimana sebuah type memiliki object-object yang mempunyai fungsionalitas/behavior yang berbeda-beda. Type tersebut sehingga bisa mendelegasikan fungsi tertentu ke object tersebut sesuai dengan behavior yang dipunyainya.

    CreditAccount
        AvailableFund()
        ProcessPayment()
        Account
            AvailableFund()
            ProcessPayment()

Ketika `AvailableFund() `dipanggil pada CreditAccount dia akan mendelegasikan ke AvailableFund() yang ada pada object internal Account. Proses yang sama terjadi ketika ProcessPayment dipanggil CreditAccount. Hal ini berarrti terjadi sebuah proses delegasi dari CreditAccount ke object Account seperti definisi _composition_.


#### Type Embedding
Type embedding adalah pendekatan di Go dalam rangka menyediakan _composition_ atau delegasi. Contohnya seperti di bawah ini:

    type Account struct {...}
    func (a *Account) AvailableFund() float32 {...}
    func (a *Account) ProcessPayment(amount float32) bool {...}


    type CreditAccount struct {
        Account
    }

Kita lihat bahwa type account tidak mempunyai nama field pada type Account, Compiler akan mengenali dan menyediakan proses delegasi dari CreditAccount ke object Account. Hasilnya kita bisa melakukan ini:

    ca := & CreditAccount {...}
    funds := ca.AvailableFund()


Yang membedakan  _composition_ dengan _inheritance_ adalah bahwa CreditAccount and Account dalam hal ini tidak _interchangeable_. 

Ingat, pada bahasa pemgrograman berorientasi object lainnya, bila CreditCard meng-extend Account, maka keduanya saling bisa ditukarkan. Contohnya :  instance CreditAccount (new CreditAccount) bisa disimpan sebagai Account atau sebaliknya, instance dari Account (new Account) disimpan sebagai CreditAccount. Hal ini juga disebut sebagai _polymorphism_ yang akan kita bahas selanjutnya.