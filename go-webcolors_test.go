package main

import "testing"

func TestNormalizeHex(t *testing.T) {
    Value := NormalizeHex("#0099CC")
    if Value != "#0099cc" {
        t.Error("Expected #0099cc, got", Value)
    }
}

func TestNormalizeIntegerTriplet(t *testing.T) {
    Value := NormalizeIntegerTriplet([]int{270, -20, 128})
    Expected := []int{255, 0, 128}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}

func TestNormalizeIntegerRGB(t *testing.T) {
    Value := NormalizeIntegerRGB(270)
    if Value != 255 {
        t.Error("Expected 255, got", Value)
    }
}

func TestNormalizePercentTriplet(t *testing.T) {
    Value := NormalizePercentTriplet([]string{"-10%", "250%", "500%"})
    Expected := []string{"0%", "100%", "100%"}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}

func TestNormalizePercentRGB(t *testing.T) {
    Value := NormalizePercentRGB("-5%")
    if Value != "0%" {
        t.Error("Expected 0%, got", Value)
    }
}

func TestNamesToHex(t *testing.T) {
    Value := NamesToHex("white", "css3")
    if Value != "#ffffff" {
        t.Error("Expected white, got", Value)
    }
}

func TestNameToRGB(t *testing.T) {
    Value := NameToRGB("navy", "css3")
    Expected := []int{0, 0, 128}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}

func TestNameToRGBPercent(t *testing.T) {
    Value := NameToRGBPercent("navy", "css3")
    Expected := []string{"0%", "0%", "50%"}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}

func TestHexToName(t *testing.T) {
    Value := HexToName("#daa520", "css3")
    if Value != "goldenrod" {
        t.Error("Expected goldenrod, got", Value)
    }
}

func TestHexToRGB(t *testing.T) {
    Value := HexToRGB("#000080")
    Expected := []int{0, 0, 128}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}

func TestHexToRGBPercent(t *testing.T) {
    Value := HexToRGBPercent("#000080")
    Expected := []string{"0%", "0%", "50%"}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}

func TestRGBToName(t *testing.T) {
    Value := RGBToName([]int{0, 0, 128}, "css3")
    if Value != "navy" {
        t.Error("Expected navy, got", Value)
    }
}

func TestRGBToHex(t *testing.T) {
    Value := RGBToHex([]int{0, 0, 128})
    if Value != "#000080" {
        t.Error("Expected #000080, got", Value)
    }
}

func TestRGBToRGBPercent(t *testing.T) {
    Value := RGBToRGBPercent([]int{218, 165, 32})
    Expected := []string{"85.49%", "64.71%", "12.50%"}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}

func TestRGBPercentToName(t *testing.T) {
    Value := RGBPercentToName([]string{"85.49%", "64.71%", "12.5%"}, "css3")
    if Value != "goldenrod" {
        t.Error("Expected goldenrod, got", Value)
    }
}

func TestRGBPercentToHex(t *testing.T) {
    Value := RGBPercentToHex([]string{"100%", "100%", "0%"})
    if Value != "#ffff00" {
        t.Error("Expected #ffff00, got", Value)
    }
}

func TestRGBPercentToRGB(t *testing.T) {
    Value := RGBPercentToRGB([]string{"0%", "0%", "50%"})
    Expected := []int{0, 0, 128}
    for i := range Value {
        if Value[i] != Expected[i] {
            t.Error("Expected", Expected[i], " got", Value[i])
        }
    }
}
