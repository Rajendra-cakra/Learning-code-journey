let nama = document.getElementById("nama")
let umur = document.getElementById("umur")

let button = document.getElementById("button")

let datanama = document.getElementById("datanama")
let dataumur = document.getElementById("dataumur")

let daftar = document.getElementById("daftar")
let alldata = []







button.addEventListener("click", 
    function(){
        if(umur.value ===""|| nama.value ==="" ){
            datanama.textContent = "data harus di isi"
            dataumur.textContent = ""
        }
        else if(nama.value.length < 8  ){
            datanama.textContent = "Nama Minimal 8 karakter"
            dataumur.textContent = ""
        }

        // berhasil

        else if(umur.value < 18  ){
            datanama.textContent = "Nama: " +nama.value
            dataumur.textContent = "Status: Belum dewasa"

            let dataobject = {
                nama: nama.value,
                umur: umur.value
            }

            alldata.push(dataobject)
            let dafnama = document.createElement("li")
            dafnama.textContent = nama.value
            daftar.appendChild(dafnama)
            console.log(alldata)

            let hapus = document.createElement("button")
            hapus.textContent = "Hapus"
            daftar.appendChild(hapus)
            hapus.addEventListener("click", 
                function(){
                    // remove aray
                    let index = alldata.indexOf(dataobject)
                    alldata.splice(index, 1)

                    // remove dom
                    dafnama.remove()
                    hapus.remove()
                    console.log(alldata)
                }
            )
        }

        else{
            datanama.textContent = "Nama: " +nama.value
            dataumur.textContent = "Status: Sudah Dewasa"

            let dataobject = {
                nama: nama.value,
                umur: umur.value
            }

            alldata.push(dataobject)
            let dafnama = document.createElement("li")
            dafnama.textContent = nama.value
            daftar.appendChild(dafnama)
            console.log(alldata)

            let hapus = document.createElement("button")
            hapus.textContent = "Hapus"
            daftar.appendChild(hapus)
            hapus.addEventListener("click", 
                function(){
                    // remove aray
                    let index = alldata.indexOf(dataobject)
                    alldata.splice(index, 1)

                    // remove dom
                    dafnama.remove()
                    hapus.remove()
                    console.log(alldata)

                    

                }
            )
        }

})
