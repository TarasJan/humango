# humango [![Build Status](https://travis-ci.com/TarasJan/humango.svg?branch=master)](https://travis-ci.com/TarasJan/humango)
Humango is a simple Go library for converting numeric input into human readable text. 

It can be used for humanizing currencies, ledger items, stock or other items that require numbers written in plaintext.

# Usage

The library exposes method `Wordify` that handles whole numbers and floats up to the second decimal point. Wordify will produce a text of English numerals up to billions.

```go
import github.com/TarasJan/humango

humango.Wordify(2020)
// Returns "two thousand twenty"

humango.Wordify(-771)
// Returns "minus seven hundred seventy one"

humango.Wordify(42.94)
// Returns "forty two point ninety four"

```

Should you want to wordify the numbers in other languages the method `WordifyWithLocale` may come in handy:
```go
humango.WordifyWithLocale(50601, "en")
// Returns "fifty thousand six hundred one" - equivalent to humango.Wordify(50601)

humango.WordifyWithLocale(50601, "jp")
// Returns "五万六百一"

humango.WordifyWithLocale(50601, "de")
// Returns "funfzigtausendsechshundertein"

humango.WordifyWithLocale(50601, "pl")
// Returns "pięćdziesiąt tysięcy sześćset jeden"
```

If you would like to add units into the mix the `WordifyWithLocaleAndUnit` is the method to go for\*:
```go
humango.WordifyWithLocaleAndUnit(50601, "en", "USD")
// Returns "fifty thousand six hundred one dollars"

humango.WordifyWithLocale(50601, "jp", "JPY")
// Returns "五万六百一円"

humango.WordifyWithLocale(50601, "de", "EUR")
// Returns "fünfzigtausendsechshundertein Euro"
```

\*Currently supported units are dollars `USD`, euro `EUR`, pounds `GBP`, yen `JPY` and zloty `PLN`)

# Upcoming Features

Other languages are in the works - Spanish, Italian and Russian are upcoming.
More units spanning currencies, weights and scientific measures will be introduced. 
Handling of Decimal numbers and strings is planned.
# Licensing

The software is licensed under MIT License.
