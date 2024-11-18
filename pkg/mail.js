    document.getElementById("contactForm").addEventListener("submit", function (e) {
        e.preventDefault();

        // Ambil data form
        const formData = new FormData(this);

        // Kirim data menggunakan fetch
        fetch("/Sendmail", {
            method: "POST",
            body: formData,
        })
            .then(response => response.json())
            .then(data => {
                if (data.success) {
                    // Tampilkan SweetAlert jika berhasil
                    Swal.fire({
                        icon: "success",
                        title: "Success",
                        text: data.message,
                        confirmButtonText: "OK",
                    }).then(() => {
                        // Arahkan ke halaman "/"
                        window.location.href = "/";
                    });
                } else {
                    // Tampilkan SweetAlert jika gagal
                    Swal.fire({
                        icon: "error",
                        title: "Error",
                        text: data.message,
                    });
                }
            })
            .catch(error => {
                console.error("Error:", error);
                Swal.fire({
                    icon: "error",
                    title: "Error",
                    text: "Something went wrong, please try again later.",
                });
            });
    });