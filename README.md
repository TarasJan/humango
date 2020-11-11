# Wordify [![Build Status](https://travis-ci.com/TarasJan/wordify.svg?branch=master)](https://travis-ci.com/TarasJan/wordify)
Wordify is a simple Go library for converting numeric input into human readable text. 

It can be used for humanizing currencies, ledger items, stock or other items that require numbers written in plaintext.

# Usage

The library exposes the single method `Wordify` that handles whole numbers and floats up to the second decimal point.

```go
import github.com/TarasJan/wordify

wordify.Wordify(2020)
// Returns "two thousand twenty"

wordify.Wordify(-771)
// Returns "minus seven hundred seventy one"

wordify.Wordify(42.94)
// Returns "forty two point ninety four"

```

# Upcoming Features

Currently only English numerals are supported though other languages are in the works. Future support for formatting currencies and common units akin to Rails' [number_to_human](https://apidock.com/rails/v5.2.3/ActionView/Helpers/NumberHelper/number_to_human) is planned.

The software is licensed under MIT License.
