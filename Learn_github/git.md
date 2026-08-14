# 📑 Kamus Perintah Git & GitHub (Cheat Sheet)

---

## 1. Persiapan & Identitas
*Dipakai sekali saja di awal untuk mengenalkan laptop ke akun GitHub.*
* `git config --global user.name "Nama Kamu"` ➜ Mendaftarkan nama lengkap.
* `git config --global user.email "Email Kamu"` ➜ Mendaftarkan email yang terhubung ke GitHub.
* `git config user.name` ➜ Mengecek nama yang saat ini terdaftar.
* `git config user.email` ➜ Mengecek email yang saat ini terdaftar.
* `git init` ➜ Mengaktifkan Git di dalam folder baru (membuat folder pembungkus).
* `git branch -m main` ➜ Mengubah nama branch utama bawaan dari `master` menjadi `main`.

---

## 2. Alur Kerja Harian (Paling Sering Dipakai)
*Alur standar dari mengambil kode, mengedit, hingga mengirim ke awan.*
* `git clone <link-url-github>` ➜ Men-download repositori dari web GitHub ke laptop.
* `git status` ➜ Mengecek file apa saja yang baru diedit (warna merah/hijau).
* `git add .` ➜ Membungkus/menandai seluruh perubahan file agar siap dikunci.
* `git commit -m "pesan commit"` ➜ Membuat *savepoint* (titik simpan) resmi di memori laptop.
* `git push origin main` ➜ Mengirim hasil *savepoint* laptop ke web GitHub.
* `git pull origin main` ➜ Menyedot update terbaru dari web GitHub ke laptop (biar seimbang).

---

## 3. Eksperimen Aman (Branching & Merging)
*Dipakai saat ingin membuat fitur baru tanpa takut merusak kode utama.*
* `git branch` ➜ Mengecek posisi kamu sedang berdiri di cabang (*branch*) mana.
* `git checkout -b <nama-branch>` ➜ Membuat cabang baru dan langsung melompat ke dalamnya.
* `git checkout main` ➜ Melompat kembali ke dimensi/cabang utama.
* `git merge <nama-branch>` ➜ Menyatukan kode dari cabang eksperimen ke cabang utama (`main`).
* `git branch -d <nama-branch>` ➜ Menghapus cabang bekas yang sudah tidak terpakai agar rapi.

---

## 4. Merapikan File & Folder (Manajemen Struktur)
*Siasat memindahkan, menghapus, atau mengganti nama file lewat terminal.*
* `mkdir <nama-folder>` ➜ Membuat folder baru di laptop.
* `git mv <nama-file> <folder-tujuan/>` ➜ Memindahkan file ke dalam folder tujuan.
* `git mv <nama-lama> <nama-baru>` ➜ Mengubah nama (*rename*) file atau folder.
* `git rm <nama-file>` ➜ Menghapus file secara total (hilang di laptop dan di GitHub).
* `git rm --cached <nama-file>` ➜ Menghapus file di web GitHub saja (di laptop aslinya tetap ada).

---

## 5. Tombol Darurat & Paksaan (Mesin Waktu)
*Dipakai jika kodingan tabrakan (conflict) atau terjadi salah langkah.*
* `git merge --abort` ➜ Membatalkan proses tabrakan kode (*merge conflict*) yang sedang menggantung.
* `git reset --hard HEAD~1` ➜ Mundur total 1 langkah ke belakang (menghapus commit terakhir yang salah).
* `git push origin main --force` ➜ Memaksa web GitHub mengikuti dan meniru 100% struktur yang ada di laptop saat ini.