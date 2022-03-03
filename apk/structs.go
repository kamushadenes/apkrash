package apk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kamushadenes/apkrash/utils"
	"strconv"
	"strings"
)

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var white = color.New(color.FgWhite).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()

type BundleToolSpec struct {
	SupportedABIs    []string `json:"supportedAbis"`
	SupportedLocales []string `json:"supportedLocales"`
	ScreenDensity    int      `json:"screenDensity"`
	SdkVersion       int      `json:"sdkVersion"`
}

type APKComparison struct {
	APK1                      *APK                    `json:"-"`
	APK2                      *APK                    `json:"-"`
	CompileSDKVersion         *APKComparisonInt       `json:"compileSDKVersion"`
	CompileSDKVersionCodename *APKComparisonString    `json:"compileSDKVersionCodename"`
	PlatformBuildVersionCode  *APKComparisonInt       `json:"platformBuildVersionCode"`
	PlatformBuildVersionName  *APKComparisonString    `json:"platformBuildVersionName"`
	PackageName               *APKComparisonString    `json:"packageName"`
	Features                  *APKComparisonArray     `json:"features"`
	Permissions               *APKComparisonArray     `json:"permissions"`
	Services                  *APKComparisonArray     `json:"services"`
	Receivers                 *APKComparisonArray     `json:"receivers"`
	Activities                *APKComparisonArray     `json:"activities"`
	Providers                 *APKComparisonArray     `json:"providers"`
	Files                     *APKComparisonFileArray `json:"files"`
	Sources                   *APKComparisonFileArray `json:"sources"`
	MainActivity              *APKComparisonString    `json:"mainActivity"`
}

type APKComparisonInt struct {
	APK1  int64 `json:"apk1"`
	APK2  int64 `json:"apk2"`
	Equal bool  `json:"equal"`
}

type APKComparisonString struct {
	APK1  string `json:"apk1"`
	APK2  string `json:"apk2"`
	Equal bool   `json:"equal"`
}

type APKComparisonArray struct {
	APK1Only []string `json:"apk1Only"`
	APK2Only []string `json:"apk2Only"`
	Common   []string `json:"common"`
	Equal    bool     `json:"equal"`
}

