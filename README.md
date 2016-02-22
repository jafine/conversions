# Conversion formulas in Golang

[![Build Status](https://travis-ci.org/davidjpeacock/conversions.svg?branch=master)](https://travis-ci.org/davidjpeacock/conversions)
[![GoDoc](https://godoc.org/github.com/davidjpeacock/conversions?status.svg)](https://godoc.org/github.com/davidjpeacock/conversions)
[![License](https://img.shields.io/badge/license-GPL-blue.svg)](https://raw.githubusercontent.com/davidjpeacock/conversions/master/LICENSE)


This library is designed to provide safe conversions of units; each unit has a unique type T to prevent accidental unit mis-conversion.

## Contributing
Pull requests are very welcome; please follow the structure and ensure you provide tests for new conversions.

## Usage

```go
package main

import (
	"fmt"

	"github.com/davidjpeacock/conversions"
)

func main() {
	t := (30)

	psi := conversions.Psi(t)
	fmt.Printf("%s = %s\n", psi, conversions.PsiToBar(psi))
}
```

### Output

```
[peacock@trashcan psitobar]$ ./psitobar
30psi = 2.068428bar
```
