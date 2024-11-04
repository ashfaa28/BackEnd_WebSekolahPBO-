fetch("/api/prestasi/showAll")
  .then((response) => response.json())
  .then((data) => {
    const sliderWrapper = document.getElementById("sliderWrapper");
    sliderWrapper.innerHTML = "";
    data.forEach((prestasi, index) => {
      const slide = document.createElement("div");
      slide.classList.add("slide");
      const img = document.createElement("img");
      const imgPath = `${prestasi.gambar}`;
      img.src = imgPath;
      img.alt = `Penghargaan ${index + 1}`;
      slide.appendChild(img);
      sliderWrapper.appendChild(slide);
    });

    const dotsContainer = document.querySelector(".dots");
    dotsContainer.innerHTML = "";
    data.forEach((_, index) => {
      const dot = document.createElement("span");
      dot.classList.add("dot");
      dot.onclick = () => currentSlide(index);
      dotsContainer.appendChild(dot);
    });

    showSlide(0);
  })
  .catch((error) => {
    console.error("Error fetching data:", error);
  });


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
  