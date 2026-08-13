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
// allert habis, allerynya gw pake google buat nulis. sebenernya ga penting ko

let daftar = [
    {
        tugas: "bangun",
        selesai: false
    },
    {
        tugas: "mandi",
        selesai: false
    },
];

function lclstorage(){
    localStorage.setItem(
    "daftar",
    JSON.stringify(daftar)
);
}

function consoledaftar(){
    console.log("data daftar:")
    for(let i = 0; i<daftar.length; i++){
    console.log(i+1 +"."+ daftar[i].tugas +"    "+daftar[i].selesai)
}  console.log("selesai")
}

function tampilkanitem(task) {
    let item = document.createElement("li")
    item.textContent = task.tugas

    if(task.selesai){
        item.classList.add("selesai")}
    else{
            item.classList.remove("selesai")
        }

    list.appendChild(item)
    let checkbox = document.createElement("input")
    checkbox.type="checkbox"
    checkbox.checked = task.selesai
    
    list.appendChild(checkbox)

    let hapus = document.createElement("button")
    hapus.textContent = "hapus"
    list.appendChild(hapus)




    hapus.addEventListener("click", function() {
        hapusfunc(task, item, hapus, checkbox)
    })
    checkbox.addEventListener("change", function(){
        task.selesai = checkbox.checked
        if(task.selesai){
        item.classList.add("selesai")}
        else{
            item.classList.remove("selesai")
        }
        lclstorage()
        consoledaftar()
        
    })
}

function hapusfunc(task, item, hapus, checkbox) {
    item.remove()
    hapus.remove()
    checkbox.remove()
    let index = daftar.indexOf(task)
    daftar.splice(index, 1)
    lclstorage()
    consoledaftar()
    AlertHapus()
}

let data = localStorage.getItem("daftar")

if(data != null){
    let dataArray = JSON.parse(data)
    daftar = dataArray
};

tambah.addEventListener("click",
    function(){
        daftar.push({
            tugas: input.value,
            selesai: false,
        })
        lclstorage()
        tampilkanitem(daftar[daftar.length -1])
        console.log("Item ["+daftar[daftar.length-1].tugas+"] Berhasil Ditambahkan")
        consoledaftar()
        AlertTambah();
    })

for (let i = 0; i < daftar.length; i++) {
    tampilkanitem(daftar[i])
}

