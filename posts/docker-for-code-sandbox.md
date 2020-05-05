---
title: "Docker for Code Sandbox : Proof of Concept"
date: 2020-04-23T02:41:15+07:00
draft: false
tags: ["php","docker", "sandbox"]
categories: ["docker"]
---

Hari ini mencoba bikin sandbox env, dimana dalam sandbox ini kita bisa meng-compile/eksekusi code yang diberikan ke stdin dari sandbox tersebut. 

Dalam percobaan ini, digunakan docker container sebagai sandbox environment. Percobaan pertama adalah membuat sandbox untuk mengeksekusi kode php.

Persiapan yang perlu dilakukan.
1. Download docker di [sini](https://docs.docker.com/docker-for-windows/install/).
2. Pastikan docker server sudah running di local machine. Bisa dites dengan mengetikkan command `docker` di terminal.
3. Pull image php (bisa pilih versi alpine supaya sizenya tidak terlalu besar, versi terbaru 7.4 sudah available).


Okay, setelah semua beres, kita bisa melakukan tes sederhana, sekedar menjadi _proof of concept_ bahwa kita bisa melakukan eksekusi code di docker.

1. Membuat container dari image yang sudah dipull di tahap persiapan di atas. Supaya memudahkan interaksi antara host machine dengan container, kita akan menyiapkan volume di container yang akan mount ke file system host machine.
    
    > docker create -v /home/codesnippet:/app --name phpcontainer php:7.4-alpine

2. Kita coba running containernya, dimana kita akan mengeksekusi file test.php (yang ada di host machine yang telah di mount). Perhatikan bahwa kita menambahkan beberapa flag  seperti flag -w yang untuk mengeset _working directory_ dari container dan flag --rm supaya container akan segera dihapus setelah selesai.

    > docker run -w /app --rm --volumes-from phpcontainer php:7.4-alpine php test.php


untuk test yang lebih sederhana kita bisa juga melakukan ini :

    echo "<?php echo 'asdf';" | docker run -i --rm php:7.4-alpine php
