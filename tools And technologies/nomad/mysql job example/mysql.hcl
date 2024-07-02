job "mysql" {
  datacenters = ["dc1"]
  type = "service"

  group "mysql" {
    count = 1

    restart {
      attempts = 10
      interval = "5m"
      delay    = "25s"
      mode     = "delay"
    }

    network {
      port "mysql" {
        static = 3306
      }
    }

    task "mysql" {
      driver = "docker"

      config {
        image = "mysql:latest"
        ports = ["mysql"]
        args = ["-c", " mysql -P 3306 -u root"]
      }

      env {
        MYSQL_DATABASE = "bacancy"
        MYSQL_USER = "mysql"
        MYSQL_PASSWORD = "admin"
        MYSQL_ROOT_PASSWORD = "admin"
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
        args = ["-c", "until mysqladmin ping -hlocalhost -uroot -padmin --silent; do echo 'Waiting for MySQL to be ready...'; sleep 2; done; mysql -hlocalhost -uroot -padmin -e 'CREATE DATABASE my_new_db;'"]
      }
    }
  }
}
