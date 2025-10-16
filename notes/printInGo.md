# Go Print Functions (fmt package)

## 1. fmt.Print()
- Prints text or variables without a newline.
- Does not add spaces automatically.

### Example:
```go
fmt.Print("Hello")
fmt.Print("World")
```
Output:
```
HelloWorld
```

---

## 2. fmt.Println()
- Prints text or variables with a newline at the end.
- Adds spaces automatically between values.

### Example:
```go
fmt.Println("Hello")
fmt.Println("World")
```
Output:
```
Hello
World
```

### Another example:
```go
fmt.Println("Hello", "World")
```
Output:
```
Hello World
```

---

## 3. fmt.Printf()
- Prints text using format specifiers.
- You control spaces, newlines, and formatting.

### Example:
```go
name := "Virta"
age := 26
fmt.Printf("My name is %s and I am %d years old.\n", name, age)
```
Output:
```
My name is Virta and I am 26 years old.
```

---

## Format Specifiers (used with Printf)

| Specifier | Meaning | Example | Output |
|------------|----------|----------|---------|
| %d | Integer | fmt.Printf("%d", 10) | 10 |
| %f | Float | fmt.Printf("%.2f", 3.14159) | 3.14 |
| %s | String | fmt.Printf("%s", "Go") | Go |
| %t | Boolean | fmt.Printf("%t", true) | true |
| %v | Default format | fmt.Printf("%v", 123) | 123 |
| %T | Type of variable | fmt.Printf("%T", name) | string |

---

## Example combining all
```go
name := "Virta"
age := 26
height := 5.6

fmt.Print("Name: ", name, " ")
fmt.Println("Age:", age)
fmt.Printf("Height: %.1f ft\n", height)
```
Output:
```
Name: Virta Age: 26
Height: 5.6 ft
```

---

## Summary

| Function | Adds Space | Adds Newline | Supports Formatting | Best For |
|-----------|-------------|---------------|----------------------|-----------|
| Print | No | No | No | Quick inline printing |
| Println | Yes | Yes | No | Simple readable output |
| Printf | You control | You control | Yes | Formatted output |
