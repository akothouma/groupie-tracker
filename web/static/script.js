const accordion = document.querySelectorAll(".accordion")

accordion.forEach((item) => {
    item.addEventListener("click", () => {

        item.classList.toggle("open_desc")

        let desc = item.querySelector(".desc")
        let plus = item.querySelector(".plus")
        let minus = item.querySelector(".minus")

        if (item.classList.contains("open_desc")){
            desc.style.height = desc.scrollHeight + "px"
            desc.style.margin = "20px 0 0"
            plus.style.display = "none"
            minus.style.display = "block"
        }else{
            desc.style.height = "0"
            desc.style.margin = "0"
            plus.style.display = "block"
            minus.style.display = "none"
        }
    })
})