fetch("/api/prestasi/showAll")
  .then((response) => response.json())
  .then((data) => {
    const sliderWrapper = document.getElementById("sliderWrapper");
    sliderWrapper.innerHTML = ""; // Kosongkan slider sebelum menambahkan gambar baru
    data.forEach((prestasi, index) => {
      const slide = document.createElement("div");
      slide.classList.add("slide");
      const img = document.createElement("img");
      const imgPath = `${prestasi.gambar}`;
      console.log(`Mengakses gambar: ${imgPath}`); // Log untuk debugging
      img.src = imgPath;
      img.alt = `Penghargaan ${index + 1}`;
      slide.appendChild(img);
      sliderWrapper.appendChild(slide);
    });

    // Update dots sesuai jumlah gambar
    const dotsContainer = document.querySelector(".dots");
    dotsContainer.innerHTML = ""; // Kosongkan dots sebelum menambah yang baru
    data.forEach((_, index) => {
      const dot = document.createElement("span");
      dot.classList.add("dot");
      dot.onclick = () => currentSlide(index);
      dotsContainer.appendChild(dot);
    });

    showSlide(0); // Menampilkan slide pertama setelah gambar dimuat
  })
  .catch((error) => {
    console.error("Error fetching data:", error);
  });


  document.addEventListener("DOMContentLoaded", function() {
    // Ambil data prestasi dari API
    fetch('/api/berita/showAll')
      .then(response => response.json())
      .then(berita => {
        const prestasiContainer = document.querySelector(".prestasi-container");
        prestasiContainer.innerHTML = ''; // Bersihkan konten sebelumnya
  
        berita.forEach(berita => {
          // Pastikan path gambar benar
          const imgPath = berita.gambar.startsWith("uploads/")
            ? berita.gambar
            : `uploads/berita/${berita.gambar}`;
  
          // Buat elemen prestasi-card baru
          const prestasiCard = document.createElement("div");
          prestasiCard.classList.add("prestasi-card");
  
          prestasiCard.innerHTML = `
            <img src="${imgPath}" alt="Prestasi ${berita.ID}">
            <h3>${berita.judul}</h3>
            <p>${berita.isi_berita}</p>
          `;
  
          // Tambahkan card ke container
          prestasiContainer.appendChild(prestasiCard);
        });
      })
      .catch(error => console.error("Error fetching data:", error));
  });
  