# nserial

Go library to dump data from Nikon F5, F6, F100, F90, N90

## Disclaimer

This project is currently a proof of concept. It can dump data, but it is not accurate.
There is absolutely a possibility that you seriously fuck up your camera trying this,
only do this if you know what you are doing.

I have only tested this on a Nikon F5 because it's the only supported camera I own.
It *should* work on the other cameras as well, but it might need adjustments.
Feel free to open an issue here on GitHub.

## Connection

Any TTL level (3.3v) UART should work.
I have seen people using 5v, it might work, but it might also destroy your camera.
3.3v definitely works.
I used LED pins to get into the connector.

![F5 example](assets/f5_example.JPG)
