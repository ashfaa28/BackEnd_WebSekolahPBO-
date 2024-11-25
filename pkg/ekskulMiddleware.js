document.addEventListener("DOMContentLoaded", function () {
    showAllEkskul();
});

// Fungsi untuk menampilkan semua ekskul
function showAllEkskul() {
    fetch('/api/ekskul/showAll')
        .then(response => response.json())
        .then(ekskuls => {
            const tableBody = document.getElementById('ekskulTableBody');
            tableBody.innerHTML = '';
            ekskuls.forEach(ekskul => {
                const tr = document.createElement('tr');

                // Tambahkan gambar dan nama ekskul
                const imgPath = ekskul.gambar.startsWith("/") ? ekskul.gambar : `/${ekskul.gambar}`;
                const imgCell = document.createElement('td');
                imgCell.innerHTML = `
                <div class="d-flex px-2 py-1">
                    <div><img src="${imgPath}" class="avatar avatar-sm me-3 border-radius-lg" width="150" height="150" alt="ekskul"></div>
                    <div class="d-flex flex-column justify-content-center">
                        <h6 class="mb-0 text-sm text-black">${ekskul.nama_ekskul}</h6>
                    </div>
                </div>`;
                tr.appendChild(imgCell);

                // Tambahkan tombol Edit dan Hapus
                const actionCell = document.createElement('td');
                actionCell.className = "align-middle text-center";
                actionCell.innerHTML = `
                <a href="javascript:void(0);" class="text-secondary font-weight-bold text-xs me-3" onclick="showEditModal(${ekskul.ID})">Edit</a>
                <a href="javascript:void(0);" class="text-secondary font-weight-bold text-xs" onclick="showDeleteModal(${ekskul.ID})">Hapus</a>`;
                tr.appendChild(actionCell);

                tableBody.appendChild(tr);
            });
        })
        .catch(error => console.error('Error fetching data:', error));
}

// Fungsi untuk menambah ekskul
function submitAddForm() {
    const formData = new FormData(document.getElementById("addEkskulForm"));

    fetch('/api/ekskul/create', {
        method: 'POST',
        body: formData
    })
        .then(response => {
            if (response.ok) {
                window.location.reload();
            } else {
                alert("Gagal menambahkan ekskul");
            }
        })
        .catch(error => console.error('Error:', error));
}

// Fungsi untuk menampilkan modal edit
function showEditModal(id) {
    fetch(`/api/ekskul/${id}`)
        .then(response => response.json())
        .then(data => {
            const idInput = document.querySelector("#editEkskulForm input[name='id']");
            const namaEkskulInput = document.querySelector("#editEkskulForm input[name='nama_ekskul']");

            idInput.value = id; // Set the ID directly from parameter
            namaEkskulInput.value = data.nama_ekskul;

            new bootstrap.Modal(document.getElementById('editModal')).show();
        })
        .catch(error => console.error('Error:', error));
}

// Fungsi untuk mengedit ekskul
function submitEditForm() {
    const form = document.getElementById("editEkskulForm");
    const formData = new FormData(form);
    const id = document.querySelector("#editEkskulForm input[name='id']").value;

    fetch(`/api/ekskul/update/${id}`, {
        method: 'PUT',
        body: formData
    })
        .then(response => response.json())
        .then(data => {
            if (data.error === false) {
                window.location.reload();
            } else {
                alert("Gagal mengedit ekskul");
            }
        })
        .catch(error => console.error('Error updating data:', error));
}

// Fungsi untuk menampilkan modal hapus
function showDeleteModal(id) {
    document.querySelector("#deleteModal .btn-danger").setAttribute("onclick", `confirmDelete(${id})`);
    new bootstrap.Modal(document.getElementById('deleteModal')).show();
}

// Fungsi untuk menghapus ekskul
function confirmDelete(id) {
    fetch(`/api/ekskul/delete/${id}`, { method: 'DELETE' })
        .then(response => {
            if (response.ok) {
                window.location.reload();
            } else {
                alert("Gagal menghapus ekskul");
            }
        })
        .catch(error => console.error('Error deleting data:', error));
}
