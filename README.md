## Property Filter Command Line

This CLI allows filtering properties based on different criteria.
We offer multiple flags to let you filter properties in various ways.

## Executing the CLI

First, navigate to `app/` directory before executing the program:

```sh
cd app/
```

Then, check if the executable is created. If not, run:

```sh
    go build -o prop-filter-cli
```

If you don't have permission, execute:

```sh
    chmod +x prop-filter-cli
```

## Example

Run the command with desired filters:

```sh
    ./prop-filter-cli -min_price 100000 -max_price 500000 -min_rooms 2 -max_rooms 5
```

## Available flags for filtering

| Flag                   | Type    | Default Value       | Description |
|------------------------|---------|---------------------|-------------|
| `-min_square_footage`  | `int`   | `0`                 | minimum square footage. |
| `-max_square_footage`  | `int`   | `100000`            | maximum square footage. |
| `-lighting_intensity`  | `string`| `""`                | lighting level (`low`, `medium`, `high`). |
| `-min_price`           | `float` | `1000`              | minimum price. |
| `-max_price`           | `float` | `1000000000`        | maximum price. |
| `-min_rooms`           | `int`   | `1`                 | minimum number of rooms. |
| `-max_rooms`           | `int`   | `200`               | maximum number of rooms. |
| `-min_bathrooms`       | `int`   | `1`                 | minimum number of bathrooms. |
| `-max_bathrooms`       | `int`   | `100`               | maximum number of bathrooms. |
| `-description`         | `string`| `""`                | search by description. |
| `-ammenities`          | `string`| `""`                | search by ammenities. |
| `-latitude`            | `float` | `0.0`               | search by latitude. |
| `-longitude`           | `float` | `0.0`               | search by longitude. |
| `-max_distance`        | `float` | `0.0`               | search by max distance. |

