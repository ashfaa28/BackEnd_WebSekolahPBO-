
function openAddModal() {
  document.getElementById('modalTitle').innerText = "Add New Student";
  document.getElementById('studentForm').reset();
  document.getElementById('studentModal').style.display = 'block';
}

function openEditModal(id) {
  document.getElementById('modalTitle').innerText = "Edit Student";
  fetch(`/db/admin/siswa/${id}`)
    .then(response => response.json())
    .then(data => {
      document.getElementById('studentId').value = data.ID;
      document.getElementById('nama').value = data.Nama;
      document.getElementById('studentModal').style.display = 'block';
    });
}

function saveStudent() {
  const id = document.getElementById('studentId').value;
  const nama = document.getElementById('nama').value;
  const url = id ? `/db/admin/siswa/edit/${id}` : '/db/admin/siswa/create';
  const method = id ? 'PUT' : 'POST';

  fetch(url, {
    method: method,
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ nama: nama })
  })
  .then(response => response.json())
  .then(data => {
    alert(data.message);
    window.location.reload();
  });
}

function deleteStudent(id) {
  if (confirm("Are you sure you want to delete this student?")) {
    fetch(`/db/admin/siswa/delete/${id}`, { method: 'DELETE' })
      .then(response => response.json())
      .then(data => {
        alert(data.message);
        window.location.reload();
      });
  }
}

function closeModal() {
  document.getElementById('studentModal').style.display = 'none';
}
