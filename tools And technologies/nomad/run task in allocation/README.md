# Problem Statement

Currently with stdin and stdout i am able to do the required task. but as i change it to custom one other then standard ones, then i am getting error.

The Default command must be /bash/bin, so that we can start terminal process and thus issue a command.

there after we have to send data to stdout and it will read it, process it and return a response. This is how it is going to be. Just like how our terminal is stdin and stdout. 

main hurdle is how we keep buffer common between them so that no anamolies happens.