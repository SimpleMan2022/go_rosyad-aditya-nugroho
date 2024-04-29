**Pertanyaan:**
Terdapat sekumpulan data mengenai tulisan dalam bentuk tweet mengenai sebuah kebijakan. Sekumpulan data tersebut ingin dikelompokkan berdasarkan sentimen dari tweet tersebut yaitu sentimen positif dan negatif. Jelaskan algoritma A.I. yang dapat digunakan untuk mengelompokkan tweet tersebut beserta alasannya.

**Jawaban:**
Untuk mengelompokkan tweet berdasarkan sentimen (positif dan negatif) dari sekumpulan data tweet mengenai sebuah kebijakan, terdapat beberapa algoritma AI yang dapat digunakan, antara lain:

**Naive Bayes**

Algoritma Naive Bayes adalah algoritma klasifikasi probabilistik yang sederhana dan mudah diimplementasikan. Algoritma ini bekerja dengan menghitung probabilitas setiap tweet termasuk dalam kategori positif atau negatif berdasarkan kata-kata yang terdapat dalam tweet tersebut.

Cara kerjanya adalah algoritma ini berdasarkan pada teorema Bayes yang mengasumsikan bahwa fitur-fitur (kata-kata dalam kasus klasifikasi teks) adalah independen satu sama lain. Walaupun asumsi ini seringkali tidak benar dalam konteks nyata, Namun, algoritma ini sering memberikan hasil yang baik dalam praktek.

Kelebihannya seperti cepat dalam pelatihan dan prediksi, terutama untuk dataset besar, membutuhkan jumlah data yang relatif kecil untuk pelatihan, dan cocok untuk klasifikasi teks.

Alasan pengunaanya karena efisiensinya dan kecenderungan memberikan hasil yang memadai untuk klasifikasi teks, Naive Bayes adalah pilihan yang populer untuk tugas ini.

Langkah langkah algoritma Naive Bayes untuk analisis sentimen tweet:

**1. Persiapan Data:**

- Kumpulkan dan bersihkan data tweet.
- Preprocessing data (tokenisasi, stemming/lemmatization).
- Bagi data menjadi set pelatihan dan pengujian.

**2. Pelatihan Model:**

- Buat representasi vektor (BOW, TF-IDF, dll.) untuk tweet.
- Hitung probabilitas kata dalam tweet positif/negatif dan probabilitas apriori kategori.
- Klasifikasikan tweet dengan teorema Bayes (probabilitas posterior tertinggi).

**3. Evaluasi Model:**

- Hitung metrik kinerja (akurasi, presisi, recall, F1-score) pada set pengujian.
- Analisis hasil untuk mengetahui performa model.
