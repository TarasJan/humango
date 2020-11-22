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

humango.WordifyWithLocale(50601, "es")
// Returns "cincuenta mil seiscientos uno"
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

# Supported languages
1. English [en]
2. Chinese [cn] 中文
3. German [de] Deutsch
4. Japanese [jp] 日本語
5. Polish [pl] Polski
6. Russian [ru] Русский
7. Spanish [es] Español

 
# Supported units

\*Currently supported units are dollars `USD`, euro `EUR`, pounds `GBP`, yen `JPY`, zloty `PLN`, `CNY` yuan, and `RUB` rubles and `MXN` pesos)

# Upcoming Features

Upcoming features may be introduced in 3 ways:

1. If you want to introduce the functionality yourself just fork the repo and make a PR for the requested functionality

2. Raise an issue on the repo

3. Contact me directly under jan.taras29@gmail.com

Handling of Decimal numbers and strings is planned.

# Licensing

The software is licensed under MIT License.
