data_dir = "/opt/nomad/data"

client {
  enabled = true
  servers = ["127.0.0.1:4647"]

  host_volume "database-data" {
    path = "/home/bacancy/Desktop/Bacancy/internship_bacancy/tools And technologies/nomad/mysql job example/postgres-init"
    read_only = false
    }

  host_volume "mysql-config" {
    path = "/home/bacancy/Desktop/Bacancy/internship_bacancy/tools And technologies/nomad/mysql job example/mysql-config"
    read_only = false
    }
}

ports {
  http = 5656  # Change to any available port
}

plugin "docker" {
    volumes {
      enabled      = true
      selinuxlabel = "z"
    }
}

