package hacknet

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type ThemeLayout string

const (
	Blue         ThemeLayout = "blue"
	Green        ThemeLayout = "green"
	White        ThemeLayout = "white"
	Mint         ThemeLayout = "mint"
	GreenCompact ThemeLayout = "greencompact"
	RipTide      ThemeLayout = "riptide"
	Colamaeleon  ThemeLayout = "colamaeleon"
	RipTide2     ThemeLayout = "riptide2"
)

type RgbColour []int

//type RgbaColour [4]int

func (rgb *RgbColour) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	parts := strings.Split(raw, ",")

	*rgb = make([]int, len(parts))

	for i, p := range parts {
		pInt, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return err
		}
		(*rgb)[i] = pInt
	}

	return nil
}

func (rgb RgbColour) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	sb := strings.Builder{}
	for i, v := range rgb {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.Itoa(v))
	}

	return e.EncodeElement(sb.String(), start)
}

type Theme struct {
	Layout          ThemeLayout `xml:"themeLayoutName"`
	BackgroundPath  string      `xml:"backgroundImagePath"`
	PreserveScaling bool        `xml:"UseAspectPreserveBackgroundScaling"`
	BackgroundFill  RgbColour   `xml:"BackgroundImageFillColor"`
	*MainColours
	*ExeColours
	*ExtraColours
}

/*
<!-- Main Colors - these will define the main feel of the theme -->
  <!-- Color of nodes on the netmap, and many other derived colors. -->
  <defaultHighlightColor>255,41,63</defaultHighlightColor>
  <defaultTopBarColor>74,7,14,255</defaultTopBarColor>
  <!-- This is used for the outlines of the module windows -->
  <moduleColorSolidDefault>0,204,132</moduleColorSolidDefault>
  <moduleColorStrong>14,40,25,80</moduleColorStrong>
  <moduleColorBacking>5,7,6,10</moduleColorBacking>
*/

type MainColours struct {
	DefaultHighlightColor   RgbColour `xml:"defaultHighlightColor"`
	DefaultTopBarColor      RgbColour `xml:"defaultTopBarColor"`
	ModuleColorSolidDefault RgbColour `xml:"moduleColorSolidDefault"`
	ModuleColorStrong       RgbColour `xml:"moduleColorStrong"`
	ModuleColorBacking      RgbColour `xml:"moduleColorBacking"`
}

type ExeColours struct {
	ExeModuleTopBar    RgbColour `xml:"exeModuleTopBar"`
	ExeModuleTitleText RgbColour `xml:"exeModuleTitleText"`
}

type ExtraColours struct {
	WarningColor            RgbColour `xml:"warningColor"`
	SubtleTextColor         RgbColour `xml:"subtleTextColor"`
	DarkBackgroundColor     RgbColour `xml:"darkBackgroundColor"`
	IndentBackgroundColor   RgbColour `xml:"indentBackgroundColor"`
	OutlineColor            RgbColour `xml:"outlineColor"`
	LockedColor             RgbColour `xml:"lockedColor"`
	BrightLockedColor       RgbColour `xml:"brightLockedColor"`
	BrightUnlockedColor     RgbColour `xml:"brightUnlockedColor"`
	UnlockedColor           RgbColour `xml:"unlockedColor"`
	LightGray               RgbColour `xml:"lightGray"`
	ShellColor              RgbColour `xml:"shellColor"`
	ShellButtonColor        RgbColour `xml:"shellButtonColor"`
	SemiTransText           RgbColour `xml:"semiTransText"`
	TerminalTextColor       RgbColour `xml:"terminalTextColor"`
	TopBarTextColor         RgbColour `xml:"topBarTextColor"`
	SuperLightWhite         RgbColour `xml:"superLightWhite"`
	ConnectedNodeHighlight  RgbColour `xml:"connectedNodeHighlight"`
	NetmapToolTipColor      RgbColour `xml:"netmapToolTipColor"`
	NetmapToolTipBackground RgbColour `xml:"netmapToolTipBackground"`
	TopBarIconsColor        RgbColour `xml:"topBarIconsColor"`
	ThisComputerNode        RgbColour `xml:"thisComputerNode"`
	ScanlinesColor          RgbColour `xml:"scanlinesColor"`
}
