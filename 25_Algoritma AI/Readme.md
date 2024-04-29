# Algoritma A.I :rocket:

## Narrow AI, Reactve Machine dan Limited Memory

Narrow AI, juga dikenal sebagai Weak AI, adalah jenis kecerdasan buatan yang terbatas pada tugas tertentu atau domain spesifik. Jenis AI ini dirancang untuk menyelesaikan masalah atau tugas yang terbatas, seperti pengenalan suara, pengenalan wajah, atau penerjemahan bahasa.

Reactive Machine, merupakan tipe AI yang tidak memiliki kemampuan memori jangka panjang dan tidak dapat belajar dari pengalaman sebelumnya. AI ini hanya merespons pada situasi yang ada di hadapannya tanpa memiliki pemahaman konteks yang lebih luas.

Limited Memory, merupakan perluasan dari tipe Reactive Machine, di mana AI dapat menyimpan informasi dalam waktu yang terbatas. Ini memungkinkan AI untuk membuat keputusan yang lebih kompleks dengan mempertimbangkan beberapa kejadian sebelumnya, meskipun tidak dapat belajar dari pengalaman dalam arti yang sama seperti AI yang lebih canggih.

## Machine Learning

Machine learning (ML) adalah cabang dari kecerdasan buatan yang mengkaji teknik-teknik yang memungkinkan komputer untuk belajar dari data dan membuat keputusan atau prediksi tanpa harus diprogram secara eksplisit. Dengan menggunakan algoritma dan model statistik, mesin dapat belajar dari pola-pola dalam data untuk membuat keputusan atau melakukan tugas tertentu.

Taxonomi Machine Learning adalah penggolongan berbagai pendekatan dan teknik dalam machine learning. Taxonomi ini dapat berbeda-beda tergantung pada perspektif dan penekanan masing-masing, tetapi secara umum terdapat beberapa kategori utama, antara lain:

1. Supervised Learning (Pembelajaran Terawasi): Model belajar dari data yang berpasangan dengan label. Tujuannya adalah untuk mempelajari hubungan antara input dan output yang diketahui.

2. Unsupervised Learning (Pembelajaran Tanpa Pengawasan): Model belajar dari data tanpa label. Tujuannya adalah untuk menemukan struktur dalam data, seperti pola-pola atau kelompok-kelompok yang tersembunyi.

3. Reinforcement Learning (Pembelajaran Penguatan): Model belajar melalui interaksi dengan lingkungannya. Tujuannya adalah untuk memaksimalkan reward yang diterima dari lingkungan tersebut.

4. Semi-Supervised Learning (Pembelajaran Semi-Terawasi): Kombinasi antara supervised dan unsupervised learning, di mana model menggunakan data yang sebagian besar tidak berlabel bersama dengan data yang berlabel untuk pembelajaran.

## Supervised Learning

Supervised learning adalah jenis pembelajaran mesin di mana model belajar dari data yang berpasangan dengan label. Dalam supervised learning, tujuan utama adalah untuk mempelajari hubungan antara input (fitur) dan output (label) yang sudah diketahui. Ada dua tugas utama dalam supervised learning, yaitu classification (klasifikasi) dan regression.

1. Classification (Klasifikasi):

- Tujuan: Memprediksi kelas atau label diskrit untuk data input yang baru.
- Contoh: Memprediksi apakah email adalah spam atau bukan, klasifikasi gambar menjadi kategori tertentu seperti kucing atau anjing, klasifikasi pelanggan menjadi kategori risiko kredit tertentu.
- Algoritma yang umum digunakan: Decision Trees, Random Forest, Support Vector Machines (SVM), Logistic Regression, Neural Networks.

2. Regression:

- Tujuan: Memprediksi nilai kontinu berdasarkan input yang diberikan.
- Contoh: Memprediksi harga rumah berdasarkan fitur-fitur tertentu seperti luas tanah, jumlah kamar, dll., memprediksi pendapatan seseorang berdasarkan pendidikan dan pengalaman kerja, memprediksi suhu berdasarkan waktu dan lokasi.
- Algoritma yang umum digunakan: Linear Regression, Polynomial Regression, Support Vector Regression (SVR), Random Forest Regression, Neural Networks.

## Unsupervised Learning

Unsupervised learning adalah jenis pembelajaran mesin di mana model belajar dari data yang tidak memiliki label. Tujuannya adalah untuk menemukan struktur yang tersembunyi dalam data tanpa bimbingan eksternal. Beberapa tugas utama dalam unsupervised learning meliputi clustering, dimensionality reduction, association rule learning, dan anomaly detection.

