let tombolSimpan = document.getElementById("simpan")
let tombolLihat = document.getElementById("lihat")
let nama = document.getElementById("nama")

tombolSimpan.addEventListener("click", function(){

    let data = nama.value

    localStorage.setItem("nama", data)

    console.log("Data tersimpan")

})

tombolLihat.addEventListener("click", function(){

    let data = localStorage.getItem("nama")

    console.log(data)

})