type APKComparisonFileArray struct {
	APK1Only []string `json:"apk1Only"`
	APK2Only []string `json:"apk2Only"`
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

func (c *APKComparison) CompareMainActivity() *APKComparisonString {
	a1 := c.APK1.Manifest.GetMainActivity()
	a2 := c.APK2.Manifest.GetMainActivity()

	cmp := &APKComparisonString{
		APK1:  a1,
		APK2:  a2,
		Equal: a1 == a2,
	}

	c.MainActivity = cmp

	return cmp
}

func (c *APKComparison) CompareCompileSDKVersion() *APKComparisonInt {
	v1, _ := strconv.ParseInt(c.APK1.Manifest.CompileSdkVersion, 10, 64)
	v2, _ := strconv.ParseInt(c.APK2.Manifest.CompileSdkVersion, 10, 64)
	cmp := &APKComparisonInt{
		APK1:  v1,
		APK2:  v2,
		Equal: v1 == v2,
	}

	c.CompileSDKVersion = cmp

	return cmp
}

func (c *APKComparison) CompareCompileSDKVersionCodename() *APKComparisonString {
	cmp := &APKComparisonString{
		APK1:  c.APK1.Manifest.CompileSdkVersionCodename,
		APK2:  c.APK2.Manifest.CompileSdkVersionCodename,
		Equal: c.APK1.Manifest.CompileSdkVersionCodename == c.APK2.Manifest.CompileSdkVersionCodename,
	}

	c.CompileSDKVersionCodename = cmp

	return cmp
}

func (c *APKComparison) ComparePlatformBuildVersionCode() *APKComparisonInt {
	v1, _ := strconv.ParseInt(c.APK1.Manifest.PlatformBuildVersionCode, 10, 64)
	v2, _ := strconv.ParseInt(c.APK2.Manifest.PlatformBuildVersionCode, 10, 64)
	cmp := &APKComparisonInt{
		APK1:  v1,
		APK2:  v2,
		Equal: v1 == v2,
	}

	c.PlatformBuildVersionCode = cmp

	return cmp
}

func (c *APKComparison) ComparePlatformBuildVersionName() *APKComparisonString {
	cmp := &APKComparisonString{
		APK1:  c.APK1.Manifest.PlatformBuildVersionName,
		APK2:  c.APK2.Manifest.PlatformBuildVersionName,
		Equal: c.APK1.Manifest.PlatformBuildVersionName == c.APK2.Manifest.PlatformBuildVersionName,
	}

	c.PlatformBuildVersionName = cmp

	return cmp
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
	c.CompareCompileSDKVersion()
	c.CompareCompileSDKVersionCodename()
	c.ComparePlatformBuildVersionCode()
	c.ComparePlatformBuildVersionName()
	c.CompareMainActivity()
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
	case "table":
		t := table.NewWriter()
		t.AppendHeader(table.Row{"Field", "APK 1", "APK 2", "Equal?"})
		t.AppendRows([]table.Row{{"Package Name", c.PackageName.APK1, c.PackageName.APK2, c.PackageName.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Platform Build Version", c.PlatformBuildVersionCode.APK1, c.PlatformBuildVersionCode.APK2, c.PlatformBuildVersionCode.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Main Activity", c.MainActivity.APK1, c.MainActivity.APK2, c.MainActivity.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Features", strings.Join(append(c.Features.APK1Only, c.Features.Common...), "\n"), strings.Join(append(c.Features.APK2Only, c.Features.Common...), "\n"), c.Features.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Permissions", strings.Join(append(c.Permissions.APK1Only, c.Permissions.Common...), "\n"), strings.Join(append(c.Permissions.APK2Only, c.Permissions.Common...), "\n"), c.Permissions.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Services", strings.Join(append(c.Services.APK1Only, c.Services.Common...), "\n"), strings.Join(append(c.Services.APK2Only, c.Services.Common...), "\n"), c.Services.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Providers", strings.Join(append(c.Providers.APK1Only, c.Providers.Common...), "\n"), strings.Join(append(c.Providers.APK2Only, c.Providers.Common...), "\n"), c.Providers.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Activities", strings.Join(append(c.Activities.APK1Only, c.Activities.Common...), "\n"), strings.Join(append(c.Activities.APK2Only, c.Activities.Common...), "\n"), c.Activities.Equal}})
		t.AppendSeparator()
		t.AppendRows([]table.Row{{"Receivers", strings.Join(append(c.Receivers.APK1Only, c.Receivers.Common...), "\n"), strings.Join(append(c.Receivers.APK2Only, c.Receivers.Common...), "\n"), c.Receivers.Equal}})
		t.SetStyle(table.StyleLight)

		return t.Render(), nil
	case "text":
		var writer bytes.Buffer

		writer.WriteString(white(bold("Package 1: ")))
		writer.WriteString(c.PackageName.APK1)
		writer.WriteString("\n")
		writer.WriteString(white(bold("Compile SDK Version: ")))
		writer.WriteString(fmt.Sprintf("%d", c.CompileSDKVersion.APK1))
		writer.WriteString(fmt.Sprintf(" (%s)", c.CompileSDKVersionCodename.APK1))
		writer.WriteString("\n")
		writer.WriteString(white(bold("Platform Build Version: ")))
		writer.WriteString(fmt.Sprintf("%d", c.PlatformBuildVersionCode.APK1))
		writer.WriteString(fmt.Sprintf(" (%s)", c.PlatformBuildVersionName.APK1))
		writer.WriteString("\n")
		writer.WriteString(white(bold("Main Activity: ")))
		writer.WriteString(c.MainActivity.APK1)
		writer.WriteString("\n")
		writer.WriteString("\n")

		writer.WriteString(white(bold("Package 2: ")))
		writer.WriteString(c.PackageName.APK2)
		writer.WriteString("\n")
		writer.WriteString(white(bold("Compile SDK Version: ")))
		writer.WriteString(fmt.Sprintf("%d", c.CompileSDKVersion.APK2))
		writer.WriteString(fmt.Sprintf(" (%s)", c.CompileSDKVersionCodename.APK2))
		writer.WriteString("\n")
		writer.WriteString(white(bold("Platform Build Version: ")))
		writer.WriteString(fmt.Sprintf("%d", c.PlatformBuildVersionCode.APK2))
		writer.WriteString(fmt.Sprintf(" (%s)", c.PlatformBuildVersionName.APK2))
		writer.WriteString("\n")
		writer.WriteString(white(bold("Main Activity: ")))
		writer.WriteString(c.MainActivity.APK2)
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
