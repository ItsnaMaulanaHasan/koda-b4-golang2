# Go Slice Manipulation Example

Program Go sederhana ini mendemonstrasikan manipulasi _slice_ dengan menyisipkan elemen baru ke dalam _slice_ yang sudah ada dan kemudian mencetak elemen-elemen dari _slice_ yang telah dimodifikasi.

## Fungsionalitas

Kode ini melakukan langkah-langkah berikut:

1.  Mendefinisikan _slice_ integer awal bernama `scores`.
2.  Membuat _slice_ integer kosong bernama `editedScores`.
3.  Menggabungkan tiga elemen pertama dari `scores` ke dalam `editedScores`.
4.  Menyisipkan nilai integer tunggal (`88`) ke dalam `editedScores`.
5.  Menggabungkan elemen-elemen dari indeks ke-4 (elemen kelima) dan seterusnya dari `scores` ke dalam `editedScores`.
6.  Melakukan iterasi melalui _slice_ `editedScores` yang telah dimodifikasi dan mencetak setiap elemen.

### Detail Slice Manipulation

- **Slice Awal (`scores`):** `[50, 75, 66, 20, 32, 90]`
- **Langkah 1:** `editedScores = append(editedScores, scores[:3]...)`
  - Mengambil `[50, 75, 66]` (indeks 0 hingga sebelum 3)
  - `editedScores` menjadi `[50, 75, 66]`
- **Langkah 2:** `editedScores = append(editedScores, 88)`
  - Menambahkan `88`
  - `editedScores` menjadi `[50, 75, 66, 88]`
- **Langkah 3:** `editedScores = append(editedScores, scores[4:]...)`
  - Mengambil `[32, 90]` (indeks 4 hingga akhir)
  - `editedScores` menjadi `[50, 75, 66, 88, 32, 90]`
- **Langkah 3:** Mengoutputkan editedScores dengan for range
