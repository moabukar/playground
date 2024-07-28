# Tailscale K8s operator

```sh

helm repo add tailscale https://pkgs.tailscale.com/helmcharts
helm repo update

helm upgrade \
  --install \
  tailscale-operator \
  tailscale/tailscale-operator \
  --namespace=tailscale \
  --create-namespace \
  --set-string oauth.clientId="kDHHapFv1A11CNTRL" \
  --set-string oauth.clientSecret="tskey-auth-kDHHapFv1A11CNTRL-G5R7DoP5HFGPmtCd9VLzFGH7WBfn5uYxi" \
  --wait


```