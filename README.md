# traefik-plugin-header

`traefik-plugin-header` is a [Traefik 2](https://doc.traefik.io/traefik/) plugin for assigning headers. This is very much redundant and useless and only intended for testing out how to write Traefik plugins.

## Usage

`traefik.toml`
```toml
# Define plugin
[experimental.plugins]
  [experimental.plugins.kg_header]
    modulename = "github.com/kevingimbel/traefik-plugin-header"
    version = "v0.1.2"

# Configure a reusable middleware called "kg_header-default"
[http.middlewares]
  [http.middlewares.kg_header-default.plugin.kg_header]
    # Rewrites all "foo" occurences by "bar"
    [[http.middlewares.kg_header-default.plugin.kg_header.headers]]
      key = "kevingimbel.de/version"
      value = "0.1.2"
    [[http.middlewares.kg_header-default.plugin.kg_header.headers]]
      key = "kevingimbel.de/works"
      value = "true"
```

I am not sure if the config above actually works and I also don't know how to test a plugin before publishing it. 😬

I think it should be usable with Docker label like so

```yaml
your-container: #
  image: your-docker-image

  labels:
    # Attach kg_header-default@file middleware (declared in file)
    - "traefik.http.routers.my-container.middlewares=kg_header-default@file"
```

## Acknowledgment

The base structure of this plugin was directly copied from the "blockpath" plugin by Traefik which can be found at [https://github.com/traefik/plugin-blockpath](https://github.com/traefik/plugin-blockpath)

## License

MIT, see LICENSE file