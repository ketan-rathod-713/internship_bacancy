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
        image_pull_timeout="10m"
      }

      env {
        POSTGRES_DB = "bacancy"
        POSTGRES_USER = "postgres"
        POSTGRES_PASSWORD = "admin"
        PGPASSWORD = "admin"
      }

      resources {
        cpu    = 500 # in MHz
        memory = 512 # in MB
      }
    }

  }
}
