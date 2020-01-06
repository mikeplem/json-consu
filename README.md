# Learning Golang, JSON, Consul

I am thinking about changing how data is stored/retrieved in dashboard-admin and dashboard-client applications. My thinking is that JSON would allow me more flexibiity with how data gets passed around and that I could add more fields the JSON object to expand the capabilities of the client.

I am using this code to help me understand how they all work together.

## Consul Configuration

<http://github.com/hashicorp/consul/api>

I am starting Consul in the following way for my testing.

```shell
./consul agent -dev
```
