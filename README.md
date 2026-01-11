# GoPing

A straightforward command-line utility written in Go for monitoring HTTP endpoints.

## Features

-   **HTTP Status Check:** Quickly checks the HTTP status of one or more URLs using `HEAD` requests.
-   **Continuous Monitoring (`watch` mode):** Periodically pings specified URLs and logs their status.
-   **Logging:** Records monitoring results to time-stamped log files in the `logs/` directory.

## Installation

Ensure you have Go installed.

```bash
git clone https://github.com/andradeatdev/goping.git
cd goping
go build -o goping ./cmd/goping
```

Place the `goping` executable in your system's PATH for easy access.

## Usage

### Ping Once

To check the status of a URL once:

```bash
goping http://example.com
goping https://google.com https://github.com
```

### Watch Mode

To continuously monitor URLs and log their status. The monitoring interval can be configured using the `-interval` flag (e.g., `-interval 5s` for 5-second intervals).

```bash
goping watch http://example.com https://another.com
goping watch -interval 5s http://example.com
```

Logs will be saved in the `logs/` directory.

## Contributing

Contributions are welcome. Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
