let button = document.getElementById("butt")
/**
 * @type {HTMLInputElement}
 */
let inp = document.getElementById("inp")

button.addEventListener('click', async () => {
    let short_url = await fetch("/short_url", {
        method: 'POST',
        body: inp.value
    })

    let lnk = await short_url.text()
    console.log(lnk)

    let link = document.createElement('p')
    link.innerText = lnk

    document.body.append(link)
})
