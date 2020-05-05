---
title: "Code Editor Web : Code Mirror"
date: 2020-04-27T03:10:35+07:00
draft: false
tags: ["editor", "codemirror","monaco"]
categories: ["editor"]
---
Editor (text/code editor) adalah tool utama dari seorang programmer. Berbagai macam code editor baik yang berbayar (seperti produk dari Jetbrain, Sublime, dll) maupun gratis (yang biasanya adalah hasil karya komunitas _open source_ seperti Atom, VS Code, dan editor text yang menolak mati seperti emacs atau vim).

Biasanya, editor ini adalah aplikasi desktop yang harus diinstall di level OS, tapi dengan perkembangan teknologi web, kita bisa menggunakan "engine" dari editor ini untuk kita sematkan di website kita. Bahkan tidak sedikit website yang sudah menggunakan teknik ini untuk sarana pembelajaran pemrograman seperti [codepen](https://codepen.io/) untuk belajar HTML, CSS, Javascript (dan beragam framework yang didukungnya) , atau [repl](https://repl.it) yang menyediakan REPL untuk berbagai
bahasa pemrograman popular saat ini. Tentu saja banyak contoh lainnya, dengan bantuan mesin pencari, kita bisa memasukkan kata kunci : repl, fiddle, online code editor, sandbox online, dsbnya.

Bila kita ingin membangun sesuatu seperti repl atau codepen, tentunya dibutuhkan usaha yang tidak sedikit, tapi mari kita coba menelisik satu bagian kecil saja, yaitu bagian editornya. Codepen menggunakan [code mirror](https://codemirror.net/) sebagai bahan utama website mereka. Bila kita ingin menggunakannya, caranya juga sangat mudah seperti yang dijelaskan di dokumentasi code mirror sendiri.

Kita bisa mendownload file zip dari library ini secara manual atau menggunakan npm untuk mendownload package code mirror

> npm install codemirror

yang akan mendownload library ini ke directory `node_modules`. 

Untuk text editor, cukup membuat sebuah file HTML dan masukkan library yang dibutuhkan (file javascript dari code mirror dan file css untuk stylingnya) di bagian HEAD dokumen HTML.

Sebagai contoh, untuk membuat text editor yang mampu mengenali bahasa pemrograman PHP, maka bagian HEAD akan berisi sebagai berikut bila kamu memakai npm untuk instalasi library code mirror, bila menggunakan zip tinggal mengganti saja path nya ke hasil ekstrasi zip file nya:

    <script src="./node_modules/codemirror/lib/codemirror.js"></script>
    <script src="./node_modules/codemirror/mode/php/php.js"></script>
    <script src="./node_modules/codemirror/mode/htmlmixed/htmlmixed.js"></script>
    <script src="./node_modules/codemirror/mode/xml/xml.js"></script>
    <script src="./node_modules/codemirror/mode/javascript/javascript.js"></script>
    <script src="./node_modules/codemirror/mode/clike/clike.js"></script>

Kemudian di bagian BODY HTML kita akan membuat sebuah elemen TEXTAREA sebagai elemen yang akan digunakan oleh editor kita sebagai berikut :
    
    <textarea id="code" name="code"></textarea>

Lalu di bagian paling bawah body, kita bisa menyisipkan kode Javascript kita di bagian SCRIPT
    
    var myTextArea = document.getElementById("code");
    var editor = CodeMirror.fromTextArea(myTextArea, {
        lineNumbers: true,
        mode: "php",
        indentUnit: 4,
        indentWithTabs: true
    });

Kode di atas akan membuat _instance_ dari code mirror yang kita simpan dalam variabel editor. Bila kita membuka dokumen HTML kita saat ini, harusnya sudah terlihat editor yang bisa digunakan dengan deteksi sintaks-sintaks PHP. 

Untuk mendapatkan nilai dari kode yang kita tuliskan di editor ini sangat mudah, kita bisa menggunakan function `getValue` yang sudah disediakan. Hal ini berguna, misalnya kita akan mengirimkan kode di editor kita ini ke sebuah API dengan XMLRequest/Fetch/Axios sehingga bisa diolah oleh API tersebut. Sebagai contoh, bila kita sudah punya API 


    const url = "http://localhost:3000/exec";
    const btn = document.getElementById("run");
    const snippet = editor.getValue();
    btn.onclick = function() {
    const snippet = editor.getValue();
        fetch(url, {
            method: 'post',
            body: JSON.stringify({
                Code: snippet
            }),
        })
            .then(r => {
                return r.json()
            })
            .then(data => console.log(data))
    }


Kode di atas akan mengirim sebuah request POST ke API dengan alamat yang disimpan di variabel URL dengan kode yang sudah dituliskan ketika elemen `<button>` di click.

Mudah bukan ? Pembaca bisa mengeksplorasi lebih lanjut dengan membuat API sederhana untuk menyimpan dari kode yang dikirim atau mengeksekusinya (dengan pertimbangan keamanan data tentunya).

Alternatif selain code mirror yang bisa dijajal juga adalah [Monaco Editor](https://microsoft.github.io/monaco-editor/) yang merupakan engine yang digunakan untuk membuat VS Code. Lain kali kita akan membahasnya, tapi untuk saat ini kita fokus dulu dengan code mirror, karena proses yang dibutuhkan untuk diimplementasikan sangat jauh lebih mudah.

Jadi text/code editor apa yang kamu gunakan ?




ps : untuk menulis artikel di situs ini, penulis menggunakan [VIM](https://www.vim.org/).

