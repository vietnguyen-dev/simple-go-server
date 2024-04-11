const button = document.getElementById("message");
const message = document.getElementById("message-text");

document.addEventListener("click", async () => {
	const data = await fetch('/messages')
	const res = await data.json()
	message.innerHTML = res[res.length - 1].message
})
