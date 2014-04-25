package webcolors

import (
    "math"
    "regexp"
    "strings"
    "strconv"
    "encoding/hex"
    "encoding/binary"
)

var HexColorRegex = regexp.MustCompile(`^#([a-fA-F0-9]{3}|[a-fA-F0-9]{6})$`)

var SupportedSpecifications = []string{"html4", "css2", "css21", "css3"}

var HTML4NamesToHex = map[string]string {
    "aqua": "#00ffff",
    "black": "#000000",
    "blue": "#0000ff",
    "fuchsia": "#ff00ff",
    "green": "#008000",
    "grey": "#808080",
    "lime": "#00ff00",
    "maroon": "#800000",
    "navy": "#000080",
    "olive": "#808000",
    "purple": "#800080",
    "red": "#ff0000",
    "silver": "#c0c0c0",
    "teal": "#008080",
    "white": "#ffffff",
    "yellow": "#ffff00",
    "orange": "#ffa500",
}

var CSS2NamesToHex = HTML4NamesToHex

var CSS21NamesToHex = HTML4NamesToHex

var CSS3NamesToHex = map[string]string {
    "aliceblue": "#f0f8ff",
    "antiquewhite": "#faebd7",
    "aqua": "#00ffff",
    "aquamarine": "#7fffd4",
    "azure": "#f0ffff",
    "beige": "#f5f5dc",
    "bisque": "#ffe4c4",
    "black": "#000000",
    "blanchedalmond": "#ffebcd",
    "blue": "#0000ff",
    "blueviolet": "#8a2be2",
    "brown": "#a52a2a",
    "burlywood": "#deb887",
    "cadetblue": "#5f9ea0",
    "chartreuse": "#7fff00",
    "chocolate": "#d2691e",
    "coral": "#ff7f50",
    "cornflowerblue": "#6495ed",
    "cornsilk": "#fff8dc",
    "crimson": "#dc143c",
    "cyan": "#00ffff",
    "darkblue": "#00008b",
    "darkcyan": "#008b8b",
    "darkgoldenrod": "#b8860b",
    "darkgray": "#a9a9a9",
    "darkgrey": "#a9a9a9",
    "darkgreen": "#006400",
    "darkkhaki": "#bdb76b",
    "darkmagenta": "#8b008b",
    "darkolivegreen": "#556b2f",
    "darkorange": "#ff8c00",
    "darkorchid": "#9932cc",
    "darkred": "#8b0000",
    "darksalmon": "#e9967a",
    "darkseagreen": "#8fbc8f",
    "darkslateblue": "#483d8b",
    "darkslategray": "#2f4f4f",
    "darkslategrey": "#2f4f4f",
    "darkturquoise": "#00ced1",
    "darkviolet": "#9400d3",
    "deeppink": "#ff1493",
    "deepskyblue": "#00bfff",
    "dimgray": "#696969",
    "dimgrey": "#696969",
    "dodgerblue": "#1e90ff",
    "firebrick": "#b22222",
    "floralwhite": "#fffaf0",
    "forestgreen": "#228b22",
    "fuchsia": "#ff00ff",
    "gainsboro": "#dcdcdc",
    "ghostwhite": "#f8f8ff",
    "gold": "#ffd700",
    "goldenrod": "#daa520",
    "gray": "#808080",
    "grey": "#808080",
    "green": "#008000",
    "greenyellow": "#adff2f",
    "honeydew": "#f0fff0",
    "hotpink": "#ff69b4",
    "indianred": "#cd5c5c",
    "indigo": "#4b0082",
    "ivory": "#fffff0",
    "khaki": "#f0e68c",
    "lavender": "#e6e6fa",
    "lavenderblush": "#fff0f5",
    "lawngreen": "#7cfc00",
    "lemonchiffon": "#fffacd",
    "lightblue": "#add8e6",
    "lightcoral": "#f08080",
    "lightcyan": "#e0ffff",
    "lightgoldenrodyellow": "#fafad2",
    "lightgray": "#d3d3d3",
    "lightgrey": "#d3d3d3",
    "lightgreen": "#90ee90",
    "lightpink": "#ffb6c1",
    "lightsalmon": "#ffa07a",
    "lightseagreen": "#20b2aa",
    "lightskyblue": "#87cefa",
    "lightslategray": "#778899",
    "lightslategrey": "#778899",
    "lightsteelblue": "#b0c4de",
    "lightyellow": "#ffffe0",
    "lime": "#00ff00",
    "limegreen": "#32cd32",
    "linen": "#faf0e6",
    "magenta": "#ff00ff",
    "maroon": "#800000",
    "mediumaquamarine": "#66cdaa",
    "mediumblue": "#0000cd",
    "mediumorchid": "#ba55d3",
    "mediumpurple": "#9370d8",
    "mediumseagreen": "#3cb371",
    "mediumslateblue": "#7b68ee",
    "mediumspringgreen": "#00fa9a",
    "mediumturquoise": "#48d1cc",
    "mediumvioletred": "#c71585",
    "midnightblue": "#191970",
    "mintcream": "#f5fffa",
    "mistyrose": "#ffe4e1",
    "moccasin": "#ffe4b5",
    "navajowhite": "#ffdead",
    "navy": "#000080",
    "oldlace": "#fdf5e6",
    "olive": "#808000",
    "olivedrab": "#6b8e23",
    "orange": "#ffa500",
    "orangered": "#ff4500",
    "orchid": "#da70d6",
    "palegoldenrod": "#eee8aa",
    "palegreen": "#98fb98",
    "paleturquoise": "#afeeee",
    "palevioletred": "#d87093",
    "papayawhip": "#ffefd5",
    "peachpuff": "#ffdab9",
    "peru": "#cd853f",
    "pink": "#ffc0cb",
    "plum": "#dda0dd",
    "powderblue": "#b0e0e6",
    "purple": "#800080",
    "red": "#ff0000",
    "rosybrown": "#bc8f8f",
    "royalblue": "#4169e1",
    "saddlebrown": "#8b4513",
    "salmon": "#fa8072",
    "sandybrown": "#f4a460",
    "seagreen": "#2e8b57",
    "seashell": "#fff5ee",
    "sienna": "#a0522d",
    "silver": "#c0c0c0",
    "skyblue": "#87ceeb",
    "slateblue": "#6a5acd",
    "slategray": "#708090",
    "slategrey": "#708090",
    "snow": "#fffafa",
    "springgreen": "#00ff7f",
    "steelblue": "#4682b4",
    "tan": "#d2b48c",
    "teal": "#008080",
    "thistle": "#d8bfd8",
    "tomato": "#ff6347",
    "turquoise": "#40e0d0",
    "violet": "#ee82ee",
    "wheat": "#f5deb3",
    "white": "#ffffff",
    "whitesmoke": "#f5f5f5",
    "yellow": "#ffff00",
    "yellowgreen": "#9acd32",
}

