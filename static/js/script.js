let currentIndex = 2; // Mulai dari logo Pemasaran (indeks 2)
const jurusanImages = document.querySelectorAll('.jurusan-wrapper img');

function updateSlides() {
    // Reset semua gambar ke default
    jurusanImages.forEach((img, index) => {
        img.classList.remove('active');
        img.style.transform = 'scale(1.5)'; // Reset skala untuk gambar yang tidak aktif
        img.style.opacity = '0.2'; // Reset opasitas untuk gambar yang tidak aktif
        img.style.zIndex = '1'; 
    });

    // Tambahkan kelas aktif ke gambar saat ini
    jurusanImages[currentIndex].classList.add('active');
    jurusanImages[currentIndex].style.transform = 'scale(2.4)'; // Perbesar gambar aktif
    jurusanImages[currentIndex].style.opacity = '1'; // Set opasitas penuh untuk gambar aktif
    jurusanImages[currentIndex].style.zIndex = '10'; 

    // Pusatkan gambar aktif
    const offset = (currentIndex - 2) * 100; // Sesuaikan offset dengan gambar aktif di tengah
    document.querySelector('.jurusan-wrapper').style.transform = `translateX(${-offset}px)`;
}

// Tombol untuk navigasi ke kanan
document.querySelector('.next').addEventListener('click', () => {
    if (currentIndex === jurusanImages.length - 1) {
        currentIndex = 0; // Kembali ke logo pertama jika sudah di akhir
    } else {
        currentIndex++;
    }
    updateSlides();
});

// Tombol untuk navigasi ke kiri
document.querySelector('.prev').addEventListener('click', () => {
    if (currentIndex === 0) {
        currentIndex = jurusanImages.length - 0; // Kembali ke logo terakhir jika sudah di awal
    } else {
        currentIndex-1;
    }
    updateSlides();
});

// Panggil awal untuk mengatur slide pertama
updateSlides();



// Tombol untuk navigasi ke kiri
document.querySelector('.prev').addEventListener('click', () => {
    if (currentIndex === 0) {
        currentIndex = 1; // Kembali ke logo Pemasaran jika sudah di awal
    } else {
        currentIndex--;
        jurusanImages[currentIndex].style.zIndex = '10';
    }
    updateSlides();
});

// Panggil awal untuk mengatur slide pertama
updateSlides();



let currentSlideIndex = 0;

function moveSlide(n) {
  const slides = document.querySelectorAll('.slide');
  const totalSlides = slides.length;  // Akan menjadi 4 sekarang

  currentSlideIndex += n;

  if (currentSlideIndex >= totalSlides) {
    currentSlideIndex = 0;
  }

  if (currentSlideIndex < 0) {
    currentSlideIndex = totalSlides - 1;
  }

  updateSlide();
}

function currentSlide(n) {
  currentSlideIndex = n;
  updateSlide();
}

function updateSlide() {
  const sliderWrapper = document.querySelector('.slider-wrapper');
  sliderWrapper.style.transform = `translateX(${-currentSlideIndex * 100}%)`;

  const dots = document.querySelectorAll('.dot');
  dots.forEach(dot => dot.classList.remove('active'));
  dots[currentSlideIndex].classList.add('active');
}

// Ganti Gambar PPDB
let ppdbIndex = 0;

// Daftar gambar untuk PPDB
const ppdbImages1 = ['../static/assets/img/ppdb1.jpg', '../static/assets/img/ppdb3.jpg']; // Gambar pertama ke gambar ketiga
const ppdbImages2 = ['../static/assets/img/ppdb2.jpeg', '../static/assets/img/ppdb4.jpg']; // Gambar kedua ke gambar keempat

// Ambil elemen gambar
const imageElement1 = document.getElementById('ppdb-1');
const imageElement2 = document.getElementById('ppdb-2');

// Fungsi untuk mengganti gambar
function changePPDBImages() {
    // Update gambar pertama dan kedua sesuai index
    imageElement1.src = ppdbImages1[ppdbIndex];
    imageElement2.src = ppdbImages2[ppdbIndex];

    // Update index untuk gambar berikutnya
    ppdbIndex = (ppdbIndex + 1) % ppdbImages1.length; // Loop kembali ke awal setelah dua gambar
}

// Ganti gambar setiap 5 detik (5000 milidetik)
setInterval(changePPDBImages, 3000);

// Panggil fungsi secara langsung untuk menampilkan gambar pertama
changePPDBImages();




