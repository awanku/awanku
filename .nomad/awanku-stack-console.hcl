job "awanku-stack-console" {
    datacenters = ["dc1"]
    reschedule {
        delay          = "30s"
        delay_function = "exponential"
        max_delay      = "1h"
        unlimited      = true
    }
    group "console" {
        task "console-webui" {
            driver = "docker"
            config {
                image = "docker.awanku.id/awanku/console-webui:latest"
                auth {
                    username = "awanku"
                    password = "rahasia"
                }
                port_map {
                    http = 3000
                }
            }
            service {
                name = "awanku-console-webui"
                port = "http"
                check {
                    type = "http"
                    path = "/"
                    port = "http"
                    interval = "10s"
                    timeout = "1s"
                    check_restart {
                        limit = 3
                        grace = "60s"
                    }
                }
                tags = [
                    "traefik.enable=true",
                    "traefik.http.routers.awanku-console-webui-http.rule=Host(`console.awanku.id`)",
                    "traefik.http.routers.awanku-console-webui-http.entrypoints=http",
                    "traefik.http.routers.awanku-console-webui-http.middlewares=httpToHttps@consul",
                    "traefik.http.routers.awanku-console-webui-https.rule=Host(`console.awanku.id`)",
                    "traefik.http.routers.awanku-console-webui-https.entrypoints=https",
                    "traefik.http.routers.awanku-console-webui-https.tls=true",
                    "traefik.http.routers.awanku-console-webui-https.tls.certresolver=gratisan",
                    "traefik.http.routers.awanku-console-webui-https.tls.options=default",
                ]
            }
            env {
                AWANKU_CORE_API_URL = ""
            }
            resources {
                network {
                    port "http" {}
                }
            }
        }
    }
}
