---
title: "Go Docker Sandbox"
date: 2020-04-24T02:52:31+07:00
draft: false
tags: ["docker", "go","php"]
categories: ["go"]
---

Melanjutkan Proof of Concept dari artikel sebelumnya. Kali ini kita coba memanggil process dari pemanggilan container docker dengan menggunakan Go. 

Fungsi yang akan menjalankan container secara otomatis akan seperti ini :


    func startDocker() {
	    fmt.Println("going to run code in docker, please wait...")
	    cmd := exec.Command("bash", "-c", "docker run -w /app --rm --volumes-from phpcontainer php:7.4-alpine php test.php")
	    var out bytes.Buffer
	    cmd.Stdout = &out
	    err := cmd.Run()
	    check(err)
	    fmt.Println(out.String())
    }


Di sini kita menggunakan exec, supaya bisa memanggil proses eksternal (Docker dalam hal ini) dan menyambungkan stdout dari container ke dalam Go dan mendisplaynya di terminal.

Ternyata tidak sesulit yang dibayangkan bukan ?

Tentunya func di atas harus ditambahkan juga beberapa validasi, seperti apakah docker executable sudah ada di system atau belum, apakah container bisa di run atau tidak dsbnya, tapi untuk tujuan _proof of concept_ sekiranya ini cukup membuktikan bahwa kita bisa menjalankan _arbitrary code_ dalam sebuah container yang disposable dan relatif aman untuk _host machine_.

Referensi untuk package  _exec_ sendiri merupakan standard library dari Go dan  bisa dipelajari lebih lanjut di [sini](https://golang.org/pkg/os/exec).
