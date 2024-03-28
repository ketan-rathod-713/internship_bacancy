const {kafka} = require("./client")

async function init(){
    const producer = kafka.producer();

    console.log("connecting")

    await producer.connect();

    console.log("connected")

    await producer.send({
        topic: "rider-updates",
        messages: [
            {
                partition: 0,
                key: "location-update", value: JSON.stringify({name: "tonny stark", loc: "SOUTH"})
            }
        ]
    })


    console.log("done")

    await producer.disconnect()

}

init()