# Gajah

Tools untuk generate PHP menjadi HTML.

## Cara Pakai

Buka [branch gh-pages](https://github.com/mzaini30/gajah/tree/gh-pages) lalu pilih salah satu `gajah` yang sesuai dengan OS yang kamu gunakan.

Lalu, `gajah` itu kamu masukkan ke PATH di laptop kamu. Misalnya kalau Linux, tempatnya di `/usr/bin/`

Lalu, buka dua Terminal.

Isi Terminal pertama:

```bash
php -S localhost:3000
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