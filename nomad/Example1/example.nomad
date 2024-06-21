job "example" {
  datacenters = ["dc1"]

  group "example" {
    task "server" {
      driver = "docker"

      config {
        image = "nginx"
        ports = ["http"]
      }

      resources {
        network {
          port "http" {
            static = 8080
          }
        }
      }
    }
  }
}
