const source = new EventSource("http://localhost:3000/sse")
source.onmessage = (event) => {
    console.log("OnMessage called:")
    console.log(event)
    console.log(JSON.parse(event.data))
}