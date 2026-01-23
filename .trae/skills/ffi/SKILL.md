---
name: ffi
description: Summarizes FFI (Foreign Function Interface) usage tips for Go projects, applicable to any C dynamic library integration. Invoke when using purego or ffi library.
---

# FFI Usage Tips for Go Projects

This skill provides a comprehensive summary of FFI (Foreign Function Interface) usage tips for Go projects, applicable to any scenario where Go code needs to interact with C dynamic libraries.

## Overview

These tips demonstrate how to use both `github.com/jupiterrider/ffi` and `github.com/ebitengine/purego` packages to directly integrate with C dynamic libraries from Go applications, enabling cross-platform support across Linux, macOS, and Windows without requiring CGo.

| Package | Primary Uses |
|---------|--------------|
| **ffi** | Loading dynamic libraries, calling C functions, handling type mappings between Go and C, working with C structs and pointers, managing dynamic library handles |
| **purego** | Creating C-compatible callback functions, initializing structs that will be passed by pointer to C functions |

## Integration Pattern

The typical integration pattern combining both packages is:

1. **Load Libraries**: Use FFI to load dynamic libraries
2. **Prepare Functions**: Use FFI to prepare C functions for calling
3. **Create Callbacks**: Use PureGo to create C-compatible callbacks
4. **Register Callbacks**: Pass callbacks to C functions using FFI
5. **Call C Functions**: Use FFI to call C functions as needed
6. **Handle Callbacks**: C library calls back to Go functions via PureGo callbacks

```go
// Example Integration Flow

import (
    "fmt"
    "github.com/jupiterrider/ffi"
    "github.com/ebitengine/purego"
    "unsafe"
)

func main() {
    // 1. Load library
    lib, err := ffi.Load("libexample.so")
    if err != nil {
        panic(err)
    }

    // 2. Prepare C functions
    setCallbackFunc, err := lib.Prep("set_callback", &ffi.TypeVoid, &ffi.TypePointer)
    if err != nil {
        panic(err)
    }

    // 3. Create callback using PureGo
    goCallback := purego.NewCallback(func(data uintptr) uintptr {
        fmt.Printf("Callback called with data: %d\n", data)
        return 0
    })

    // 4. Register callback with C library
    setCallbackFunc.Call(nil, unsafe.Pointer(&goCallback))

    // 5. Call other C functions as needed
    // ...
}
```

## Core Concepts

### 1. Import Packages

```go
import (
    "github.com/jupiterrider/ffi"
    "github.com/ebitengine/purego"
    "unsafe"
)
```

### 2. Load Dynamic Library Functions

Use `lib.Prep()` to load C functions from dynamic libraries:

```go
lib, err := ffi.Load("libexample.so")
if err != nil {
    panic(err)
}

// Load different types of C functions
initFunc, err := lib.Prep("initialize", &ffi.TypeVoid)                     // No params, no return
addFunc, err := lib.Prep("add", &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32) // With params and return value
```

### 3. Function Call Patterns

```go
// No parameters, no return value
initFunc.Call(nil)

// With return value
var count uint64
getCountFunc.Call(unsafe.Pointer(&count))

// With parameters and return value
var result int32
a, b := int32(10), int32(20)
addFunc.Call(unsafe.Pointer(&result), unsafe.Pointer(&a), unsafe.Pointer(&b))
```

### 4. Type Mapping

Map C types to Go types:

| C Type | Go Type | FFI Type |
|--------|---------|----------|
| `void` | `nil` | `ffi.TypeVoid` |
| `int` | `int32` | `ffi.TypeSint32` |
| `uint` | `uint32` | `ffi.TypeUint32` |
| `int64_t` | `int64` | `ffi.TypeSint64` |
| `size_t` | `uint64` | `ffi.TypeUint64` |
| `bool` | `bool` | `ffi.TypeUint8` |
| `char*` | `*byte` | `ffi.TypePointer` |
| `void*` | `uintptr` | `ffi.TypePointer` |
| `float` | `float32` | `ffi.TypeFloat` |
| `double` | `float64` | `ffi.TypeDouble` |

### 5. Struct Handling

#### Define Matching Go Structs

Create Go structs that match the layout of C structs and define FFI types for them:

```go
// C struct: struct example_struct { int32_t id; const char* name; float value; }

type ExampleStruct struct {
    ID    int32   // id field
    Name  *byte   // name field
    Value float32 // value field
}

// Create FFI type for the struct
var FFITypeExampleStruct = ffi.NewType(
    &ffi.TypeSint32,  // id
    &ffi.TypePointer, // name
    &ffi.TypeFloat    // value
)
```

