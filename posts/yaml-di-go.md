---
title: "Yaml Di Go"
date: 2020-05-09T03:46:55+07:00
draft: false
tags: ["go", "yaml", "hugo"]
categories: ["go"]
---
Yet Ain't Markup Language atau biasa disebut [yaml](https://yaml.org/) adalah sebuah standard untuk melakukan [_serialization_](https://en.wikipedia.org/wiki/Serialization) yang berorientasi object, yang diklaim lebih mudah dibaca karena mengikuti gaya indentation ala [python](https://www.python.org/).

Aplikasi-aplikasi cloud kekinian, seperti docker, kubernetes, atau bahkan aplikasi berbasis web seperti hugo bisa menggunakan format ini, biasanya sebagai bagian dari menyimpan konfigurasi seperti config.yaml, docker-compose.yml, atau yang lainnya (ekstensi dari file bisa .yaml saja atau .yml).

Untuk menggunakan file/format yaml, tentunya kita perlu yaml parser, untuk bahasa pemrograman Go, ada dua parser yang direkomendasikan dari website resmi yaml yaitu : [yaml.v2](https://github.com/go-yaml/yaml) atau [go-gypsy](https://github.com/kylelemons/go-gypsy) yang sudah lama tidak diupdate maintainer projectnya. Jadi pilihan yang masuk akal bila kita menggunakan yaml.v2 dalam project kita.

Untuk menggunakannya kita import saja package-nya 

    import (
        "gopkg.in/yaml.v2"
    )

kita menggunakan variabel yang kita deklarasikan secara langsung seperti ini (untuk mempermudah)

    var data = `
    title: yaml di go!
    post:
    id: 2
    tags: [go, yaml]
    `

Nah, data di atas sudah memenuhi kualifikasi sebagai standar yaml, sehingga kita bisa mencoba mem-parsingnya. Tapi sebelum itu siapkan dulu type berupa struct untuk menampung hasil parsingnya nanti :

    type Content struct {
        Title string 
        Post
    }

    // Post represent blog post
    type Post struct {
        ID string `yaml: "id"`
        Tags [] string `yaml: "tags"`
    }

Tentu saja kita bisa menggunakan satu struct saja supaya lebih sederhana :). 

Kemudian kita coba untuk memparsing dan men-decode yaml format di atas dengan seperti ini

	// From YAML to Struct
	c := &Content{}
	// unmarshal/decode yaml file to struct 
	err := yaml.Unmarshal([]byte(data), c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	
Bila kita ingin menampilkan hasilnya 

    fmt.Printf("--- content struct:\n%v\n\n", c)

Kita juga bisa membalikkan operasinya, misalnya dari type struct Content kita ke format yaml dengan seperti ini :

	d, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

Hasilnya bila kita print 

	fmt.Printf("--- yaml:\n%s\n\n", d)

Okay, sekarang semuanya jadi lebih karena kita bisa mempunyai keleluasaan menentukan apakah format yaml itu ada dalam sebuah variabel atau disimpan di file terpisah seperti misalnya config.yaml. Adalah pekerjaan rumah buat pembaca untuk mencoba mengeksplorasi membaca yaml dari sebuah file.

hint : Cara paling sederhana adalah dengan menggunakan package :  ioutil. 