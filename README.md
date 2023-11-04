# Gopetuah

## Description

Gopetuah is a lightweight Go application designed to send requests to multiple upstream servers. Its primary purpose is to route TCP packages from one endpoint to multiple upstream servers, making it ideal for data ingestion and traffic distribution.

## Purpose/Usecase

The primary use case for Gopetuah is to efficiently route TCP traffic from a single endpoint to multiple upstream servers. This can be valuable in scenarios where you need to collect data from various sources or distribute incoming traffic across multiple destinations.

## Features

- **Route Any TCP Traffic**: Gopetuah is capable of routing any TCP traffic, allowing you to forward data from one source to multiple destinations with ease.

- **Lightweight**: Built with efficiency in mind, Gopetuah is a lightweight application that won't strain your system resources.

- **Built on Go Standard Library**: Gopetuah leverages the power and reliability of the Go Standard Library to ensure stability and performance.

- **Able to Handle Concurrent Connections**: Gopetuah is designed to manage concurrent connections effectively, ensuring that your data ingestion remains smooth and uninterrupted.

## Getting Started

To get started with Gopetuah, follow these simple steps:

1. [Download the latest release](https://github.com/thaddeuscleo/gopetuah/releases).

2. Configure your upstream servers in the `config.yaml` file.

3. Run Gopetuah using the following command:

   ```
   ./gopetuah proxy
   ```

4. Gopetuah will start routing incoming TCP traffic to your configured upstream servers.

## Configuration

You can configure Gopetuah by editing the `config.yaml` file. Customize the settings to meet your specific requirements, including specifying the endpoints of your upstream servers.

```yaml
proxy:
  host: localhost
  port: 3000

upstreams:
  anyname_upstream1:
    host: localhost
    port: 8000
  anyname_upstream2:
    host: localhost
    port: 9000
```

## License

This project is licensed under the MIT License

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## Contact

For any questions or feedback, please contact the project maintainer at thaddeuscleo@gmail.com.

Thank you for using Gopetuah! We hope it serves your data ingestion needs effectively.