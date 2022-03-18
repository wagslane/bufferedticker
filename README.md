# bufferedticker

A ticker that keeps a buffer of ticks. Particularly useful for rate-limiting when the rolling window of time on which the rate limit is calculated is larger than the time between requests.

[![](https://godoc.org/github.com/wagslane/bufferedticker?status.svg)](https://godoc.org/github.com/wagslane/bufferedticker)![Deploy](https://github.com/wagslane/bufferedticker/workflows/Tests/badge.svg)

## Motivation

Imagine you are able to make 100 requests to an API each hour and create a normal ticker to limit how often you make requests.

```go
ticker := time.NewTicker(time.Hour / 100)
for _ := range ticker.C {
    go makeRequest()
}
```

In this scenario, you'll make exactly 100 requests per hour. However, what if you are making requests on behalf of clients? The requests don't come in to your system in a uniform manner, but you still can't make more than 100 per hour to the API you depend on. In this scenario, if you use a normal ticker and all the requests come in the last half of the hour, you'll only make 50 requests for the hour! That wastes half your budget!

The buffered ticker is a very thin wrapper around a normal ticker that lets you "save" the ticks for a rainy day. In the scenario described above, you could do the following:

```go
ticker := bufferedticker.NewTicker(time.Hour / 100, 100)
for _ := range ticker.C {
    go makeRequest()
}
```

This ensures you never make more than 100 requests in a single hour, but also that you can spend your entire hourly budget, even if it's just in the last minutes of the hour.

## ⚙️ Installation

Inside a Go module:

```bash
go get github.com/wagslane/bufferedticker
```

## Stability

Note that the API is currently in `v0`. I don't plan on any huge changes, but there may be some small breaking changes before we hit `v1`.

## 💬 Contact

[![Twitter Follow](https://img.shields.io/twitter/follow/wagslane.svg?label=Follow%20Wagslane&style=social)](https://twitter.com/intent/follow?screen_name=wagslane)

Submit an issue (above in the issues tab)

## Transient Dependencies

None.

## 👏 Contributing

I love help! Contribute by forking the repo and opening pull requests. Please ensure that your code passes the existing tests and linting, and write tests to test your changes if applicable.

All pull requests should be submitted to the `main` branch.
