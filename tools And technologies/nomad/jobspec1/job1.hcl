job "http-echo" {
    datacenters = ["dc1"]
    group "echo" {
        count = 2

        network {
            # mbits = 10 # Bandwidth allocation for the task
            port "http" {
              # static = 8080 # Now our port has become dynamic assignment. static means port assigned to this application, This nomad application. Not internall or anything.
              #  to = 8080 # what do mean by it.
              # Now our port is dynamic so what port we will give to our docker container or whatever application. so for that nomad gives a run time variables to deal with it. NOMAD_PORT_<label>, where <label> is the name we gave to our port, which in our case is simply http.
            }
        }

        task "server" {
            driver = "docker"

            config {
                image = "hashicorp/http-echo:latest"
                args = [
                    # as now we are using dynamic port so we will change it from 8080 to the dynamic port name that we have defined.
                    "-listen", ":${NOMAD_PORT_http}",
                    "-text", "Hello and welcome to localhost${NOMAD_PORT_http}, I am being served by this machine which is a client of Nomad, on port ${NOMAD_PORT_http}"
                ]
                ports = ["http"]
            }

            resources {
                cpu = 500 # MHZ
                memory = 128 # MB
                # No network configuration here
            }
        }
    }
}

# Each of the 5 allocations will be using a different dynamic port. While it is possible to query Nomad to obtain these ports, it can become somewhat laborious, so let’s look at an easier way to access our application.

# we need a way to run consul and need to register http-echo application in catalog.
# we can solve both this problem using nomad.
# schedule consul in nomad to run natively, rather then using container in docker engine, raw_exec will run a task in the same OS as nomad is runnning  in.
