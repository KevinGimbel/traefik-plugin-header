# traefik-plugin-header

`traefik-plugin-header` is a [Traefik 2](https://doc.traefik.io/traefik/) plugin for assigning headers. This is very much redundant and useless and only intended for testing out how to write Traefik plugins.

## Usage

`traefik.yaml`
```yaml
pilot:
    token: "xxxx"
experimental:
    plugins:
        traefik-plugin-header:
            moduleName: "github.com/kevingimbel/traefik-plugin-header"
            version: "v0.1.2"

middlewares:
    my-traefik-plugin-header:
        plugin:
            traefik-plugin-header:
                Headers:
                    - key: de.kevingimbel/version
                      value: "1.0"
```

I am not sure if the config above actually works and I also don't know how to test a plugin before publishing it. 😬

I think it should be usable with Docker label like so

```yaml
your-container: #
  image: your-docker-image

  labels:
    # Attach kg_header-default@file middleware (declared in file)
    - "traefik.http.routers.my-container.middlewares=my-traefik-plugin-header@file"
```

## Acknowledgment

The base structure of this plugin was directly copied from the "blockpath" plugin by Traefik which can be found at [https://github.com/traefik/plugin-blockpath](https://github.com/traefik/plugin-blockpath)

## License

MIT, see LICENSE file