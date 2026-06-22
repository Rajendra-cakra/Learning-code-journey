let input = document.getElementById("input")
let tambah = document.getElementById("tambah")
let list = document.querySelector("#list")

// allert
function AlertTambah() {
  const alertBox = document.getElementById("customAlert");
  
  // 1. Memunculkan alert
  alertBox.style.display = "block";
  alertBox.textContent = input.value + " Berhasil Ditambahkan"
  
  // 2. menyembunyikan alert. 2000 = 2detik
  setTimeout(() => {
    alertBox.style.display = "none";
  }, 2000);
}

function AlertHapus() {
  const alertBox = document.getElementById("customAlert");
  
  // 1. Memunculkan alert
  alertBox.style.display = "block";
  alertBox.textContent = "Data Berhasil Dihapus"
  
  // 2. menyembunyikan alert. 2000 = 2detik
  setTimeout(() => {
    alertBox.style.display = "none";
  }, 2000);
}
// allert habis


let daftar = [
    "Bangun",
    "Mandi",
];

tambah.addEventListener("click",
    function(){
        daftar.push(input.value)
        let item = document.createElement("li")
        item.textContent = input.value;
        list.appendChild(item)
        console.log("Item ["+input.value+"] Berhasil Ditambahkan")
        AlertTambah();

        // hapus
        let hapus = document.createElement("button")
        hapus.textContent = "hapus";
        list.appendChild(hapus);

        hapus.addEventListener("click",
            function(){
            item.remove()
            hapus.remove()
            console.log("Sebelum:", daftar);
            let index = daftar.indexOf(item.textContent);
            console.log("Index:", index);
            daftar.splice(index, 1);
            console.log("Sesudah:", daftar);
            AlertHapus();
            })
    })


for (let i = 0; i < daftar.length; i++) {
    let item = document.createElement("li");
    item.textContent = daftar[i];
    list.appendChild(item);

    let hapus = document.createElement("button")
    hapus.textContent = "hapus";
    list.appendChild(hapus);

    hapus.addEventListener("click",
            function(){
            item.remove()
            hapus.remove()
            console.log("Sebelum:", daftar);
            let index = daftar.indexOf(item.textContent);
            console.log("Index:", index);
            daftar.splice(index, 1);
            console.log("Sesudah:", daftar);
            AlertHapus();
            })
}
