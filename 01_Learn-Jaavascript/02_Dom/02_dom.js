let judul = document.getElementById("judul");
let datanama = document.getElementById("datanama");
let dataumur = document.getElementById("dataumur");
let datajurusan = document.getElementById("datajurusan");
let tombol = document.getElementById("tombol");
let i = 1;


let inpt = document.getElementById("inpt")
let umr = document.getElementById("umur")
let jrsn = document.getElementById("jurusan")
let dataset = inpt + umr + jrsn;


let daftar = document.getElementById("daftar");
let klik = document.getElementById("klik");
let daftarNama = []
let hapush = []


let remove = document.getElementById("remove");



tombol.addEventListener("click", 
    function(){
    if (inpt.value === ""||umr.value===""||jrsn.value===""){
        alert("data harus diisi!")
        datanama.textContent = "Data Harus Di Isi Lengkap!";
        dataumur.textContent = "";
        datajurusan.textContent = "";
    }
    else if (inpt.value.length < 8){
        datanama.textContent = "Nama Minimal 8 karakter";
        dataumur.textContent = "";
        datajurusan.textContent = "";
    }
    else if (umr.value > 60){
        datanama.textContent = "Umur maksimal 60 tahun";
        dataumur.textContent = "";
        datajurusan.textContent = "";
    }

        // berhasil

    else if(umr.value < 18){
    datanama.textContent = "Nama: "+inpt.value;
    dataumur.textContent = "Status: Belum Dewasa";
    datajurusan.textContent = "Jurusan: "+jrsn.value;

    daftarNama.push(inpt.value)
    let item = document.createElement("li")
    item.textContent = inpt.value
    daftar.appendChild(item)

    // hapus
    hapush.push (inpt.value)
    let hapus = document.createElement("button")
    hapus.textContent = "Hapus"
    daftar.appendChild(hapus)
    hapus.addEventListener("click", function(){
        item.remove()
        hapus.remove()
    })
    }
    
    else{
    datanama.textContent = "Nama: "+inpt.value;
    dataumur.textContent = "Status: Dewasa";
    datajurusan.textContent = "Jurusan: "+jrsn.value;

    daftarNama.push(inpt.value)
    let item = document.createElement("li")
    item.textContent = inpt.value
    daftar.appendChild(item)

    hapush.push (inpt.value)
    let hapus = document.createElement("button")
    hapus.textContent = "Hapus"
    daftar.appendChild(hapus)
    hapus.addEventListener("click", function(){
        item.remove()
        hapus.remove()
    })
    }}
);

remove.addEventListener("click", function(){
    
    daftar.innerHTML=""
})


