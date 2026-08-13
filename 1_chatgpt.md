# Ringkasan Belajar JavaScript (Posisi Saat Ini)

## Tujuan

Saya sedang belajar JavaScript dari dasar. Tolong lanjutkan dari progres ini, jangan mulai dari nol lagi.

---

# Materi yang Sudah Dipahami

## 1. Variabel

Sudah paham:

```js
let nama = "Rajendra";
let umur = 18;
```

Perbedaan string dan number juga sudah paham secara dasar.

---

## 2. If Else

Sudah paham:

```js
if (umur < 18) {
    console.log("Belum Dewasa");
} else {
    console.log("Dewasa");
}
```

Sudah paham:

```js
if (...)
else if (...)
else
```

Contoh:

```js
if (umur < 13) {
    ...
}
else if (umur < 18) {
    ...
}
else {
    ...
}
```

Sudah mengerti bahwa:

* else tidak bisa punya kondisi.
* jika ingin kondisi tambahan harus memakai else if.

---

## 3. For Loop

Sudah paham:

```js
for (let i = 1; i <= 5; i++) {
    console.log(i);
}
```

Sudah bisa menebak output loop sederhana.

Sudah memahami:

```js
i++
```

berarti menambah 1 setiap perulangan.

---

## 4. Function

Sudah paham:

```js
function halo() {
    console.log("Halo");
}
```

dan pemanggilan:

```js
halo();
```

---

## 5. Parameter

Sudah paham:

```js
function sapa(nama) {
    console.log("Halo " + nama);
}
```

---

## 6. Return

Sudah paham:

```js
function tambah(a, b) {
    return a + b;
}
```

Sudah memahami bahwa:

```js
console.log(tambah(5, 3));
```

akan mencetak hasil return.

Sudah memahami bahwa:

```js
console.log(tes());
```

akan mencetak nilai yang dikembalikan function.

Sudah memahami bahwa function tanpa return menghasilkan:

```js
undefined
```

jika hasilnya dicetak.

---

## 7. Array

Sudah paham:

```js
let data = ["A", "B", "C"];
```

Sudah memahami:

```js
data[0]
data[1]
data[2]
```

Sudah paham:

```js
data.length
```

menghasilkan jumlah elemen.

Sudah paham:

```js
push()
```

untuk menambah data.

Sudah bisa melakukan loop array:

```js
for (let i = 0; i < data.length; i++) {
    console.log(data[i]);
}
```

---

## 8. Object

Sudah paham object dasar:

```js
let orang = {
    nama: "Rajendra",
    umur: 18
};
```

Mengakses:

```js
orang.nama
orang.umur
```

Mengubah data:

```js
orang.umur = 19;
```

Menambah property:

```js
orang.jurusan = "Informatika";
```

---

## 9. Object Bersarang

Sudah paham:

```js
let mahasiswa = {
    Rajendra: {
        umur: 18
    },

    Budi: {
        umur: 19
    }
};
```

Mengakses:

```js
mahasiswa.Rajendra.umur
```

---

## 10. Dot Notation dan Bracket Notation

Sudah memahami:

```js
mahasiswa.Rajendra
```

sama dengan:

```js
mahasiswa["Rajendra"]
```

Sudah memahami bahwa bracket notation berguna jika menggunakan variabel:

```js
let nama = "Rajendra";

mahasiswa[nama]
```

---

# DOM (Document Object Model)

## 11. Mengambil Elemen HTML

Sudah paham:

```js
document.getElementById()
```

Contoh:

```js
let judul = document.getElementById("judul");
```

---

## 12. Mengubah Isi Elemen

Sudah paham:

```js
judul.textContent = "Halo";
```

---

## 13. Event Listener

Sudah paham:

```js
tombol.addEventListener("click", function() {

});
```

Sudah memahami:

* "click" adalah nama event.
* bukan variabel.
* event lain juga ada seperti:

  * dblclick
  * input
  * keydown

---

## 14. Input User

Sudah paham:

```js
let inputNama = document.getElementById("nama");
```

Mengambil isi:

```js
inputNama.value
```

Sudah memahami bahwa:

```js
.value
```

mengambil teks yang diketik user.

---

## 15. Menampilkan Input ke Halaman

Sudah berhasil membuat mini form:

Input:

* Nama
* Umur
* Jurusan

Ketika tombol ditekan:

```txt
Nama: ...
Umur: ...
Jurusan: ...
```

ditampilkan ke halaman menggunakan:

```js
textContent
```

---

## 16. Validasi Form Dasar

Sudah berhasil membuat validasi:

* Nama tidak boleh kosong.
* Password tidak boleh kosong.
* Minimal 8 karakter.

Menggunakan:

```js
if
else if
```

---

# Pemahaman Penting yang Sudah Didapat

## Kapan Pakai Else If

Gunakan:

```js
if (...)
else if (...)
else
```

Jika hanya ingin SATU hasil.

Contoh:

* Status umur
* Ranking
* Kategori nilai

---

## Kapan Pakai Banyak If

Gunakan:

```js
if (...)
if (...)
if (...)
```

Jika semua syarat harus diperiksa.

Contoh:

* Validasi password
* Validasi form
* Registrasi akun

---

## Return Pada Validasi

Sudah memahami:

```js
if (...) {
    return;
}
```

digunakan untuk menghentikan function lebih awal.

Sudah memahami bahwa:

* return bukan khusus untuk if.
* return bisa digunakan di if maupun else if.
* return digunakan jika function harus berhenti saat itu juga.

Sudah memahami bahwa:

```js
if (...) {
}
else if (...) {
}
else {
}
```

biasanya tidak membutuhkan return karena alurnya sudah terkontrol.

---

# Tingkat Kemampuan Saat Ini

Sudah cukup nyaman dengan:

* Variabel
* If Else
* For Loop
* Function
* Parameter
* Return
* Array
* Object
* DOM dasar
* Input User
* Event Listener
* Validasi Form sederhana

---

# Lanjut Belajar Dari Sini

Jangan ulang materi dasar lagi kecuali saya minta.

Materi berikutnya yang sebaiknya dilanjutkan:

1. DOM lebih lanjut

   * querySelector
   * createElement
   * appendChild
   * remove

2. Mini Project DOM

   * Counter
   * To Do List
   * Kalkulator sederhana

3. Method Array

   * map()
   * filter()
   * find()

4. JavaScript Modern (ES6)

   * template string
   * arrow function
   * destructuring

5. Async JavaScript

   * fetch()
   * API
   * JSON

Tolong lanjutkan dari level ini.
