function loadScript(){
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve("Script is Loaded")
        }, 2000);
    })
}

loadScript().then((result, error) => {
    if(error){
        console.log("erorr ",error)
        // it is not runnig ha ha
    } else {
        console.log("result ",result)
    }
}).catch(err => {
    console.log("error ", err)
})