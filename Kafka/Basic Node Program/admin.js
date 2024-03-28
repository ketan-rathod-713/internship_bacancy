const kafka = require("./client").kafka
// there can be multiple brokers

async function init(){
    console.log("start")
    const admin = kafka.admin();
    console.log("admin connecting..")
    admin.connect()
    console.log("admin connected successfully")


    console.log("creating topic [rider-updates]")
    admin.createTopics({
        topics: [{
            topic: "rider-updates",
            numPartitions: 2, // North india and south india
        }]
    })

    console.log("topic created success..")

    console.log("disconnecting admin..")
    await admin.disconnect();
}

init();