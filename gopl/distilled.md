# The Go programming language

## CHAPTER 1. TUTORIAL

- `index, element := range array`, e.g. `i, arg := range os.Args[1:]`
- A map is a reference to the data structure created by `make`, e.g. `make(map[string]int)`

### 1.5. Fetching a URL

- package `net/http` and `http.Get(url)`
- package `io/ioutil` and `ioutil.ReadAll`

## CHAPTER 3 BASIC TYPES

### 3.5 Strings

#### 3.5.4

- The conversions between strings and byte slices:

```
s := "abc"
b := []byte(s)
s2 := string(b)
```

## CHAPTER 4 COMPOSITE TYPES

- arrays: `numbers := [...]int{1, 2, 3}`
- array literal of index and value pairs:

```

type Currency int

const (
    USD = iota
    RMB
)

currencySymbols := [...]string{USD: "$", RMB: "Y"}

```

- `r := [...]int{99: 1}`: `r` has 100 elements and all of them are `0` except
  for the last one, which is `1`.

- Two arrays are equal if all of their corresponding elements are equal

- Arrays of the same element type and the same size can be assigned:

```go
a := [4]int{}
b := [4]int{3: 1}
fmt.Printf("%d\n", b) // [0 0 0 1]
b = a
fmt.Printf("%d\n", b) // [0 0 0 0]
```

- slices: A slice is a structure of three properties: a pointer to the underlying array, a length, and a capacity.

## CHAPTER 10 PACKAGES AND THE GO TOOL

### 10.1. Introduction

- Packages provide modularity, a name space, and encapsulation.

### 10.2. Import Paths

- Each package identified by a unique string called its _import path_.

- Use the factored (mathematically) import statement

```go
import (
        "fmt",
        "math/rand"
)
```

### 10.3. The Package Declaration

- A `package` declaration is required at the start of every Go source
  file, which defines a default identifier for the package, called the
  `package name` when it is imported.

- For example when `import "math/rand"`, the package name is `rand`
  and every file in the `rand` package starts with `package rand`.

- Conventionally, the package name is the last segment of the import
  path.

- There are three major exceptions to the "last segment" conventions
  - An executable program or a command always has the package name
    `main` so that `go build` knows it needs to invoke the linker to
    make an executable file.
  - `_test` (to come back after the test chapter)
  - Version number suffixes are excluded from the package name

### 10.4. Import Declarations

- Use blank lines to group imported packages by domains
- Use alternative names to avoid package name conflicts or give a
  context to enhance readability

```go
import (
        mathrand "math/rand" // mathrand is an alternative name to rand
)
```

### 10.5. Blank Imports

- Unused imports raise the "unused import" error.
- Use the blank identifier as the alternative name for the blank
  import: : `import _ "image/png"`.
- Blank imports are used to link code not explicitly referred in your
  programs, such as drivers for different databases

### 10.6. Packages and Naming

- The most frequently used packages in the standard library are
  `bufio`, `bytes`, `flag`, `fmt`, `http`, `io`, `json`, `os`, `sort`,
  `sync`, and `time`. TODO: be aware of their existence, and
  familiarise with their offerings gradually.

- Concise, descriptive, and unambiguous. Avoid choosing names that
  commonly used for related local variables.

- Use the singular form. Be aware that the standard packages `bytes`,
  `errors`, and `strings` are chosen to avoid hiding predeclared
  types, and `go/types` to avoid conflict with a keyword.

- Package members use camel case.

- Define a `New` function to create instances for `single-type packages`, which expose one principal data type.

### 10.7. The Go Tool

#### 10.7.1. Workspace Organisation

- `GOPATH` specifies the root of the workspace. Change `GOPATH` to
  switch between workspace.
- `GOPATH` has three subdirections: `src`, `bin`, and `pkg`.
- `GOROOT` specifies the root directory of the Go distribution.
- `go env` to print the effective values of the current environment.

#### 10.7.5. Internal Packages
