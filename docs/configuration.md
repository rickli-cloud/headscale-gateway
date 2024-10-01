# Configuration

| Environment variable               | default                                  | Description               |
| ---------------------------------- | ---------------------------------------- | ------------------------- |
| `HSGW_OIDC_ISSUER`                 |                                          | **Required**              |
| `HSGW_OIDC_CLIENT_ID`              |                                          | **Required**              |
| `HSGW_OIDC_CLIENT_SECRET`          |                                          |                           |
| `HSGW_LISTEN_ADDR`                 | 0.0.0.0:8000                             | Server listen address     |
| `HSGW_HEADSCALE_SOCKET`            | unix:///var/run/headscale/headscale.sock | The headscale unix socket |
| `HSGW_ACCESS_CONTROL_ALLOW_ORIGIN` |                                          | CORS sources              |