#### Pass Structs to C Functions

```go
// Create and initialize a struct
namePtr, _ := GoStringToCString("test")
myStruct := ExampleStruct{
    ID:    1,
    Name:  namePtr,
    Value: 3.14,
}

// Pass struct to C function
processStructFunc.Call(nil, unsafe.Pointer(&myStruct))
```

#### Struct Initialization with PureGo

When C functions expect a pointer to a struct for initialization purposes, use PureGo to properly register and initialize these structs:

```go
// In the library implementation
func (c *CStableDiffusionImpl) CtxParamsInit() CtxParams {
    // Allocate struct on heap to prevent memory drift
    structPtr := new(CtxParams)
    
    // Use PureGo to call the C initialization function
    c.ctxParamsInit(uintptr(unsafe.Pointer(structPtr)))
    
    return *structPtr
}

// Prepare the initialization function using PureGo
purego.RegisterLibFunc(&impl.ctxParamsInit, libSd.Addr, "sd_ctx_params_init")
```

### 6. Callback Functions

#### Create C Callbacks with PureGo

```go
// Example: Simple boolean callback
type AbortFunc func() bool

func newAbortCallback(fn AbortFunc) uintptr {
    return purego.NewCallback(func(data uintptr) uintptr {
        if fn() {
            return 1 // true in C
        }
        return 0 // false in C
    })
}
```

#### Callback Signature Mapping

Map Go function signatures to C callback signatures correctly:

| C Callback Type | Go Function Signature |
|----------------|----------------------|
| `bool (*)(void*)` | `func(data uintptr) uintptr` |
| `void (*)(int32_t, char*, void*)` | `func(level int32, text, data uintptr) uintptr` |
| `int (*)(int, int)` | `func(a, b uintptr) uintptr` |
| `void (*)(void*)` | `func(data uintptr) uintptr` |
| `char* (*)(const char*)` | `func(str uintptr) uintptr` |

#### Pass Callbacks to C Functions

```go
// Example: Passing a callback to a C function
ctx := uintptr(0) // Your context pointer
callback := newAbortCallback(func() bool {
    return shouldAbort // Your abort logic
})
nilPtr := uintptr(0) // Proper null pointer for void*

setAbortCallbackFunc.Call(nil, unsafe.Pointer(&ctx), unsafe.Pointer(&callback), unsafe.Pointer(&nilPtr))
```

#### Callback Best Practices

- **Always initialize all parameters**: Create valid variables for all parameters, even if they are null pointers
- **Use uintptr(0) for void* null pointers**: Never use `nil` directly for `void*` parameters
- **Pass pointers correctly**: Use `unsafe.Pointer(&variable)` for parameters
- **Follow exact C signature**: Ensure the number and type of parameters match the C function exactly
- **Handle void* properly**: When dealing with `void*` user data, always pass a valid memory address
- **Avoid stack-allocated variables**: Don't use pointers to stack-allocated variables that might go out of scope
- **Thread safety**: Ensure callbacks are thread-safe if called from multiple threads
- **No Go pointers**: Don't store Go pointers in C memory unless absolutely necessary

### 7. void* Type Handling

#### Passing Null as void*

```go
// Correct: Using uintptr(0) for void* null pointer
userData := uintptr(0)
cFunction.Call(nil, unsafe.Pointer(&userData))

// Incorrect: Using nil directly
cFunction.Call(nil, nil) // WRONG: nil is not a valid void* null pointer
```

#### Passing Go Data to void*

```go
// Create a structure to hold your data
type MyUserData struct {
    Value int
    Name  string
}

// Allocate data on heap to avoid stack issues
userData := &MyUserData{Value: 42, Name: "test"}
userDataPtr := uintptr(unsafe.Pointer(userData))

// Pass to C function
cFunction.Call(nil, unsafe.Pointer(&userDataPtr))
```

#### Receiving void* Data in Callbacks

```go
cCallback := purego.NewCallback(func(arg1 int32, data uintptr) uintptr {
    // Convert void* back to Go pointer
    userData := (*MyUserData)(unsafe.Pointer(data))
    
    // Access the data
    fmt.Printf("Received data: Value=%d, Name=%s\n", userData.Value, userData.Name)
    
    return 0
})
```

### 8. String Handling

