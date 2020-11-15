package humango

const CONFIG = `{
   "en": {
        "name": "en-US",
        "separator": " ",
        "decimal_separator_word": "point",
        "negative_sign": "minus",
        "quantify_points": [1000000000, 1000000, 1000, 100],
        "quantifiers": {
            "100": {"singular": "hundred", "plural": "hundred"},
            "1000": {"singular": "thousand", "plural": "thousand"},
            "1000000": {"singular": "million", "plural": "million"},
            "1000000000": {"singular": "billion", "plural": "billion"}
        },
        "tens": ["", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"],
        "glyphs": ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "forteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"],
        "rules": {},
        "units": {
            "USD": { 
                "unit": {"singular": "dollar", "plural": "dollars"},
                "subunit":  {"singular": "cent", "plural": "cents"}
            },
            "EUR": { 
                "unit": {"singular": "euro", "plural": "euro"},
                "subunit":  {"singular": "eurocent", "plural": "eurocents"}
            },
            "GBP": { 
                "unit": {"singular": "pound", "plural": "pounds"},
                "subunit":  {"singular": "penny", "plural": "pennies"}
            },
            "JPY": { 
                "unit": {"singular": "yen", "plural": "yen"},
                "subunit":  {"singular": "sen", "plural": "sen"}
            },
            "PLN": { 
                "unit": {"singular": "zloty", "plural": "zlotys"},
                "subunit":  {"singular": "groszy", "plural": "groszys"}
            }
        }
    },
    "de": {
        "name": "de-DE",
        "separator": "",
        "decimal_separator_word": " Komma ",
        "negative_sign": "minus ",
        "quantify_points": [1000000000, 1000000, 1000, 100],
        "quantifiers": {
            "100": {"singular": "hundert", "plural": "hundert"},
            "1000": {"singular": "tausend", "plural": "tausend"},
            "1000000": {"singular": " Million ", "plural": " Millionen "},
            "1000000000": {"singular": " Milliarde ", "plural": " Milliarden "}
        },
        "tens": ["", "", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"],
        "glyphs": ["null", "ein", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechszehn", "siebzehn", "achtzehn", "neunzehn"],
        "rules": {
            "agglunative_tens": {
                "name": "agglunative_tens",
                "context": {
                    "joiner_word":"und"
                }
            }
        },
        "units": {
            "USD":{ 
                "unit": {"singular": " Dollar ", "plural": " Dollar "},
                "subunit":  {"singular": " Cent ", "plural": " Cent"}
            },
            "GBP": { 
                "unit": {"singular": " Pfund ", "plural": " Pfund "},
                "subunit":  {"singular": " Penny ", "plural": " Penny"}
            },
            "EUR": { 
                "unit": {"singular": " Euro ", "plural": " Euro "},
                "subunit":  {"singular": " Eurocent ", "plural": " Eurocent"}
            },
            "JPY": { 
                "unit": {"singular": " Yen ", "plural": " Yen "},
                "subunit":  {"singular": " Sen ", "plural": " Sen"}
            },
            "PLN": { 
                "unit": {"singular": " Zloty ", "plural": " Zloty "},
                "subunit":  {"singular": " Grosch ", "plural": " Groschen"}
            }
        }
    },
    "jp": {
        "name": "jp-JP",
        "separator": "",
        "decimal_separator_word": ".",
        "negative_sign": "-",
        "quantify_points": [10000, 1000, 100],
        "quantifiers": {
            "100": {"singular": "百", "plural": "百"},
            "1000": {"singular": "千", "plural": "千"},
            "10000": {"singular": "万", "plural": "万"} 
        },
        "tens": ["", "", "二十", "三十", "四十", "五十", "六十", "七十", "八十", "九十"],
        "glyphs": ["零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九"],
        "rules": {},
        "units": {
            "USD": { 
                "unit": {"singular": "ドル", "plural": "ドル"},
                "subunit":  {"singular": "セント", "plural": "セント"}
            },
            "GBP": { 
                "unit": {"singular": "ポンド", "plural": "ポンド"},
                "subunit":  {"singular": "ペニー", "plural": "ペニー"}
            },
            "EUR": { 
                "unit": {"singular": "ユーロ", "plural": "ユーロ"},
                "subunit":  {"singular": "ユーロセント", "plural": "ユーロセント"}
            }, 
            "JPY": { 
                "unit": {"singular": "円", "plural": "円"},
                "subunit":  {"singular": "銭", "plural": "銭"}
            },
            "PLN": { 
                "unit": {"singular": "個のzlotys", "plural": "個のzlotys"},
                "subunit":  {"singular": "セント", "plural": "セント"}
            }
        }
    },
    "pl": {
        "name": "pl-PL",
        "separator": " ",
        "decimal_separator_word": ",",
        "negative_sign": "minus",
        "quantify_points": [1000000000, 1000000, 1000, 100],
        "quantifiers": {
            "100": {"singular": "sto", "plural": "-"},
            "1000": {"singular": "tysiąc", "plural": "tysięcy"},
            "1000000": {"singular": "milion", "plural": "milionów"},
            "1000000000": {"singular": "miliard", "plural": "miliardów"}
        },
        "tens": ["", "", "dwadzieścia", "trzydzieści", "czterdzieści", "pięćdziesiąt", "sześćdziesiąt", "siedemdziesiąt", "osiemdziesiąt", "dziewięćdziesiąt"],
        "glyphs": ["zero", "jeden", "dwa", "trzy", "cztery", "pięć", "sześć", "siedem", "osiem", "dziewięć", "dziesięć", "jedenaście", "dwanaście", "trzynaście", "czternaście", "piętnaście", "szesnaście", "siedemnaście", "osiemnaście", "dziewiętnaście"],
        "rules": {
            "custom_hundreds": {
                "name": "custom_hundreds",
                "context": {
                    "hundreds": ["", "sto", "dwieście", "trzysta", "czterysta", "pięćset", "sześćset", "siedemset", "osiemset", "dziewięćset"]
                }
            },
            "slavic_plural": {
                "name": "slavic_plural",
                "context": {
                    "quantifiers": {"1000": "tysiące", "1000000": "miliony", "1000000000": "miliardy"}
                }
            }
        },
        "units": {
            "USD": { 
                "unit": {"singular": "dolar", "plural": "dolarów"},
                "subunit":  {"singular": "cent", "plural": "centów"}
            },
            "GBP": { 
                "unit": {"singular": "funt", "plural": "funtów"},
                "subunit":  {"singular": "pens", "plural": "pensów"}
            },
            "EUR": { 
                "unit": {"singular": "euro", "plural": "euro"},
                "subunit":  {"singular": "eurocent", "plural": "eurocentów"}
            },
            "JPY": { 
                "unit": {"singular": "jen", "plural": "jenów"},
                "subunit":  {"singular": "sen", "plural": "senów"}
            },
            "PLN": { 
                "unit": {"singular": "złoty", "plural": "złotych"},
                "subunit":  {"singular": "grosz", "plural": "groszy"}
            }
        }
    }
}`