func ReverseMap(m map[string]string) map[string]string {
    n := make(map[string]string)
    for k, v := range m {
        n[v] = k
    }
    return n
}

// # Mappings of Normalized hexadecimal color Values to color Names.
// #################################################################

var HTML4HexToNames = ReverseMap(HTML4NamesToHex)
var CSS2HexToNames = HTML4HexToNames
var CSS21HexToNames = ReverseMap(CSS21NamesToHex)
var CSS3HexToNames = ReverseMap(CSS3NamesToHex)

// Normalization routines.
// #################################################################

// Normalize a hexadecimal color value
func NormalizeHex(HexValue string) string {
    HexDigits := HexColorRegex.FindStringSubmatch(HexValue)[1]
    if len(HexDigits) == 3 {
        finalhex := []string{}
        for i := range HexDigits {
            finalhex = append(finalhex, strings.Repeat(string(HexDigits[i]), 2))
        }
        return "#" + strings.ToLower(strings.Join(finalhex, ""))
    } else {
        return "#" + strings.ToLower(HexDigits)
    }
    return ""
}

// Normalize an integer rgb triplet so that all values are within the range 0-255 inclusive.
func NormalizeIntegerTriplet(RGBTriplet []int) []int {
    IntegerTriplet := []int{NormalizeIntegerRGB(RGBTriplet[0]), NormalizeIntegerRGB(RGBTriplet[1]), NormalizeIntegerRGB(RGBTriplet[2])}
    return IntegerTriplet
}

// Normalize value for use in an integer rgb triplet
func NormalizeIntegerRGB(Value int) int {
    if Value >= 0 && Value <= 255 {
        return Value
    } else if Value < 0 {
        return 0
    } else if Value > 255 {
        return 255
    }
    return 0
}

// Normalize a percentage rgb triplet to that all values are within the range 0%-100% inclusive.
func NormalizePercentTriplet(RGBTriplet []string) []string {
    FinalTriplet := []string{}
    for i := range RGBTriplet {
        FinalTriplet = append(FinalTriplet, NormalizePercentRGB(RGBTriplet[i]))
    }
    return FinalTriplet
}

// Normalize value for use in a percentage rgb triplet
func NormalizePercentRGB(Value string) string{
    var Percent = strings.Split(Value, "%")[0]

    if strings.Contains(Percent, ".") {
        Percent, err := strconv.ParseFloat(Percent, 64)
        if err == nil {
            if Percent >= 0 && Percent <= 100{
                return strconv.FormatFloat(Percent, 'g', 4, 64) + "%"
            } else if Percent < 0 {
                return "0%"
            } else if Percent > 100 {
                return "100%"
            }
        }
    } else {
        Percent, err := strconv.Atoi(Percent)
        if err == nil {
            if Percent >= 0 && Percent <= 100{
                return strconv.Itoa(Percent) + "%"
            } else if Percent < 0 {
                return "0%"
            } else if Percent > 100 {
                return "100%"
            }
        }
    }
    return ""
}

//# Conversions from color Names to various formats.
// #################################################################

func Contains(s []string, e string) bool {
    for _, output := range s { if output == e { return true } }
    return false
}

// Convert a color name to a normalized hexadecimal color value
func NamesToHex(Name string, Spec string) string {
    if Contains(SupportedSpecifications, Spec) == true {
        Normalized := strings.ToLower(Name)
        if Spec == "html4" {
            return HTML4NamesToHex[Normalized]
        } else if Spec == "css2" {
            return CSS2NamesToHex[Normalized]
        } else if Spec == "css21" {
            return CSS21NamesToHex[Normalized]
        } else if Spec == "css3" {
            return CSS3NamesToHex[Normalized]
        } else {
            panic(Name + "has no defined color Name in " + Spec)
        }
    } else {
        panic(Spec + "is not output supported Specification for color Name lookups")
    }
    return ""
}

// Convert a color name to a 3-tuple of integers suitable for use in an rgb triplet specifying that color
func NameToRGB(Name string, Spec string) []int {
    return HexToRGB(NamesToHex(Name, Spec))
}

// Convert a color name to a 3-tuple of percentages suitable for use in an rgb triplet specifying that color
func NameToRGBPercent(Name string, Spec string) []string {
    return RGBToRGBPercent(NameToRGB(Name, Spec))
}

// # Conversions from hexadecimal color Values to various formats.
// #################################################################

// Convert a hexadecimal color value to its corresponding normalized color name, if any such name exists
func HexToName(HexValue string, Spec string) string {
    if Contains(SupportedSpecifications, Spec) == true {
        Normalized := NormalizeHex(HexValue)
        if Spec == "html4" {
            return HTML4HexToNames[Normalized]
        } else if Spec == "css2" {
            return CSS2HexToNames[Normalized]
        } else if Spec == "css21" {
            return CSS21HexToNames[Normalized]
        } else if Spec == "css3" {
            return CSS3HexToNames[Normalized]
        } else {
            panic(HexValue + "has no defined color Name in " + Spec)
        }
    } else {
        panic(Spec + "is not output supported Specification for color Name lookups")
    }
    return ""
}

func ByteToInt(Input []byte) int {
    var Output uint32
    l := len(Input)
    for i, b := range Input {
        shift := uint32((l-i-1) * 8)
        Output |= uint32(b) << shift
    }
    return int(Output)
}

// Convert a hexadecimal color value to a 3-tuple of integers suitable for use in an rgb triplet specifying that color
func HexToRGB(HexValue string) []int {
    HexDigits := NormalizeHex(HexValue)
    RGBTuple := []int{}
    PartialHex1, err := hex.DecodeString(HexDigits[1:3])
    if err == nil {
        RGBTuple = append(RGBTuple, ByteToInt(PartialHex1))
    }
    PartialHex2, err := hex.DecodeString(HexDigits[3:5])
    if err == nil {
        RGBTuple = append(RGBTuple, ByteToInt(PartialHex2))
    }
    PartialHex3, err := hex.DecodeString(HexDigits[5:7])
    if err == nil {
        RGBTuple = append(RGBTuple, ByteToInt(PartialHex3))
    }
    return RGBTuple
}

// Convert a hexadecimal color value to a 3-tuple of percentages suitable for use in an rgb triplet representing that color
func HexToRGBPercent(HexValue string) []string {
    return RGBToRGBPercent(HexToRGB(HexValue))
}

// # Conversions from  integer rgb() triplets to various formats.
// #################################################################

// Convert a 3-tuple of integers, suitable for use in an rgb color triplet, to its corresponding normalized color name, if any such name exists
func RGBToName(RGBTriplet []int, Spec string) string {
    return HexToName(RGBToHex(NormalizeIntegerTriplet(RGBTriplet)), Spec)
}

// Convert a 3-tuple of integers, suitable for use in an rgb color triplet, to a normalized hexadecimal value for that color
func RGBToHex(RGBTriplet []int) string {
    IntegerTriplet := NormalizeIntegerTriplet(RGBTriplet)
    HexString := "#"
    for i := range IntegerTriplet {
        ByteCoded := make([]byte, 2)
        binary.BigEndian.PutUint16(ByteCoded, uint16(IntegerTriplet[i]))
        HexString = HexString +  hex.EncodeToString(ByteCoded[1:2])
    }
    return HexString
}

// Convert a 3-tuple of integers, suitable for use in an rgb color triplet, to a 3-tuple of percentages suitable for use in representing that color
func RGBToRGBPercent(RGBTriplet []int) []string {
    Specials := map[int] string {
        255: "100%",
        128: "50%",
        64: "25%",
        32: "12.50%",
        16: "6.25%",
        0: "0%",
    }
    RGBPercentTriplet := []string{}
    NormalizedTriplet := NormalizeIntegerTriplet(RGBTriplet)

    for i := range NormalizedTriplet {
        if Name, ok := Specials[NormalizedTriplet[i]]; ok {
            RGBPercentTriplet = append(RGBPercentTriplet, Name)
        } else {
            PercentVal := (float64(NormalizedTriplet[i]) / 255.0) * 100
            RGBPercentTriplet = append(RGBPercentTriplet, strconv.FormatFloat(PercentVal, 'g', 4, 64) + "%")
        }
    }
    return RGBPercentTriplet
}

// # Conversions from Percentage rgb() triplets to various formats.
// #################################################################

// Convert a 3-tuple of percentages, suitable for use in an rgb color triplet, to its corresponding normalized color name, if any such name exists
func RGBPercentToName(RGBPercentTriplet []string, Spec string) string {
    return RGBToName(RGBPercentToRGB(NormalizePercentTriplet(RGBPercentTriplet)), Spec)
}

// Convert a 3-tuple of percentages, suitable for use in an rgb color triplet, to a normalized hexadecimal color value for that color
func RGBPercentToHex(RGBPercentTriplet []string) string {
    return RGBToHex(RGBPercentToRGB(NormalizePercentTriplet(RGBPercentTriplet)))
}

// Internal helper for converting a percentage value to an integer between 0 and 255 inclusive
func PercentToInteger(Percent string) int {
    Num, err := strconv.ParseFloat(strings.Split(Percent, "%")[0], 64)
    if err == nil {
        Num = 255 * (Num / 100.0)
        e := Num - math.Floor(Num)
        if e < 0.5 {
            return int(math.Floor(Num))
        } else {
            return int(math.Ceil(Num))
        }
    }
    return 0
}

// Convert a 3-tuple of percentages, suitable for use in an rgb color triplet, to a 3-tuple of integers suitable for use in representing that color
func RGBPercentToRGB(RGBPercentTriplet []string) []int {
    RGBTriplet := []int{}
    NormalizedTriplet := NormalizePercentTriplet(RGBPercentTriplet)
    for i := range NormalizedTriplet {
        RGBTriplet = append(RGBTriplet, PercentToInteger(NormalizedTriplet[i]))
    }
    return RGBTriplet
}
