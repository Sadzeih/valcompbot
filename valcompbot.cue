package main

import (
        "dagger.io/dagger"
        "universe.dagger.io/docker"
)

// Example usage in a plan
dagger.#Plan & {
        client: {
                filesystem: "./": read: contents: dagger.#FS
                env: {
                        TAG: string
                        GITHUB_ACTOR: string
                        GITHUB_TOKEN: dagger.#Secret
                }
        }

        actions: {
                build: docker.#Dockerfile & {
                        source: client.filesystem."./".read.contents
                        dockerfile: path: "./Dockerfile"
                }

                push:  docker.#Push & {
                        image: build.output
                        dest:  "ghcr.io/sadzeih/valcompbot:\(client.env.TAG)"

                        auth: {
                                username: client.env.GITHUB_ACTOR
                                secret: client.env.GITHUB_TOKEN
                        }
                }
        }
}
