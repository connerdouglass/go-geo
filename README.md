# go-geo

Geographic coordinates and distances library written in Go.

## Representing coordinates

```go
myLocation := geo.Coordinate{
    Latitude: 39.7684426,
    Longitude: -86.1583112,
}
```

## Calculating distances

```go
a := geo.Coordinate{
    Latitude: 39.7684426,
    Longitude: -86.1583112,
}
b := geo.Coordinate{
    Latitude: 38.3531417,
    Longitude: -81.6388193,
}

// Two ways of calculating the distance:
distance := a.DistanceTo(b)
distance := geo.DistanceBetween(a, b)

fmt.Println("Distance is: ", distance)

// Prints:
// --> 420.6896449044601km
```

*<small>(Note: Ending up with a 420.69km example was a complete accident but I'm not changing it)</small>*

## Working with `geo.Distance`

The `geo.Distance` type represents distances in geographical space. Some predefined constants can be used for convenience:

```go
geo.Millimeter
geo.Centimeter
geo.Meter
geo.Kilometer
geo.Mile
```

You can construct distances like so:

```go
twoMiles := 2 * geo.Mile
tenKm := 10 * geo.Kilometer
```

You can also fetch the `float64` conversion of a distance represented in any units you like:

```go
// Construct a distance (type geo.Distance)
distance := 15 * geo.Mile

// Get the equivalent number of kilometers (type float64)
inKilometers := distance.Kilometers()
```

In the above example, the value of `inKilometers` is equal to `24.14015969842181`, which is the number of kilometers equal to 15 miles.
