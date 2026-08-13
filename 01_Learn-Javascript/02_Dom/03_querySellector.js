let input = document.getElementById("input")
let tombol = document.getElementById("tombol")
let rmv = document.getElementById("rmv")
let daftar = document.querySelector("#buah");


let buah = [
    "Mangga",
    "Apel",
    "Jeruk",
];


tombol.addEventListener("click",
function(){
    buah.push(input.value)
    let item = document.createElement("li");
    item.textContent = input.value;
    daftar.appendChild(item);
    console.log(buah)}
)

rmv.addEventListener("click",
function(){
    buah=[]
    daftar.innerHTML = "";
    console.log(buah)
}
)


for (let i = 0; i < buah.length; i++) {
    let item = document.createElement("li");
    item.textContent = buah[i];
    daftar.appendChild(item);}


// jadi. tapi klo daftarnya yang di remove, ntar gabisa buat list lagi