1 Clustering (Pengelompokan):

- Tujuan: Mengelompokkan data ke dalam kelompok-kelompok yang serupa berdasarkan pola yang ada dalam data.
- Contoh: Mengelompokkan pelanggan berdasarkan perilaku pembelian mereka, mengelompokkan dokumen berita ke dalam topik-topik tertentu, mengelompokkan data geospasial berdasarkan pola spasialnya.
- Algoritma yang umum digunakan: K-Means, Hierarchical Clustering, DBSCAN, Gaussian Mixture Models (GMM).

2. Dimensionality Reduction (Reduksi Dimensi):

- Tujuan: Mengurangi jumlah fitur (dimensi) dalam data, tetapi tetap mempertahankan informasi yang penting.
- Contoh: Memproyeksikan data ke ruang yang lebih rendah untuk mempermudah visualisasi, mengurangi noise dan redundansi dalam data, mempercepat waktu komputasi untuk tugas tertentu.
- Algoritma yang umum digunakan: Principal Component Analysis (PCA), t-Distributed Stochastic Neighbor Embedding (t-SNE), Linear Discriminant Analysis (LDA).

3. Association Rule Learning (Pembelajaran Aturan Asosiasi):

- Tujuan: Menemukan hubungan atau ketergantungan antara item-item dalam kumpulan data.
- Contoh: Menemukan asosiasi antara barang yang dibeli secara bersamaan dalam transaksi, menemukan pola pembelian dalam data e-commerce, menemukan hubungan antara gejala dan penyakit dalam data medis.
- Algoritma yang umum digunakan: Apriori algorithm, Eclat algorithm.

4. Anomaly Detection (Deteksi Anomali):

- Tujuan: Mengidentifikasi pola yang tidak biasa atau anomali dalam data yang mungkin menunjukkan kejadian langka, kesalahan, atau perilaku mencurigakan.
- Contoh: Mendeteksi transaksi kartu kredit yang mencurigakan, mendeteksi kegagalan dalam sistem manufaktur, mendeteksi serangan siber.
- Algoritma yang umum digunakan: Isolation Forest, One-Class SVM, K-Means.

## Semi-Supervised Leaning

Semi-supervised learning adalah jenis pembelajaran mesin di mana model belajar dari kombinasi data yang berlabel (data yang memiliki label) dan data yang tidak berlabel (data yang tidak memiliki label). Tujuannya adalah untuk memanfaatkan informasi yang terdapat dalam data yang tidak berlabel untuk meningkatkan kinerja model.

Contoh-contoh algoritma semi-supervised learning yang umum digunakan termasuk:

1. SVM (Support Vector Machine): SVM dapat digunakan dalam framework semi-supervised learning dengan memanfaatkan teknik seperti Self-Training atau Co-Training. Dalam Self-Training, model SVM dilatih dengan data berlabel, dan kemudian digunakan untuk memprediksi label data yang tidak berlabel. Data yang diprediksi dengan tingkat kepercayaan tertentu kemudian ditambahkan ke data berlabel untuk melatih ulang model. Dalam Co-Training, dua model SVM dilatih secara independen menggunakan dua set fitur yang saling eksklusif, dan kemudian saling bertukar prediksi mereka untuk meningkatkan akurasi.

2. K-Nearest Neighbors (KNN): Dalam semi-supervised learning dengan KNN, data yang tidak berlabel dapat digunakan untuk memperluas himpunan tetangga yang digunakan oleh algoritma untuk klasifikasi. Hal ini dapat membantu meningkatkan ketepatan prediksi dengan memanfaatkan informasi dari data yang tidak berlabel.

3. Decision Tree: Dalam semi-supervised learning dengan decision tree, informasi dari data yang tidak berlabel dapat digunakan untuk memperbaiki struktur pohon keputusan. Metode seperti Co-Training dapat digunakan di sini juga, di mana dua pohon keputusan dilatih secara independen dengan subset fitur yang berbeda, dan kemudian digunakan untuk memperbaiki prediksi satu sama lain.

## Reinforcement Learning

Reinforcement learning (RL) adalah jenis pembelajaran mesin di mana agen belajar memilih tindakan dalam lingkungan tertentu untuk mencapai tujuan tertentu. Tujuan agen adalah untuk memaksimalkan kumulatif reward yang diterima dari lingkungan seiring waktu. RL cocok untuk masalah di mana tindakan yang diambil secara berurutan memengaruhi lingkungan, dan di mana keputusan yang optimal harus dipelajari melalui pengalaman.

Contoh-contoh algoritma reinforcement learning termasuk:

