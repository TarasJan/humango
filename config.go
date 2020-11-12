package humango

const CONFIG = `{
   "en": {
        "name": "en-US",
        "separator": " ",
        "decimal_separator_word": "point",
        "negative_sign": "minus",
        "quantify_points": [1000000000, 1000000, 1000, 100],
        "quantifiers": {"100": "hundred", "1000": "thousand", "1000000": "million", "1000000000":"billion"},
        "tens": ["", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"],
        "glyphs": ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "forteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"]
    },
    "de": {
        "name": "de-DE",
        "separator": " ",
        "decimal_separator_word": "Komma",
        "negative_sign": "minus",
        "quantify_points": [1000000000, 1000000, 1000, 100],
        "quantifiers": {"100": "hundert", "1000": "tausend", "1000000": "Million", "1000000000": "Milliarde"},
        "tens": ["", "", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"],
        "glyphs": ["null", "ein", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechszehn", "siebzehn", "achtzehn", "neunzehn"]
    },
    "jp": {
        "name": "jp-JP",
        "separator": "",
        "decimal_separator_word": ".",
        "negative_sign": "-",
        "quantify_points": [10000, 1000, 100],
        "quantifiers": {"100": "百", "1000": "千", "10000": "万"},
        "tens": ["", "", "二十", "三十", "四十", "五十", "六十", "七十", "八十", "九十"],
        "glyphs": ["零", "一", "二", "三", "四", "五", "六", "七", "八", "九", "十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九"]
    }
}`
