fetch("/api/prestasi/showAll")
  .then((response) => response.json())
  .then((data) => {
    const maxItemsPerPage = 6; // Maksimal 6 item per halaman
    const galleryContainer = document.getElementById("galleryContainer"); // Container untuk semua halaman
    galleryContainer.innerHTML = ""; // Kosongkan isi sebelumnya

    // Hitung jumlah halaman
    const totalPages = Math.ceil(data.length / maxItemsPerPage);

    for (let page = 0; page < totalPages; page++) {
      const pageDiv = document.createElement("div");
      pageDiv.classList.add("gallery");
      if (page === 0) pageDiv.classList.add("active"); // Set halaman pertama sebagai aktif
      pageDiv.id = `page${page + 1}`;

      // Ambil 6 item per halaman
      const start = page * maxItemsPerPage;
      const end = start + maxItemsPerPage;
      const pageItems = data.slice(start, end);

      // Tambahkan item ke halaman
      pageItems.forEach((prestasi, index) => {
        const galleryItem = document.createElement("div");
        galleryItem.classList.add("gallery-item");

        const img = document.createElement("img");
        img.src = `${prestasi.gambar}`;
        img.alt = `Prestasi ${start + index + 1}`;
        galleryItem.appendChild(img);

        pageDiv.appendChild(galleryItem);
      });

      // Masukkan halaman ke dalam container
      galleryContainer.appendChild(pageDiv);
    }

    // Buat navigasi dots
    const dotsContainer = document.querySelector(".dots");
    dotsContainer.innerHTML = ""; // Kosongkan dots sebelumnya

    for (let i = 0; i < totalPages; i++) {
      const dot = document.createElement("span");
      dot.classList.add("dot");
      if (i === 0) dot.classList.add("active"); // Set dot pertama sebagai aktif
      dot.onclick = () => showPage(i);
      dotsContainer.appendChild(dot);
    }
  })
  .catch((error) => {
    console.error("Error fetching data:", error);
  });

// Fungsi untuk menampilkan halaman berdasarkan indeks
function showPage(pageIndex) {
  const allPages = document.querySelectorAll(".gallery");
  const allDots = document.querySelectorAll(".dot");

  allPages.forEach((page, index) => {
    if (index === pageIndex) {
      page.classList.add("active");
    } else {
      page.classList.remove("active");
    }
  });

  allDots.forEach((dot, index) => {
    if (index === pageIndex) {
      dot.classList.add("active");
    } else {
      dot.classList.remove("active");
    }
  });
}

  document.addEventListener("DOMContentLoaded", function() {
    fetch('/api/berita/showAll')
      .then(response => response.json())
      .then(berita => {
        const prestasiContainer = document.querySelector(".prestasi-container");
        prestasiContainer.innerHTML = '';
  
        berita.forEach(berita => {
          // Pastikan path gambar benar
          const imgPath = berita.gambar.startsWith("uploads/")
            ? berita.gambar
            : `${berita.gambar}`;
  
          const prestasiCard = document.createElement("div");
          prestasiCard.classList.add("prestasi-card");
  
          prestasiCard.innerHTML = `
            <img src="${imgPath}" alt="Prestasi ${berita.ID}">
            <h5 style="color: black;">${berita.judul}</h5>
            <p>${berita.isi_berita}</p>
          `;
  
          prestasiContainer.appendChild(prestasiCard);
        });
      })
      .catch(error => console.error("Error fetching data:", error));
  });
  