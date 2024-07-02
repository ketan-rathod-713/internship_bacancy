job "mysql-server" {
  datacenters = ["dc1"]
  
  group "mysql-server" {
    count = 1
    
    task "mysql-server" {
      driver = "docker"
      
      config {
        image = "mysql/mysql-server:8.0"
        
        port_map {
          db = 3306
        }
        
        volumes = [
          "docker-entrypoint-initdb.d/:/docker-entrypoint-initdb.d/",
        ]
      }
      
      template {
        data = <<EOH
CREATE DATABASE dbwebappdb;
CREATE USER 'dbwebapp'@'%' IDENTIFIED BY 'dbwebapp';
GRANT ALL PRIVILEGES ON dbwebappdb.* TO 'dbwebapp'@'%';
GRANT ALL PRIVILEGES ON . TO 'dbwebapp'@'localhost' IDENTIFIED BY 'dbwebapp';
EOH
        
        destination = "db.sql"
      }
      
      resources {
        cpu = 500
        memory = 1024
        
        network {
          mbits = 10
          
          port "db" {
            static = 3306
          }
        }
      }
    }
  }
}