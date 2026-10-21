# Ticket

A super simple program for opening tickets from CLI.

## Usage

- The [`Makefile`](./Makefile) handles most things to do with compilation, installation & cleanup
  ```sh
  make              # Compile binary locally
  make clean        # Remove local binary
  make install      # Install binary to go path
  make uninstall    # Install binary to go path
  ```
- Binary takes input over stdin & as positional args. Multiple inputs & tickets are supported
  ```sh
  ticket 100                                    # Positional
  echo 100 | ticket                             # stdin
  printf '%d\n' {100..105} | ticket {106..110}  # Both
  ```

## Development

- Run linter locally in a container
  ```sh
  make lint                         # Defaults to `podman`
  make lint CONTAINER_ENGINE=docker # Run linter in docker
  ```
