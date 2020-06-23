job "awanku-stack-landing" {
    datacenters = ["dc1"]
    reschedule {
        delay          = "30s"
        delay_function = "exponential"
        max_delay      = "1h"
        unlimited      = true
    }
    group "landing" {
        task "landing-webui" {
            driver = "docker"
            config {
                image = "docker.awanku.id/awanku/landing-webui:latest"
                auth {
                    username = "awanku"
                    password = "rahasia"
                }
                port_map {
                    http = 3000
                }
            }
            service {
                name = "awanku-landing-webui"
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
                    "traefik.http.routers.awanku-stack-landing-webui-http.rule=Host(`awanku.id`)",
                    "traefik.http.routers.awanku-stack-landing-webui-http.entrypoints=http",
                    "traefik.http.routers.awanku-stack-landing-webui-http.middlewares=httpToHttps@consul",
                    "traefik.http.routers.awanku-stack-landing-webui-https.rule=Host(`awanku.id`)",
                    "traefik.http.routers.awanku-stack-landing-webui-https.entrypoints=https",
                    "traefik.http.routers.awanku-stack-landing-webui-https.tls=true",
                    "traefik.http.routers.awanku-stack-landing-webui-https.tls.options=default",
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
