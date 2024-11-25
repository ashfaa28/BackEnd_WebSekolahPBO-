document.getElementById("siswaForm").addEventListener("submit", function (event) {
    event.preventDefault(); 
  
    const formData = new FormData(this);
    const formObject = {};
    formData.forEach((value, key) => {
      formObject[key] = value;
    });
  
    // Kirim data ke server
    fetch("/api/siswa/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formObject),
    })
      .then(response => response.json())
      .then(data => {
        if (data.error) {
          if (data.message.includes("NISN sudah terdaftar")) {
            Swal.fire({
              icon: 'error',
              title: 'Yah Gagal karena',
              text: 'NISN sudah terdaftar.',
              confirmButtonText: 'OK',
            });
          } else {
            Swal.fire({
              icon: 'error',
              title: 'Gagal',
              text: data.message,
              confirmButtonText: 'OK',
            });
          }
        } else {
          Swal.fire({
            icon: 'success',
            title: 'Berhasil',
            text: 'Selamat anda sudah mendaftar tunggu info selanjutnya ya!!!!.',
            confirmButtonText: 'OK',
          });
        }
      })
      .catch(error => {
        console.error("Error:", error);
        Swal.fire({
          icon: 'error',
          title: 'Terjadi Kesalahan',
          text: 'Terjadi kesalahan saat mengirim data.',
          confirmButtonText: 'OK',
        });
      });
  });
  