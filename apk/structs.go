package apk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/kamushadenes/apkrash/utils"
)

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()

type APKComparison struct {
	APK1        *APK                    `json:"-"`
	APK2        *APK                    `json:"-"`
	PackageName *APKComparisonString    `json:"packageName"`
	Features    *APKComparisonArray     `json:"features"`
	Permissions *APKComparisonArray     `json:"permissions"`
	Services    *APKComparisonArray     `json:"services"`
	Receivers   *APKComparisonArray     `json:"receivers"`
	Activities  *APKComparisonArray     `json:"activities"`
	Providers   *APKComparisonArray     `json:"providers"`
	Files       *APKComparisonFileArray `json:"files"`
	Sources     *APKComparisonFileArray `json:"sources"`
}

type APKComparisonString struct {
	APK1  string `json:"manifest1"`
	APK2  string `json:"manifest2"`
	Equal bool   `json:"equal"`
}

type APKComparisonArray struct {
	APK1Only []string `json:"manifest1Only"`
	APK2Only []string `json:"manifest2Only"`
	Common   []string `json:"common"`
	Equal    bool     `json:"equal"`
}

type APKComparisonFileArray struct {
	APK1Only []string `json:"manifest1Only"`
	APK2Only []string `json:"manifest2Only"`
	Common   []string `json:"common"`
	Changed  []string `json:"changed"`
	Equal    bool     `json:"equal"`
}

func (ma *APKComparisonFileArray) WriteText(writer *bytes.Buffer, onlyDiffs bool) {
	if ma != nil {
		if !onlyDiffs {
			for _, v := range ma.Common {
				writer.WriteString("* ")
				writer.WriteString(v)
				writer.WriteString("\n")
			}
		}
		for _, v := range ma.APK1Only {
			writer.WriteString(red(bold("- ")))
			writer.WriteString(red(bold(v)))
			writer.WriteString("\n")
		}
		for _, v := range ma.APK2Only {
			writer.WriteString(green(bold("+ ")))
			writer.WriteString(green(bold(v)))
			writer.WriteString("\n")
		}
		for _, v := range ma.Changed {
			writer.WriteString(yellow(bold("M ")))
			writer.WriteString(yellow(bold(v)))
			writer.WriteString("\n")
		}
		writer.WriteString("\n")
	}
}

func (ma *APKComparisonArray) WriteText(writer *bytes.Buffer, onlyDiffs bool) {
	if ma != nil {
		if !onlyDiffs {
			for _, v := range ma.Common {
				writer.WriteString("* ")
				writer.WriteString(v)
				writer.WriteString("\n")
			}
		}
		for _, v := range ma.APK1Only {
			writer.WriteString(red(bold("- ")))
			writer.WriteString(red(bold(v)))
			writer.WriteString("\n")
		}
		for _, v := range ma.APK2Only {
			writer.WriteString(green(bold("+ ")))
			writer.WriteString(green(bold(v)))
			writer.WriteString("\n")
		}
		writer.WriteString("\n")
	}
}

func (c *APKComparison) CompareFeatures() *APKComparisonArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.Manifest.GetFeatures(), c.APK2.Manifest.GetFeatures())

	arr := &APKComparisonArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}

	c.Features = arr
	return arr
}

func (c *APKComparison) ComparePermissions() *APKComparisonArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.Manifest.GetPermissions(), c.APK2.Manifest.GetPermissions())

	arr := &APKComparisonArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}
	c.Permissions = arr
	return arr
}

func (c *APKComparison) CompareActivities() *APKComparisonArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.Manifest.GetActivities(), c.APK2.Manifest.GetActivities())

	arr := &APKComparisonArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}
	c.Activities = arr
	return arr
}

func (c *APKComparison) CompareServices() *APKComparisonArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.Manifest.GetServices(), c.APK2.Manifest.GetServices())

	arr := &APKComparisonArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}
	c.Services = arr
	return arr
}

func (c *APKComparison) CompareReceivers() *APKComparisonArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.Manifest.GetReceivers(), c.APK2.Manifest.GetReceivers())

	arr := &APKComparisonArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}
	c.Receivers = arr
	return arr
}

