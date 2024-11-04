document.getElementById('registrasi-form').addEventListener('submit', function(e) {
    e.preventDefault();
    
    const formData = new FormData(this);
    const data = {};
    formData.forEach((value, key) => data[key] = value);

    fetch('/api/siswa/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
        const successMessage = document.getElementById('success-message');
        const errorMessage = document.getElementById('error-message');
        const daftarButton = document.getElementById('daftar-button');

        if (data.error) {
            // Tampilkan pesan error dari server di elemen #error-message
            errorMessage.style.display = 'block';
            errorMessage.textContent = data.message;  // Pesan dari server
            successMessage.style.display = 'none';
        } else {
            // Jika tidak ada error, tampilkan pesan sukses dan sembunyikan tombol daftar
            successMessage.style.display = 'block';
            errorMessage.style.display = 'none';
            daftarButton.remove();
            daftarButton.disabled = true; // Nonaktifkan tombol

            // Bersihkan pesan error sebelumnya
            errorMessage.textContent = "";
        }
    })
    .catch(err => {
        // Jika terjadi masalah jaringan, tampilkan pesan error umum
        const errorMessage = document.getElementById('error-message');
        errorMessage.style.display = 'block';
        errorMessage.textContent = 'Terjadi kesalahan pada server. Mohon coba lagi.';
    });
});