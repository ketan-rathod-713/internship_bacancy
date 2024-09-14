module example.com/hello

go 1.18

// TODO: accessing local modules
replace example.com/greetings/greetings => ../greetings/

require example.com/greetings/greetings v0.0.0-00010101000000-000000000000
