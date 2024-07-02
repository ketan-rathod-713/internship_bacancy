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
      port "http" {
        static = 3306
      }
    }

    task "mysql" {
      driver = "docker"

      config {
        image = "mysql:latest"
        ports = ["http"]
        image_pull_timeout="10m"
      }

      env {
        MYSQL_ROOT_PASSWORD = admin
        MYSQL_DATABASE = bacancy
        MYSQL_USER = bacancy
        MYSQL_PASSWORD = admin
      }

      resources {
        cpu    = 500 # in MHz
        memory = 512 # in MB
      }
      
      volume_mount {
        volume      = "database-data"
        destination = "/docker-entrypoint-initdb.d" #<-- in the container
        read_only   = false
	  }

      volume_mount {
        volume      = "mysql-config"
        destination = "/etc/mysql/conf.d" #<-- in the container
        read_only   = false
	  }

    }
    
    volume "database-data" {
        type      = "host"
        read_only = false
        source    = "database-data"
    }

    volume "mysql-config" {
        type      = "host"
        read_only = false
        source    = "mysql-config"
    }

  }
}
