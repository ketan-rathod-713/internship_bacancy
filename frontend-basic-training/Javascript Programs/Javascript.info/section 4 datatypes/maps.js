const mp = new Map()

mp.set("key", "value")
mp.set("key2", "value2")
mp.set("key3", "value3")

console.log(mp)

for(let i of mp.keys()){
    console.log(i)
}

for(let i of mp.values()){
    console.log(i)
}

for(let i of mp.entries()){
    console.log(i)
}