let nama    =   document.getElementById("nama")
let pwd     =   document.getElementById("pwd")
let button  =   document.getElementById("button")

button.addEventListener("click", 
    function(){
        if      (nama.value == ""){
            console.log("Nama Harus Di isi")
        }
        
        else if (pwd.value == ""){
            console.log("Password Harus Di isi")
        }
        else if (nama.value.length < 8    ||    pwd.value.length < 8){
            console.log("Minimal 8 Karakter")
        }
        else{
            console.log("Success")
        }
    })