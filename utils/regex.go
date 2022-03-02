package utils

import "regexp"

var PackageRegex = regexp.MustCompile("^package (.*?);")
var ImportRegex = regexp.MustCompile("import (.*?);")
