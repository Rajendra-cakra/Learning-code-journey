let input = document.getElementById("input")
let submit = document.getElementById("submit")
let tampil = document.getElementById("tampil")
let hapus = document.getElementById("hapus")

let data = []
let datatotal = data.length

function sapa(){
for(let i = 0; i < data.length; i++){
    console.log(data[i])
    
}}

submit.addEventListener("click",
    function(){
        data.push((data.length+1)+". "+input.value)
        console.log(data)
    }
)

tampil.addEventListener("click",
    function sapa(){
    for(let i = 0; i < data.length; i++){
        console.log(data[i])
        
    }}
)

hapus.addEventListener("click",
    function(){
        for(let a = 0; a < 100; a++){
        let index = data.indexOf(data)
        data.splice(index, 1)}
        sapa()
    }
)


