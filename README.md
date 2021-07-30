# felipegfalcao/gocache-prometheus-exporter

> A GoCache API Exporter for Prometheus.

---

## Table of Contents

1. [Introduction](#introduction)
2. [Usage](#usage)
3. [Contributing](#contributing)
4. [License](#license)

### Introduction

The _gocache_exporter_ is a simple server that removes the GoCache API to obtain domain access statistics and exports them via HTTP for consumption by Prometheus.


_First I would like to share with everyone that this was my first code using GO. I believe this will 
definitely not be the best code you will see. But it was the only alternative to 
capture metrics from the GoCache platform. I decided to publish this repository not only because I lack knowledge in the language 
but to be able to use the strength of a community to improve this code. I count on everyone's help._

### Usage

The first step is to grab an API key from the [GoCache site]. Metrics are captured based on 
this doc. [GoCache Api].

**Important:** Be advised to set a high scrape interval (e.g. 5min). Each scrape
performs a direct API call and to frequent requests can lead to the
_deauthorization_ of your API key!

**Note:** Since _gocache_ isn't a very handy word, the metric namespace is
`gocache`.

#### Installation

TODO: The easiest way to run the _gocache_ is by grabbing the latest binary from
the [release page][release]. 

##### Building from source

```bash
git clone https://github.com/felipegfalcao/gocache-prometheus-exporter
cd gocache-prometheus-exporter
go build .
```

#### Using the application

```bash
./gocache-prometheus-exporter [flags]
```

#### Example ARGS "REQUIRED":

```bash
./gocache-prometheus-exporter -domain <DOMAIN> -token <token_GoCache>
```

At the moment we can pass another _-host_ argument that allows filtering by subdomain. 
The value must be the complete subdomain (ex: www.yourdomain.com)

```bash
./gocache-prometheus-exporter -domain <DOMAIN> -token <TOKEN_GoCache> -host <host>
```


####**We are currently with a fixed value in the range of capturing latest value**

### Contributing

Feel free to submit PRs or to fill Issues. Every kind of help is appreciated.

### License

Distributed under Apache License (`Apache License, Version 2.0`).

See [LICENSE](LICENSE) for more information.

<!-- Links -->

[GoCache site]: https://painel.gocache.com.br/home.php#account-info
[gocache api]: https://docs.gocache.com.br/api/#api-Analytics-GetAnalytics