# wthr

A command-line tool that outputs the current weather conditions and
forecast.

wthr is still in-development.

wthr uses [OpenWeatherMap](http://openweathermap.org) for weather data.

## Usage

Use the wthr command to output the current climate conditions:

```
wthr --city Portland
```

## Configuration

The OpenWeatherMap API Key is loaded through `os.Getenv("WTHR_OWM_KEY")`
so be sure to make that value accessible to the environment with:

```
export WTHR_OWM_KEY=YOUR_KEY
```
