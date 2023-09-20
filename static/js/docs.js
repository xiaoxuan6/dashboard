load()

async function load() {
    let target = localStorage.getItem("docs")
    const item = await (
        await fetch(`/apis/docs${target}`)
    ).json()

    if (item && item.status === 200) {
        console.log(item)
    }
}
