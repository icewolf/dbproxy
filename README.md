# dbproxy

## Usage

To run, use a similar command to the old one, but with `cloudflared dbconnect [args] replaced with dbproxy.exe start [args]`

A new `--port` parameter is required. This is the port that the db connect API will be served from, so it should be a port that is available on the server. `--hostname` is no longer required.


## How it works

- `dbproxy` and `cloudflared` run side-by-side as services.
- `dbproxy` is setup as a service using something like `nssm`
  - Sample command: `dbproxy.exe start --port DB-PROXY-PORT --url mssql://user:pass@domain:port --application-aud YOUR-ACCESS-POLICY-ID --auth-domain yourdomain.cloudflareaccess.com`
- Latest version of cloudflared is setup as a service with the UI configuration, but now pointing to `https://localhost:DB-PROXY-PORT`
- Once correctly setup, you should be able to make a GET request to `https://example.com/ping`, with headers `cf-access-client-id` and `cf-access-client-secret` set to the values of a service token.
