const {kafka} = require("./client")

async function init(){
    const consumer = kafka.consumer({groupId: "user-1"});

    console.log("connecting")

    await consumer.connect();

    console.log("connected")

    consumer.subscribe({topics: ["rider-updates"], fromBeginning: true})

    await consumer.run({
        eachMessage: async ({topic, partition, message}) => {
            console.log("TOPIC",topic, "PARTITION",partition)
            console.log(message.value.toString())
        }
    })

    // await consumer.disconnect()
}

init()