func (c *APKComparison) CompareProviders() *APKComparisonArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.Manifest.GetProviders(), c.APK2.Manifest.GetProviders())

	arr := &APKComparisonArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}
	c.Providers = arr
	return arr
}

func (c *APKComparison) CompareSources() *APKComparisonFileArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.GetSources(), c.APK2.GetSources())

	arr := &APKComparisonFileArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}

	for _, f1 := range c.APK1.Sources {
		for _, f2 := range c.APK2.Sources {
			if f1.Path == f2.Path && f1.Filename == f2.Filename {
				if f1.Size != f2.Size || f1.Hash != f2.Hash {
					arr.Changed = append(arr.Changed, f1.GetFullPath())
				}
			}
		}
	}

	c.Sources = arr
	return arr
}

func (c *APKComparison) CompareFiles() *APKComparisonFileArray {
	aOnly, bOnly, both, equal := utils.CompareStringArrays(c.APK1.GetFiles(), c.APK2.GetFiles())

	arr := &APKComparisonFileArray{
		APK1Only: aOnly,
		APK2Only: bOnly,
		Common:   both,
		Equal:    equal,
	}

	for _, f1 := range c.APK1.Files {
		for _, f2 := range c.APK2.Files {
			if f1.Path == f2.Path && f1.Filename == f2.Filename {
				if f1.CompressedSize != f2.CompressedSize || f1.UncompressedSize != f2.UncompressedSize || f1.CRC32 != f2.CRC32 {
					arr.Changed = append(arr.Changed, f1.GetFullPath())
				}
			}
		}
	}

	c.Files = arr
	return arr
}

func (c *APKComparison) ComparePackageName() *APKComparisonString {
	s := &APKComparisonString{
		APK1:  c.APK1.Manifest.GetPackageName(),
		APK2:  c.APK2.Manifest.GetPackageName(),
		Equal: c.APK1.Manifest.GetPackageName() == c.APK2.Manifest.GetPackageName(),
	}
	c.PackageName = s
	return s
}

func (c *APKComparison) CompareAll() {
	c.ComparePackageName()
	c.CompareActivities()
	c.CompareFeatures()
	c.ComparePermissions()
	c.CompareProviders()
	c.CompareReceivers()
	c.CompareServices()
	c.CompareFiles()
	c.CompareSources()
}

func (c *APKComparison) GetComparison(format string, onlyDiffs bool, includeFiles bool) (string, error) {
	c.CompareAll()

	switch format {
	case "json":
		b, err := json.Marshal(c)
		return string(b), err
	case "json_pretty":
		b, err := json.MarshalIndent(c, "", "  ")
		return string(b), err
	case "text":
		var writer bytes.Buffer

		writer.WriteString(white(bold("Package 1: ")))
		writer.WriteString(c.PackageName.APK1)
		writer.WriteString("\n")
		writer.WriteString(white(bold("Package 2: ")))
		writer.WriteString(c.PackageName.APK2)
		writer.WriteString("\n")
		writer.WriteString("\n")

		writer.WriteString(white(bold("Features\n")))
		c.Features.WriteText(&writer, onlyDiffs)

		writer.WriteString(white(bold("Permissions\n")))
		c.Permissions.WriteText(&writer, onlyDiffs)

		writer.WriteString(white(bold("Services\n")))
		c.Services.WriteText(&writer, onlyDiffs)

		writer.WriteString(white(bold("Activities\n")))
		c.Activities.WriteText(&writer, onlyDiffs)

		writer.WriteString(white(bold("Receivers\n")))
		c.Receivers.WriteText(&writer, onlyDiffs)

		writer.WriteString(white(bold("Providers\n")))
		c.Providers.WriteText(&writer, onlyDiffs)

		if includeFiles {
			writer.WriteString(white(bold("Files\n")))
			c.Files.WriteText(&writer, onlyDiffs)

			writer.WriteString(white(bold("Sources\n")))
			c.Sources.WriteText(&writer, onlyDiffs)
		}

		return writer.String(), nil
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
