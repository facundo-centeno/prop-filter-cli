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

## Execute

Run the command with desired filters:

Some examples:
```sh
    ./prop-filter-cli -min_price 300000 -max_price 500000
    ./prop-filter-cli -min_rooms 4 -lighting_intensity high
    ./prop-filter-cli -max_distance 50 -min_bathrooms 3
    ./prop-filter-cli -ammenities pool -max_price 350000
    ./prop-filter-cli -description loft -max_rooms 3
```

## Available flags for filtering

| flag                   | type    | default value       | description |
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
| `-max_distance`        | `float` | `0.0`               | search by max distance. |
