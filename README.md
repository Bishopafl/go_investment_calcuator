# 📈 Investment Calculator (Go)

This simple Go program helps estimate the **future value** of an investment, factoring in both **expected return rate** and **inflation**. It uses basic console input/output and the Go `math` package to compute compound growth over a given number of years.

---

## 🧠 Features

* Calculates **future value** of an investment with compound interest.
* Adjusts for **inflation** to estimate the **real future value**.
* Prompts users interactively for:

  * Investment amount
  * Expected annual return rate
  * Years of investment
* Demonstrates Go fundamentals:

  * Constant and variable declarations
  * User input via `fmt.Scan()`
  * Mathematical computations with `math.Pow()`

---

## 💻 Example Usage

### Run the program

```bash
go run main.go
```

### Example Input/Output

```
Investment Amount: 10000
Expected Return Rate: 5.5
Years of Investment: 10
```

**Output:**

```
17078.928
13391.444
```

Where:

* `17078.928` → Future Value (without inflation)
* `13391.444` → Future Real Value (adjusted for inflation)

---

## ⚙️ Formula Breakdown

### Future Value (FV)

[
FV = P \times (1 + \frac{r}{100})^{t}
]

### Future Real Value (Adjusted for Inflation)

[
FV_{real} = \frac{FV}{(1 + \frac{i}{100})^{t}}
]

Where:

* `P` = Investment amount
* `r` = Expected return rate
* `t` = Years of investment
* `i` = Inflation rate

---

## 🧩 Code Overview

```go
const inflationRate = 2.5
var investmentAmount, years float64
expectedReturnRate := 5.5

fmt.Scan(&investmentAmount)
fmt.Scan(&expectedReturnRate)
fmt.Scan(&years)

futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
```

---

## 🛠️ Requirements

* Go 1.18+
* Basic terminal access

---

## 🚀 Run and Build

To build an executable:

```bash
go build -o investment-calculator
./investment-calculator
```

---

## 📜 License

This project is released under the [MIT License](LICENSE).
