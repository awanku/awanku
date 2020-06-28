job "awanku-stack-core-api" {
    datacenters = ["dc1"]
    group "core-api" {
        task "core-api" {
            driver = "docker"
            config {
                image = "docker.awanku.id/awanku/core-api:latest"
                auth {
                    username = "awanku"
                    password = "rahasia"
                }
                force_pull = true
                port_map {
                    http = 3000
                }
            }
            service {
                name = "awanku-core-api"
                port = "http"
                check {
                    type = "http"
                    path = "/status"
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
                    "traefik.http.routers.awanku-stack-core-api-http.rule=Host(`api.awanku.id`)",
                    "traefik.http.routers.awanku-stack-core-api-http.entrypoints=http",
                    "traefik.http.routers.awanku-stack-core-api-http.middlewares=httpToHttps@consul",
                    "traefik.http.routers.awanku-stack-core-api-https.rule=Host(`api.awanku.id`)",
                    "traefik.http.routers.awanku-stack-core-api-https.entrypoints=https",
                    "traefik.http.routers.awanku-stack-core-api-https.tls=true",
                    "traefik.http.routers.awanku-stack-core-api-https.tls.options=default",
                ]
            }
            env {
                DB_URL = "postgres://awanku:rahasia@${NOMAD_IP_postgresql_pg}:${NOMAD_PORT_postgresql_pg}/awanku?sslmode=disable"
            }
            resources {
                network {
                    port "http" {}
                }
            }
        }
        task "postgresql" {
            driver = "docker"
            config {
                image = "postgres:12"
                port_map {
                    pg = 5432
                }
                volumes = [
                    "/awanku/maindb/pgdata:/var/lib/postgresql/data"
                ]
            }
            service {
                name = "awanku-db-main"
                port = "pg"
                check {
                    type     = "tcp"
                    port     = "pg"
                    interval = "10s"
                    timeout  = "1s"
                    check_restart {
                        limit = 3
                        grace = "30s"
                    }
                }
            }
            env {
                POSTGRES_USER = "awanku"
                POSTGRES_PASSWORD = "rahasia"
                POSTGRES_DB = "awanku"
            }
            resources {
                network {
                    port "pg" {}
                }
            }
        }
    }
}
