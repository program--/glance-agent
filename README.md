# Simple Glance Agent

This is a temporary [glance](https://github.com/glanceapp/glance) agent service written in Go for the [server-stats widget](https://github.com/glanceapp/glance/blob/main/docs/configuration.md#server-stats). See https://github.com/glanceapp/glance/issues/360 for more details.

## Usage

The following environment variables are available:

- `GLANCE_AGENT_SECRET` (required): secret token to use. Must match what is set in glance configuration.
- `GLANCE_AGENT_HOST` (optional): host interface address to bind service to. Defaults to an empty string.
- `GLANCE_AGENT_PORT` (optional): port to bind service to. Defaults to `8080`.

## References
- [Server stats widget internals](https://github.com/glanceapp/glance/blob/c88fd526e55117445c7f4440c83b661faa402047/internal/glance/widget-server-stats.go)
- [glance sysinfo package](https://github.com/glanceapp/glance/blob/c88fd526e55117445c7f4440c83b661faa402047/pkg/sysinfo/sysinfo.go)
