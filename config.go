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
        "glyphs": ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"],
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
            },
            "CNY": { 
                "unit": {"singular": "yuan", "plural": "yuans"},
                "subunit":  {"singular": "fen", "plural": "fens"}
            },
            "RUB": {
                "unit": {
                    "singular": "ruble",
                    "plural": "rubles"
                },
                "subunit": {
                    "singular": "kopek",
                    "plural": "kopeks"
                }
            },
            "MXN": {
                "unit": {
                    "singular": "peso",
                    "plural": "peso"
                },
                "subunit": {
                    "singular": "centavo",
                    "plural": "centavo"
                }
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
                    "joiner_word":"und",
                    "swap": true
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
            },
            "CNY": { 
                "unit": {"singular": "Yuan", "plural": "Yuan"},
                "subunit": {"singular": "Fen", "plural": "Fen"}
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
            },
            "CNY": {
                "unit": {"singular": "元", "plural": "元"},
                "subunit":  {"singular": "分", "plural": "分"}
            }
        }
    },
    "zh": {
        "name": "zh-CN",
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
                "unit": {"singular": "美元", "plural": "美元"},
                "subunit":  {"singular": "美分", "plural": "美分"}
            },
            "GBP": { 
                "unit": {"singular": "英镑", "plural": "英镑"},
                "subunit":  {"singular": "分钱", "plural": "分钱"}
            },
            "EUR": { 
                "unit": {"singular": "欧元", "plural": "欧元"},
                "subunit":  {"singular": "欧分", "plural": "欧分"}
            }, 
            "JPY": { 
                "unit": {"singular": "日元", "plural": "日元"},
                "subunit":  {"singular": "銭", "plural": "銭"}
            },
            "PLN": { 
                "unit": {"singular": "金色的", "plural": "金色的"},
                "subunit":  {"singular": "銭", "plural": "銭"}
            },
            "CNY": {
                "unit": {"singular": "元", "plural": "元"},
                "subunit":  {"singular": "分", "plural": "分"}
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
            "100": {
                "singular": "sto",
                "plural": "-",
                "splural": "-"
            },
            "1000": {
                "singular": "tysiąc",
                "plural": "tysięcy",
                "splural": "tysiące"
            },
            "1000000": {
                "singular": "milion",
                "plural": "milionów",
                "splural": "miliony"
            },
            "1000000000": {
                "singular": "miliard",
                "plural": "miliardów",
                "splural": "miliardy"
            }
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
                "context": {}
            }
        },
        "units": {
            "USD": { 
                "unit": {
                    "singular": "dolar",
                    "plural": "dolarów",
                    "splural": "dolary"
                },
                "subunit":  {
                    "singular": "cent",
                    "plural": "centów",
                    "splural": "centy"
                }
            },
            "GBP": { 
                "unit": {
                    "singular": "funt",
                    "plural": "funtów",
                    "splural": "funty"
                },
                "subunit":  {
                    "singular": "pens",
                    "plural": "pensów",
                    "splural": "pensy"
                }
            },
            "EUR": { 
                "unit": {
                    "singular": "euro",
                    "plural": "euro",
                    "splural": "euro"
                },
                "subunit":  {
                    "singular": "eurocent",
                    "plural": "eurocentów",
                    "splural": "eurocenty"
                }
            },
            "JPY": { 
                "unit": {
                    "singular": "jen",
                    "plural": "jenów",
                    "splural": "jeny"
                },
                "subunit":  {
                    "singular": "sen",
                    "plural": "senów",
                    "splural": "seny"
                }
            },
            "PLN": { 
                "unit": {
                    "singular": "złoty",
                    "plural": "złotych",
                    "splural": "złote"
                },
                "subunit":  {
                    "singular": "grosz",
                    "plural": "groszy",
                    "splural": "grosze"
                }
            },
            "CNY": {
                "unit": {
                    "singular": "juan",
                    "plural": "juanów",
                    "splural": "juany"
                },
                "subunit": {
                    "singular": "fen",
                    "plural": "fenów",
                    "splural": "feny"
                }
            }
        }
    },
    "ru": {
        "name": "ru-RU",
        "separator": " ",
        "decimal_separator_word": ",",
        "negative_sign": "минус",
        "quantify_points": [1000000000, 1000000, 1000, 100],
        "quantifiers": {
            "100": {
                "singular": "сто",
                "plural": "-",
                "splural": "-"
            },
            "1000": {
                "singular": "тысяча",
                "plural": "тысяч",
                "splural": "тысячи"
            },
            "1000000": {
                "singular": "миллион",
                "plural": "миллионов",
                "splural": "миллиона"
            },
            "1000000000": {
                "singular": "миллиард",
                "plural": "миллиардов",
                "splural": "миллиарда"
            }
        },
        "tens": ["", "", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"],
        "glyphs": ["нуль", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять", "десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"],
        "rules": {
            "custom_hundreds": {
                "name": "custom_hundreds",
                "context": {
                    "hundreds": ["", "сto", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"]
                }
            },
            "slavic_plural": {
                "name": "slavic_plural",
                "context": {}
            }
        },
        "units": {
            "USD": { 
                "unit": {
                    "singular": "долларов",
                    "plural": "доллар",
                    "splural": "доллара"
                },
                "subunit":  {
                    "singular": "цент",
                    "plural": "центов",
                    "splural": "цента"
                }
            },
            "GBP": { 
                "unit": {
                    "singular": "фунт",
                    "plural": "фунтов",
                    "splural": "фунта"
                },
                "subunit":  {
                    "singular": "пенс",
                    "plural": "пенсов",
                    "splural": "пенса"
                }
            },
            "EUR": { 
                "unit": {
                    "singular": "евро",
                    "plural": "евро",
                    "splural": "евро"
                },
                "subunit":  {
                    "singular": "евроcent",
                    "plural": "евроцентов",
                    "splural": "евроцента"
                }
            },
            "JPY": { 
                "unit": {
                    "singular": "иен",
                    "plural": "иен",
                    "splural": "иены"
                },
                "subunit":  {
                    "singular": "сен",
                    "plural": "сен",
                    "splural": "сены"
                }
            },
            "PLN": { 
                "unit": {
                    "singular": "злотый",
                    "plural": "злотых",
                    "splural": "злотых"
                },
                "subunit":  {
                    "singular": "цент",
                    "plural": "центов",
                    "splural": "цента"
                }
            },
            "CNY": {
                "unit": {
                    "singular": "юань",
                    "plural": "юаней",
                    "splural": "юаня"
                },
                "subunit": {
                    "singular": "болото",
                    "plural": "болот",
                    "splural": "болота"
                }
            },
            "RUB": {
                "unit": {
                    "singular": "рубль",
                    "plural": "рублей",
                    "splural": "рубля"
                },
                "subunit": {
                    "singular": "копейка",
                    "plural": "копеек",
                    "splural": "копейки"
                }
            }
        }
    },
    "es": {
        "name": "es-ES",
        "separator": " ",
        "decimal_separator_word": ",",
        "negative_sign": "minus",
        "quantify_points": [1000000000, 1000000, 1000, 100],
        "quantifiers": {
            "100": {
                "singular": "cento",
                "plural": "-"
            },
            "1000": {
                "singular": "mil",
                "plural": "mil"
            },
            "1000000": {
                "singular": "millón",
                "plural": "millónes"
            },
            "1000000000": {
                "singular": "millardo",
                "plural": "millardos"
            }
        },
        "tens": ["", "", "", "trenta", "quarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"],
        "glyphs": ["cero", "uno", "dos", "tres", "cuarto", "cinco", "seis", "siete", "ocho", "nueve", "diez", "once", "doce", "trece", "catorce", "quince", "dieciséis", "diecisiete", "dieciocho", "diecinueve", "veinte", "veintiuno", "veintidós", "veintitrés", "veinticuatro", "veinticinco", "veintiséis", "veintisiete", "veintiocho", "veintinueve"],
        "rules": {
            "custom_hundreds": {
                "name": "custom_hundreds",
                "context": {
                    "hundreds": ["", "ciento", "doscientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos"]
                }
            },
            "agglunative_tens": {
                "name": "agglunative_tens",
                "context": {
                    "joiner_word":" y ",
                    "swap": false
                }
            }
        },
        "units": {
            "USD": { 
                "unit": {
                    "singular": "dolar",
                    "plural": "dolares"
                },
                "subunit":  {
                    "singular": "centavo",
                    "plural": "centavos"
                }
            },
            "GBP": { 
                "unit": {
                    "singular": "libra",
                    "plural": "libras"
                },
                "subunit":  {
                    "singular": "peniques",
                    "plural": "peniques"
                }
            },
            "EUR": { 
                "unit": {
                    "singular": "euro",
                    "plural": "euros"
                },
                "subunit":  {
                    "singular": "eurocentavo",
                    "plural": "eurocentavos"
                }
            },
            "JPY": { 
                "unit": {
                    "singular": "yen",
                    "plural": "yenes"
                },
                "subunit":  {
                    "singular": "sen",
                    "plural": "senes"
                }
            },
            "PLN": { 
                "unit": {
                    "singular": "zloty",
                    "plural": "zlotys"
                },
                "subunit":  {
                    "singular": "grosz",
                    "plural": "groszys"
                }
            },
            "CNY": {
                "unit": {
                    "singular": "yuan",
                    "plural": "yuanes"
                },
                "subunit": {
                    "singular": "fen",
                    "plural": "fenes"
                }
            },
            "RUB": {
                "unit": {
                    "singular": "rublo",
                    "plural": "rublos"
                },
                "subunit": {
                    "singular": "kopeks",
                    "plural": "kopeks"
                }
            },
            "MXN": {
                "unit": {
                    "singular": "peso",
                    "plural": "pesos"
                },
                "subunit": {
                    "singular": "centavo",
                    "plural": "centavos"
                }
            }
        }
    }
}`
