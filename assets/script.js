let currentIndex = 2; // Mulai dari logo Pemasaran (indeks 2)
const jurusanImages = document.querySelectorAll('.jurusan-wrapper img');

function updateSlides() {
    // Reset semua gambar ke default
    jurusanImages.forEach((img) => {
        img.classList.remove('active');
        img.style.transform = 'scale(1)'; // Reset skala untuk gambar yang tidak aktif
        img.style.opacity = '0.5'; // Reset opasitas untuk gambar yang tidak aktif
    });

    // Tambahkan kelas aktif ke gambar saat ini
    jurusanImages[currentIndex].classList.add('active');
    jurusanImages[currentIndex].style.transform = 'scale(1.3)'; // Perbesar gambar aktif
    jurusanImages[currentIndex].style.opacity = '1'; // Set opasitas penuh untuk gambar aktif

    // Pusatkan gambar aktif
    const offset = (currentIndex - 2) * 170; // Sesuaikan offset dengan gambar aktif di tengah
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
        currentIndex--;
    }
    updateSlides();
});

// Panggil awal untuk mengatur slide pertama
updateSlides();





// Tombol untuk navigasi ke kiri
document.querySelector('.prev').addEventListener('click', () => {
    if (currentIndex === 0) {
        currentIndex = 2; // Kembali ke logo Pemasaran jika sudah di awal
    } else {
        currentIndex--;
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
