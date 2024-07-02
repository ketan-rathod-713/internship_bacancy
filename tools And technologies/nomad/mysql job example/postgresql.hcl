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
    }

    task "setup-db" {
      driver = "exec"

      lifecycle {
        hook = "poststart"
        sidecar = false
      }

      config {
        command = "/bin/sh"
        args = ["-c", "until psql -h localhost -U postgres -c '\\q'; do echo 'Waiting for postgres to be ready...'; sleep 2; done; psql -h localhost -U postgres -c 'CREATE DATABASE my_new_db;'"]
      }

      env {
        PGPASSWORD = "admin"
      }
    }
  }
}
