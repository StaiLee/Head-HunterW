
export default function hideLoader(time) {
    const loader = document.getElementById('loader');
    const body = document.getElementById('body-container')

    document.body.classList.add("stop-scrolling")

    setTimeout(() => {
        try {
            // @ts-ignore
            loader.style.opacity = 0;
            // @ts-ignore
            loader.style.height = 0;
            // @ts-ignore
            loader.style.width = 0;
            // @ts-ignore
            body.style.opacity = 1;
        } catch {
            console.log("Petite erreur")
        }
        document.body.classList.remove("stop-scrolling")
    }, time)
}