1. Dyna: Dyna adalah sebuah pendekatan untuk reinforcement learning yang menggabungkan pembelajaran langsung dari pengalaman dengan perencanaan berbasis model. Ini berarti agen tidak hanya belajar dari interaksi langsung dengan lingkungan tetapi juga membangun model lingkungan dan menggunakan model tersebut untuk melakukan perencanaan untuk memperbarui kebijakan atau nilai-nilai yang diperkirakan.

2. Policy Gradient: Algoritma ini berfokus pada pembelajaran langsung kebijakan (policy) optimal tanpa memerlukan perencanaan atau model lingkungan. Dengan menggunakan pendekatan gradient descent, agen belajar untuk memperbarui kebijakan dengan mengikuti arah yang meningkatkan reward yang diharapkan.

3. ADP (Approximate Dynamic Programming): ADP adalah pendekatan yang menggabungkan konsep pembelajaran mesin, optimasi, dan kontrol untuk menyelesaikan masalah pengendalian optimal dalam kasus-kasus di mana model lingkungan tidak diketahui atau tidak dapat diandalkan. ADP sering digunakan dalam kasus di mana pemodelan lingkungan yang tepat sulit atau mahal.

## Deep Learning

Deep learning adalah cabang dari machine learning yang menggunakan arsitektur jaringan saraf tiruan (neural networks) yang dalam (deep) untuk memodelkan dan memahami data yang kompleks. Deep learning telah menghasilkan kemajuan signifikan dalam berbagai bidang seperti pengenalan gambar dan suara, pemrosesan bahasa alami, dan pengenalan pola.

Beberapa konsep utama dalam deep learning meliputi:

1. Neural Networks: Deep learning menggunakan jaringan saraf tiruan yang terinspirasi dari struktur dan fungsi otak manusia. Jaringan ini terdiri dari lapisan-lapisan neuron (unit pengolahan informasi) yang saling terhubung.
2. Deep Neural Networks (DNN): DNN adalah jaringan saraf tiruan dengan beberapa lapisan (biasanya lebih dari tiga) yang memungkinkan model untuk mempelajari representasi yang semakin abstrak dan kompleks dari data.
3. Convolutional Neural Networks (CNN): CNN adalah jenis jaringan saraf tiruan yang sangat efektif untuk tugas-tugas pengolahan gambar. Mereka menggunakan lapisan konvolusi untuk mengekstrak fitur-fitur penting dari gambar.
4. Recurrent Neural Networks (RNN): RNN adalah jenis jaringan saraf tiruan yang efektif untuk pemrosesan data berurutan seperti teks atau suara. Mereka memiliki kemampuan untuk "mengingat" informasi sebelumnya dalam proses pembelajaran.
5. Long Short-Term Memory (LSTM): LSTM adalah jenis RNN yang dirancang untuk mengatasi masalah "vanishing gradient" yang sering terjadi dalam pembelajaran jangka panjang. LSTM memiliki mekanisme "gate" yang memungkinkannya untuk menjaga informasi dalam jangka waktu yang lama.
6. Generative Adversarial Networks (GAN): GAN adalah jenis arsitektur deep learning di mana dua jaringan saraf, yang disebut generator dan diskriminator, bersaing satu sama lain. Generator mencoba untuk menghasilkan data yang menyerupai data asli, sedangkan diskriminator mencoba untuk membedakan antara data asli dan data yang dihasilkan.

## Framework dan Libraries untuk AI

1. Python:

- TensorFlow: Framework yang dikembangkan oleh Google untuk deep learning.
- PyTorch: Framework deep learning yang dikembangkan oleh Facebook.
  scikit-learn: Library machine learning yang menyediakan algoritma-algoritma machine learning yang efisien dan mudah digunakan.
- NLTK (Natural Language Toolkit): Library untuk pemrosesan bahasa alami.
- OpenCV: Library untuk pengolahan gambar dan video.

2. Java:

- Deeplearning4j: Framework deep learning untuk Java dan Scala yang didesain untuk kecepatan dan skalabilitas.
- Weka: Library untuk pemrosesan data dan analisis prediktif.
- DL4J (DeepLearning4J): Framework deep learning untuk Java yang dapat digunakan untuk berbagai tugas machine learning.

3. Go:

- Gorgonia: Library deep learning yang mirip dengan TensorFlow, tetapi untuk Go.
- Golearn: Library machine learning yang menyediakan berbagai algoritma machine learning untuk Go.
- Gonum: Library numerik untuk Go, yang dapat digunakan untuk komputasi ilmiah dan analisis data.

# Thank You :star2:
