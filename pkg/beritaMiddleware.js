document.addEventListener("DOMContentLoaded", function() {
    showAllBerita();
  });
  
  // Fungsi untuk menampilkan semua berita
  function showAllBerita() {
    fetch('/api/berita/showAll')
      .then(response => response.json())
      .then(beritas => {
        const tableBody = document.getElementById('beritaTableBody');
        tableBody.innerHTML = '';
        beritas.forEach(berita => {
          const tr = document.createElement('tr');
  
          // Tambahkan gambar dan isi berita
          const imgPath = berita.gambar.startsWith("uploads/") ? berita.gambar : `uploads/berita/${berita.gambar}`;
          const imgCell = document.createElement('td');
          imgCell.innerHTML = `
            <div class="d-flex px-2 py-1">
              <div><img src="${imgPath}" class="avatar avatar-sm me-3 border-radius-lg" width="150" height="150" alt="berita"></div>
              <div class="d-flex flex-column justify-content-center">
                <h6 class="mb-0 text-sm">${berita.isi_berita}</h6>
              </div>
            </div>`;
          tr.appendChild(imgCell);
  
          // Tombol Edit dan Hapus
          const actionCell = document.createElement('td');
          actionCell.className = "align-middle text-center";
          actionCell.innerHTML = `<a href="javascript:void(0);" class="text-secondary font-weight-bold text-xs" onclick="showDeleteModal(${berita.ID})">Hapus</a>`;
          tr.appendChild(actionCell);
  
          tableBody.appendChild(tr);
        });
      })
      .catch(error => console.error('Error fetching data:', error));
  }
  
  // Fungsi untuk menambah berita
  function submitAddForm() {
    const formData = new FormData(document.getElementById("addBeritaForm"));
  
    fetch('/api/berita/create', {
      method: 'POST',
      body: formData
    })
    .then(response => {
      if (response.ok) {
        window.location.reload();
      } else {
        alert("Gagal menambahkan berita");
      }
    })
    .catch(error => console.error('Error:', error));
  }
  
  // Fungsi untuk menampilkan modal edit
  function showEditModal(id) {
    fetch(`/api/berita/${id}`)
        .then(response => response.json())
        .then(data => {
            const idInput = document.querySelector("#editBeritaForm input[name='id']");
            const isiBeritaInput = document.querySelector("#editBeritaForm textarea[name='isi_berita']");
            
            idInput.value = id;  // Set the ID directly from parameter
            isiBeritaInput.value = data.isi_berita;
            
            new bootstrap.Modal(document.getElementById('editModal')).show();
        })
        .catch(error => console.error('Error:', error));
}


  
  // Fungsi untuk mengedit berita
function submitEditForm() {
    const form = document.getElementById("editBeritaForm");
    const formData = new FormData(form);
    const id = document.querySelector("#editBeritaForm input[name='id']").value;

    fetch(`/api/berita/update/${id}`, {
        method: 'PUT',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.error === false) {
            window.location.reload();
        } else {
            alert("Gagal mengedit berita");
        }
    })
    .catch(error => console.error('Error updating data:', error));
}



  
  // Fungsi untuk menampilkan modal hapus
  function showDeleteModal(id) {
    document.querySelector("#deleteModal .btn-danger").setAttribute("onclick", `confirmDelete(${id})`);
    new bootstrap.Modal(document.getElementById('deleteModal')).show();
  }
  
  // Fungsi untuk menghapus berita
  function confirmDelete(id) {
    fetch(`/api/berita/delete/${id}`, { method: 'DELETE' })
      .then(response => {
        if (response.ok) {
          window.location.reload();
        } else {
          alert("Gagal menghapus berita");
        }
      })
      .catch(error => console.error('Error deleting data:', error));
  }
  
  async function fetchBerita(id) {
    try {
        const response = await fetch(`/api/berita/show/${id}`);
        if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);
        const berita = await response.json();
        console.log(berita);
        // Populate edit modal with berita data
    } catch (error) {
        console.error("Error fetching data:", error);
    }
}
