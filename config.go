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
            "USD": {"singular": "dollar", "plural": "dollars"},
            "EUR": {"singular": "euro", "plural": "euro"},
            "GBP": {"singular": "pound", "plural": "pounds"},
            "JPY": {"singular": "yen", "plural": "yen"}
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
            "USD": {"singular": " Dollar", "plural": " Dollar"},
            "GBP": {"singular": " Pfund", "plural": " Pfund"},
            "EUR": {"singular": " Euro", "plural": " Euro"},
            "JPY": {"singular": " Yen", "plural": " Yen"}
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
            "USD": {"singular": "ドル", "plural": "ドル"},
            "GBP": {"singular": "ポンド", "plural": "ポンド"},
            "EUR": {"singular": "ユーロ", "plural": "ユーロ"},
            "JPY": {"singular": "円", "plural": "円"}
        }
    }
}`
