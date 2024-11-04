document.addEventListener("DOMContentLoaded", function() {
  fetch('/api/prestasi/showAll') 
    .then(response => response.json())
    .then(prestasis => {
      const tableBody = document.getElementById('prestasiTableBody');
      prestasis.forEach(prestasi => {
        console.log(`Gambar: ${prestasi.gambar}`);
        const tr = document.createElement('tr');

        const imgPath = prestasi.gambar.startsWith("/uploads/")
          ? prestasi.gambar
          : `/${prestasi.gambar}`;

        const imgCell = document.createElement('td');
        imgCell.innerHTML = `
          <div class="d-flex px-2 py-1">
            <div>
              <img src="${imgPath}" class="avatar avatar-sm me-3 border-radius-lg" width="150" height="150" alt="prestasi">
            </div>
            <div class="d-flex flex-column justify-content-center">
              <h6 class="mb-0 text-sm">Sertifikat</h6>
            </div>
          </div>`;
        tr.appendChild(imgCell);

        const dateCell = document.createElement('td');
        const date = new Date(prestasi.CreatedAt);
        dateCell.className = "align-middle text-center";
        dateCell.innerHTML = `<span class="text-secondary text-xs font-weight-bold">${date.toLocaleDateString()}</span>`;
        tr.appendChild(dateCell);

        const actionCell = document.createElement('td');
        actionCell.className = "align-middle text-center";
        actionCell.innerHTML = `<a href="#" class="text-secondary font-weight-bold text-xs" onclick="deletePrestasi(${prestasi.ID})">Hapus</a>`;
        tr.appendChild(actionCell);

        tableBody.appendChild(tr);
      });
    })
    .catch(error => console.error('Error fetching data:', error));
});


document.getElementById('prestasiForm').addEventListener('submit', async (e) => {
  e.preventDefault();
  
  const formData = new FormData(e.target);
  
  try {
    const response = await fetch('/api/prestasi/create', {
      method: 'POST',
      body: formData
    });
    
    if (response.ok) {
      window.location.reload();
    } else {
      alert('Gagal mengunggah prestasi');
    }
  } catch (error) {
    console.error('Error:', error);
    alert('Terjadi kesalahan');
  }
});

document.querySelector('#addModal .btn-primary').addEventListener('click', () => {
  document.getElementById('prestasiForm').dispatchEvent(new Event('submit'));
});

function formatDate(dateString) {
  const date = new Date(dateString);
  return date.toLocaleDateString("id-ID");
}

function deletePrestasi(id) {
  fetch(`/api/prestasi/delete/${id}`, { method: 'DELETE' })
    .then(response => {
      if (response.ok) {
        alert("Data prestasi berhasil dihapus");
        location.reload(); 
      } else {
        alert("Gagal menghapus data prestasi");
      }
    })
    .catch(error => console.error('Error deleting data:', error));
}




