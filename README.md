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

## Pengecualian Convert PHP

Pada dasarnya, dia akan meng-convert semua teks `php` menjadi `html` pada file-file `php`. Contohnya seperti ini:

```html
<a href="about.php">menuju about</a>
```

Maka, dia akan berubah menjadi:

```html
<a href="about.html">menuju about</a>
```

Untuk mencegahnya, kita buat file `ignore.txt` yang berisi dengan teks-teks yang nggak boleh diconvert jadi `.html`. Contoh isinya:

```
satu.php
about.php?hai
```

## File dan Folder yang Dikecualikan

Saat menyalin semua file dan folder dari `src/` ke `build/`, ada beberapa file dan folder yang dikecualikan, yaitu:

- .gitignore
- vendor
- composer.json
- composer.lock
- node_modules
- package.json
- pnpm-lock.yaml

## Donate

<p align='center'>
    <a href='https://www.nihbuatjajan.com/mzaini30'>
        <img src='https://d4xyvrfd64gfm.cloudfront.net/buttons/default-cta.png'/>
    </a>
</p>