# Unblocker
## Unblock your favorite blocked website

Ini merupakan versi alpha dari Unblocker yang dapat digunakan untuk membuka blocked website di komputer anda. Caranya adalah menjalankan aplikasi as Administrator. Kemudian masukkan alamat website. Lalu klik UNBLOCK IT!.

Selamat, web yang tadi diblokir sudah bisa diakses!

![screenshot](https://raw.githubusercontent.com/almanalfaruq/unblocker/master/screenshot/screenshot.png)

## Penjelasan
Aplikasi ini akan mempersingkat langkah-langkah yang tertulis pada [thread ini di twitter](https://twitter.com/almanalfaruq/status/1232341411237773312?s=20) dengan langsung mengubah file `hosts` yang ada pada komputer anda. Sehingga, apabila anda mengunjungi halaman website tersebut. Anda akan diarahkan ke IP yang seharusnya tanpa melalui routing yang seharusnya.

## Petunjuk Compile
Requirements:

* Golang (version 1.12+)
* [Wails](https://github.com/wailsapp/wails) (Cek cara install-nya juga)

Petunjuk:

1. Clone repo ini
2. Pindah ke direktori hasil clone
3. Jalankan perintah `wails build`. Apabila menggunakan selain Linux, gunakan perintah `wails build -p`.