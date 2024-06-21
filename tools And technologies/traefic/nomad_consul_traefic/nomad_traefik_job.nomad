job "traefik" {
  datacenters = ["dc1"]

  group "traefik" {

    # group level network blog
    network {
          port "http" {
            static = 80
          }
        }

    task "traefik" {
      driver = "docker"

      config {
        image = "traefik:v3.0"
        ports = ["http"]

        volumes = [
          "local/traefik.toml:/etc/traefik/traefik.toml"
        ]

        command = [
          "--configFile=/etc/traefik/traefik.toml"
        ]
      }

      service {
        name = "traefik"
        port = "http"
      }
    }
  }
}
