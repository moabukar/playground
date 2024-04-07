# Setup

```sh


echo -n 'your-github-pat-here' | base64
Replace your-github-pat-here with your actual PAT in the secrets.yml. The output will be your token encoded in Base64.

k apply -f secret.yml
k apply -f deploy.yml
```
