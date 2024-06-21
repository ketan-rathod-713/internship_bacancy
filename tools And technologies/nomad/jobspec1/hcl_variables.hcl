variable "http_port" {
    type = string
    default = 8080
    description = "port for listening to port on given docker image"
}

job "example_nginx" {
    group "example" {
        task "nginx" {
            driver = "docker"

            config {
                image = "nginx:latest"
                port_map {
                    http = 80 # Wow port map need to see it TODO:
                }
            }

            resources {
                cpu = 500
                memory = 256
                network {
                    mbits = 10
                    port "http" {
                        static = 8080
                    }
                }
            }
        }
    }
}

# Notes
# config for docker important. It specifies internal port to bind with.
# see any documentation related to configs of driver in task section