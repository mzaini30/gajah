# Gajah

<p align="center">
 <img src="https://i.postimg.cc/DZpd1HrW/ade1270e9f707db76889d3d9dc628e5a.jpg">
</p>

Tools untuk generate PHP menjadi HTML.

## Struktur Folder

```
.
├── build
│   ├── hello.html
│   ├── index.html
└── src
    ├── hello.php
    ├── index.php

```

Jadi, kita meletakkan semua file PHP dan file assets di folder `src`.

## Cara Pakai

Buka [branch gh-pages](https://github.com/mzaini30/gajah/tree/gh-pages) lalu pilih salah satu `gajah` yang sesuai dengan OS yang kamu gunakan.

Lalu, `gajah` itu kamu masukkan ke PATH di laptop kamu. Misalnya kalau Linux, tempatnya di `/usr/bin/`

Lalu, buka dua Terminal pada folder projectmu.

Isi Terminal pertama:

```bash
php -S localhost:3000 -t src
```

Isi Terminal kedua:

```bash
gajah 3000
```

Maka, otomatis semua file PHP akan digenerate menjadi HTML.

## Minify

```bash
gajah 3000 minify
```

## Donate

<p align='center'>
    <a href='https://www.nihbuatjajan.com/mzaini30'>
        <img src='https://d4xyvrfd64gfm.cloudfront.net/buttons/default-cta.png'/>
    </a>
</p>