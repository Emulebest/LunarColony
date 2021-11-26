# Lunar Colony time converters

## Description

This package is responsible for converting Earth UTC date to a Lunar Standard Time (LST).
For detailed description on what is the Lunar Standard Time refer to http://lunarclock.org/what-is-lunar-standard-time.php

## Algorithm basis

The algorithm for conversion is based on the assumption that:
a day, counting from noon to noon, lasts about 29.27 to 29.83 Earth days. It is not a constant. The mean value, roughly 29.530589, is not a constant either! Over time it will be longer.
The solution for this piece is to define the moon-second as 29.530589/30, which means that **for every 30 seconds of Earth time
there are 29.530589 seconds of Moon time.**

Conversion can be done knowing that:
The Lunar year consists of twelve days, named after the first men who walked on the Moon. Each day is divided into 30 cycles of time,
with each cycle being divided into 24 moon-hours. Each moon-hour then has 60 moon-minutes,
which in turn of course are made up of 60 moon-seconds each.

The standard notation is: Year-Day-Cycle ∇ Hour:Minute:Second

## Algorithm steps

1) Convert the Moon starting date (Neil Armstrong set foot on the Moon surface on July 21th 1969 at 02:56:15 UT) and the target date to UTC
2) Subtract starting date from target date
3) Convert the difference to seconds (These are the seconds from '01-01-01 delta 00:00:00 LST')
4) Transform seconds to the LST (while transforming add 1 years, 1 day, 1 cycle since the start date is from 1 and not 0)