document.addEventListener("DOMContentLoaded", () => {
    const detailButtons = document.querySelectorAll(".detail-btn");
    const popups = document.querySelectorAll(".popup-container");

    detailButtons.forEach((button) => {
        button.addEventListener("click", () => {
            const targetId = button.getAttribute("data-target");
            const popup = document.getElementById(targetId);
            if (popup) {
                popup.style.display = "flex";
                document.body.classList.add("popup-active");
            }
        });
    });

    popups.forEach((popup) => {
        popup.addEventListener("click", (e) => {
            if (e.target === popup) {
                popup.style.display = "none";
                document.body.classList.remove("popup-active");
            }
        });
    });
});


document.addEventListener("DOMContentLoaded", function () {
    const nums = document.querySelectorAll('.num');

    nums.forEach((num) => {
        let startVal = 0;
        let endVal = parseInt(num.getAttribute('data-val'));
        let duration = 2000; // Durasi animasi dalam ms
        let increment = endVal / (duration / 50); // Menentukan kecepatan kenaikan angka per interval
        let interval = setInterval(() => {
            startVal += increment;
            num.textContent = Math.floor(startVal);
            
            if (startVal >= endVal) {
                num.textContent = endVal;
                clearInterval(interval);
            }
        }, 50);
    });
    
    // Menambahkan class 'show' untuk memulai animasi ketika halaman dimuat
    setTimeout(() => {
        nums.forEach((num) => {
            num.classList.add('show');
        });
    }, 8); // Delay 0.5 detik sebelum animasi dimulai
});




// Fungsi untuk menampilkan pop-up
function openPopup() {
    document.getElementById("pplg-detail").style.display = "flex";
}

// Fungsi untuk menutup pop-up
function closePopup() {
    document.getElementById("pplg-detail").style.display = "none";
}

// Menutup pop-up saat halaman dimuat
window.onload = function () {
    closePopup();
};

// Fungsi untuk membuka pop-up
function openPopup(id) {
    document.getElementById(id).style.display = "flex";
}

// Fungsi untuk menutup pop-up
function closePopup() {
    const popups = document.querySelectorAll(".popup-container");
    popups.forEach((popup) => (popup.style.display = "none"));
}


let currentIndex = 0;  // Indeks untuk slider saat ini
const items = document.querySelectorAll('.team-member');  // Ambil semua elemen tim
const sliderContainer = document.querySelector('.slider-container');
const totalItems = items.length;  // Total jumlah item tim
let maxSlides = 5;  // Batasan untuk slide ke kanan, default 5

// Fungsi untuk mendeteksi apakah perangkat dalam mode mobile
function isMobile() {
  return window.innerWidth < 768; // Misalnya, 768px dianggap mobile
}

// Fungsi untuk memperbarui batas slide berdasarkan mode perangkat
function updateMaxSlides() {
  if (isMobile()) {
    maxSlides = 6;  // Jika mobile, max slides 6
  } else {
    maxSlides = 3;  // Jika desktop, max slides 5
  }
  // Setelah memperbarui batas, sesuaikan posisi slider sesuai mode
  updateSliderPosition();
}

// Fungsi untuk memperbarui posisi slider
function updateSliderPosition() {
  const slideWidth = items[0].offsetWidth + 20; // Lebar item + margin (sesuaikan dengan lebar item)
  sliderContainer.style.transform = `translateX(-${currentIndex * slideWidth}px)`; // Geser slider sesuai index
}

// Fungsi untuk pergi ke slide berikutnya
function goToNextSlide() {
  // Jika sudah mencapai batas slide, reset ke awal
  if (currentIndex < totalItems - 1) {
    currentIndex++;
  } else {
    currentIndex = 0; // Kembali ke awal jika sudah di slide terakhir
  }

  if (currentIndex >= maxSlides) {
    currentIndex = 0; // Reset ke awal jika melebihi batas maxSlides
  }

  updateSliderPosition();
}

// Fungsi untuk pergi ke slide sebelumnya
function goToPrevSlide() {
  if (currentIndex > 0) {
    currentIndex--;
  } else {
    currentIndex = totalItems - 1; // Kembali ke item terakhir jika di awal
  }
  updateSliderPosition();
}

// Event listeners untuk tombol next dan prev
document.querySelector('.next').addEventListener('click', goToNextSlide);
document.querySelector('.prev').addEventListener('click', goToPrevSlide);

// Memperbarui maxSlides saat halaman dimuat atau ukuran jendela berubah
window.addEventListener('resize', updateMaxSlides);
updateMaxSlides(); // Menyesuaikan maxSlides saat halaman pertama kali dimuat


// Menyimpan referensi gambar yang sedang diperbesar
let currentExpanded = null;

function expandEkstrakulikuler(card, description) {
  // Reset semua gambar ke ukuran kecil dan hapus class 'expanded'
  const allCards = document.querySelectorAll('.ekstrakulikuler-card');
  allCards.forEach((c) => {
    c.style.width = '100px';
    c.style.height = '300px';
    c.classList.remove('expanded');
  });

  // Jika gambar yang diklik berbeda dari gambar yang sedang diperbesar
  if (currentExpanded !== card) {
    // Perbesar gambar yang diklik
    card.style.width = '300px';
    card.style.height = '300px';
    card.classList.add('expanded');
    currentExpanded = card;

    // Tampilkan deskripsi
    const descriptionElement = document.getElementById('ekstrakulikuler-description');
    descriptionElement.innerText = description;
    descriptionElement.style.opacity = '1'; // Pastikan deskripsi terlihat
  } else {
    // Kembali ke ukuran semula
    currentExpanded = null;

    // Kosongkan deskripsi
    const descriptionElement = document.getElementById('ekstrakulikuler-description');
    descriptionElement.innerText = '';
    descriptionElement.style.opacity = '0'; // Sembunyikan deskripsi
  }
}

let currentPage = 1;
const totalPages = 3;

function showPage(pageNumber) {
    const pages = document.querySelectorAll('.gallery');
    pages.forEach((page, index) => {
        page.classList.toggle('active', index + 1 === pageNumber);
    });
    document.querySelector('.page-number').textContent = pageNumber;
}

function nextPage() {
    currentPage = (currentPage % totalPages) + 1;
    showPage(currentPage);
}

document.addEventListener('DOMContentLoaded', () => {
    showPage(currentPage);
    setInterval(nextPage, 4000);
});