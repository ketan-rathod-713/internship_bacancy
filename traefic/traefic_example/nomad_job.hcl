job "nomadContainers" {
    group "nomadContainersGroup" {
        count = 1
        task "nomadTraefic" {
            driver = "docker"

            config {
                image = "nginx"
            }
        }


        service {
            name = "nomadService"
            tags = [
                "traefik.http.routers.my-router.rule=Host(`example.com`)"
            ]
        }
    }
}