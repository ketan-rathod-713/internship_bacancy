job "postgresql" {
  datacenters = ["dc1"]
  type = "service"

  group "postgresql" {
    count = 1

    restart {
      attempts = 10
      interval = "5m"
      delay    = "25s"
      mode     = "delay"
    }

    network {
      port "http" {
        static = 5432
      }
    }

    task "postgresql" {
      driver = "docker"

      config {
        image = "postgres:latest"
        ports = ["http"]
      }

      env {
        POSTGRES_DB = "bacancy"
        POSTGRES_USER = "postgres"
        POSTGRES_PASSWORD = "admin"
      }

      resources {
        cpu    = 500 # in MHz
        memory = 512 # in MB
      }

      volume_mount {
        volume      = "postgres-init"
        destination = "/docker-entrypoint-initdb.d"
        propagation_mode = "private"
      }
    }

    volume "postgres-init" {
      type      = "host"
      source    = "postgres-init"
      read_only = false
    }
  }
}