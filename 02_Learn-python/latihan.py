a = 1       # int (integer/angka) ===> a += a ==> 2
b = "1"     # string/huruf        ===> b += b ==> 11
c = 1.5     # float (desimal)     ===> sama kaya int
d = True    #Boolean


#       ###List
buah = ["apel", "belimbing", "ceri"]       #List, bisa berisi int, str, bolean.
                                                # List digunakan untuk menyimpan banyak data.
                                                #index dimulai dari nol (0).
# buah[1] = "pisang"                            #mengganti belimbing menjadi pisang
# buah[-1] = "pisang"                           #mengganti ceri menjadi pisang

#       ###Tuple
# buah = ("mangga","apel")                  # Tuple tidak bisa dirubah, seperti const
                                            # Apabila diubah = eror


# jumlah = len(buah)                            # len = panjang list, hasilnya 3


# buah.append("durian")                      #append, untuk push data ke list.
# buah.insert(0, "mangga")                   #insert, untuk push data dengan menyisipkan. artinya: masukkan mangga ke index 0 (pertama)
# print(buah)                                #[ 'mangga', 'apel', 'belimbing', 'ceri', 'durian']


# while a <= 5:
#     print(a)
#     a += 1

### remove isi list

# buah = ["apel", "jeruk", "mangga", "jeruk"]

# # 1. Pakai remove (hapus berdasarkan nama)
# buah.remove("jeruk")
# print(buah)  # Hasil: ["apel", "mangga", "jeruk"] (jeruk pertama hilang)

# # 2. Pakai pop (hapus berdasarkan index)
# buah.pop(0)  # Menghapus "apel" di index 0
# print(buah)  # Hasil: ["mangga", "jeruk"]

# # 3. Pakai pop kosongan
# buah.pop()   # Otomatis menghapus yang paling belakang ("jeruk")
# print(buah)  # Hasil: ["mangga"]

# ### remove seluruh isi list
# keranjang = ["buku", "pena"]
# keranjang.clear()
# print(keranjang)  # Hasil: []