```go
// Convert Go string to C string (null-terminated)
func GoStringToCString(s string) (*byte, error) {
    switch runtime.GOOS {
    case "windows":
        return windows.BytePtrFromString(s)
    default:
        return unix.BytePtrFromString(s)
    }
}

// Convert C string to Go string
func CStringToGoString(p *byte) string {
    switch runtime.GOOS {
    case "windows":
        return windows.BytePtrToString(p)
    default:
        return unix.BytePtrToString(p)
    }
}
```

## Advanced Concepts

### 1. Memory Management

#### Heap Allocation to Prevent Memory Drift

When working with FFI, allocate structs on the heap instead of the stack to prevent memory drift:

```go
// Safe: Heap allocation
func (c *LibraryImpl) GoodStructInit() MyStruct {
    structOnHeap := new(MyStruct) // Heap-allocated struct
    c.initStruct.Call(nil, unsafe.Pointer(structOnHeap))
    return *structOnHeap
}
```

#### Function Signature Design

Design user-friendly function signatures that return structs directly:

```go
// User-friendly: Library manages memory
func (c *LibraryImpl) InitStruct() StructParams {
    structPtr := new(StructParams)
    c.initStruct.Call(nil, unsafe.Pointer(structPtr))
    return *structPtr
}
```

### 2. Struct Offset Validation

Use Go's reflection package to verify that the memory layout of Go structs matches the expected C struct layout:

```go
// CheckStructOffsets validates the offsets of a struct using reflection
func CheckStructOffsets(structType ffi.Type, structValue interface{}) StructOffsetResult {
    // Implementation details...
}
```

### 3. Cross-Platform Library Loading

```go
func LoadLibrary(name string) (ffi.Lib, error) {
    libName := GetLibraryFilename(name)
    return ffi.Load(libName)
}

func GetLibraryFilename(libName string) string {
    switch runtime.GOOS {
    case "linux", "freebsd":
        return fmt.Sprintf("lib%s.so", libName)
    case "windows":
        return fmt.Sprintf("%s.dll", libName)
    case "darwin":
        return fmt.Sprintf("lib%s.dylib", libName)
    default:
        return libName
    }
}
```

## Best Practices

### 1. Code Organization

Organize your FFI code using a simple, library-based structure:

```
// Basic Structure
.
├── {library-name}/
│   ├── {library-name}.go  // Library initialization and loading
│   ├── types.go             // Types and constants
│   └── functions.go         // All FFI functions and wrappers
```

### 2. Naming Conventions

- Use the pattern `{functionName}Func` for FFI function pointers
- Use descriptive names for types and structs
- Use `FFIType{StructName}` for FFI type definitions
- Use PascalCase for wrapper functions with clear names

### 3. Error Handling

```go
func loadFunction(lib ffi.Lib, name string, retType *ffi.Type, argTypes ...*ffi.Type) (ffi.Fun, error) {
    fun, err := lib.Prep(name, retType, argTypes...)
    if err != nil {
        return ffi.Fun{}, fmt.Errorf("failed to load %s: %w", name, err)
    }
    return fun, nil
}
```

### 4. Resource Management

- **Explicit Resource Release**: Provide clear functions to release resources
- **Avoid Memory Leaks**: Ensure all allocated resources are freed
- **Document Ownership**: Clarify who owns the memory (Go or C)

```go
// Example: Proper memory management for models
func ModelLoad(path string) (Model, error) {
    // Load model implementation...
}

func ModelFree(model Model) error {
    // Free model implementation...
}
```

## Common Pitfalls and Solutions

| Problem | Root Cause | Solution |
|---------|------------|----------|
| Memory access errors when passing callbacks | Passing invalid pointer types, not initializing all parameters, using `nil` directly instead of `uintptr(0)` for void* | Always initialize all parameters, use `uintptr(0)` for void* null pointers, follow exact C signature |
| Memory drift with structs | Stack-allocated structs being overwritten after function returns | Allocate structs on the heap using `new()` or `&struct{}` |
| Callback signature mismatch | Incorrect mapping between Go and C callback signatures | Follow the callback signature mapping guidelines |
| Cross-platform library loading issues | Using platform-specific library filenames | Implement cross-platform library loading function |

## Conclusion

By following these tips and best practices, you can create robust, maintainable, and safe Go bindings for any C dynamic library. These techniques are universally applicable to any scenario where you need to integrate Go with C dynamic libraries, enabling you to leverage the extensive ecosystem of C libraries in your Go projects.

Key takeaways:
- Use the right package for each task: ffi for library operations, purego for callbacks and struct initialization
- Follow consistent naming conventions and code organization
- Implement proper error handling and resource management
- Ensure cross-platform compatibility
- Use heap allocation for structs passed to C functions
- Follow exact C function signatures for callbacks