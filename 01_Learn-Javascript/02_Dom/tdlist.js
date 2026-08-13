tambah = document.getElementById("tambah")
input = document.getElementById("input")
let list = document.querySelector("#list")

let daftar = ["halo"]

tambah.addEventListener("click", 
    function(){
        daftar.push(input.value)
        console.log(daftar)
        let item = document.createElement("li")
        item.textContent = (input.value)
        list.appendChild(item)

        // hapus
        let hapus = document.createElement("button")
        hapus.textContent="hapus";
        list.appendChild(hapus)
        hapus.addEventListener("click", 
            function(){
                item.remove()
                hapus.remove()
                let index = daftar.indexOf(item.textContent);
                daftar.splice(index, 1);
                console.log(daftar)
            }
        )
    }
)
