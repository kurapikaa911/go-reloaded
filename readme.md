

---

# ✨ **Go Reloaded: Text Transformer** 🚀

Welcome to **Go Reloaded**, where text correction meets automation! Transform your words with ease and let our tool handle the tedious work for you!

---

## 🎯 **Mission Brief**

The **Text Transformer** ingests a text file, scans for specific markers, and seamlessly modifies words, numbers, and phrases. Whether it’s converting hexadecimal values, adjusting text case, or cleaning up messy sentences, we've got you covered. The result? A polished `result.txt` ready to impress!

---

## 🌟 **Features**

🔍 **What the Text Transformer Does:**

### 🧮 Hexadecimal to Decimal Conversion
Got a hex number? No worries!  
`(hex)` converts it like magic!  
- `"1A (hex)"` → `"26"`

### 💻 Binary to Decimal
Have binary numbers in your text?  
`(bin)` converts them to decimal in a flash!  
- `"1101 (bin)"` → `"13"`

### 🎨 Case Transformation
Want to change the text case? We support everything!  
- `(up)` for UPPERCASE  
  - `"word (up)"` → `"WORD"`
- `(low)` for lowercase  
  - `"WORD (low)"` → `"word"`
- `(cap)` for Capitalized  
  - `"bridge (cap)"` → `"Bridge"`
  
Need to convert multiple words? Use numbers with `(up, X)`, `(low, X)`, or `(cap, X)` for greater control!  
- `"this is fun (up, 2)"` → `"THIS IS fun"`

### ✒️ Perfect Punctuation
Keep your commas, periods, and exclamation marks in check!  
- `"Hello , world !"` becomes `"Hello, world!"`

### 📝 Smart Article Correction
Don’t worry about tricky articles—your "A" will automatically change to "An" where needed.  
- `"A apple"` → `"An apple"`

### 📖 Quote Handling
This project features two quote handling methods: **basic** and **advanced**.  
To use the advanced method, simply uncomment the second quote function in the code!

---

## ⚙️ **How It Works**

Our tool processes a text file by taking **two filenames** as input:

1. **firstFilename**: The text file to transform  
2. **secondFilename**: A placeholder file for validation (don’t worry, it’s not processed)  

Once fed with these inputs, the text undergoes a series of validation checks using `tools.Checktype`, ensuring that the files are of the correct type (only `.txt` please!). The content is then transformed using `Text_cleaner` and `Text_corrector`.

### 💾 **Output**
Your transformed text is saved in `result.txt`—ready for whatever comes next!

---

## ❗ **Error Handling**

- **Wrong File Types**: If the files aren’t of the correct format, an error message will guide you.  
- **File Read/Write Issues**: Any hiccups during reading or writing are handled gracefully, with clear error messages.

---

## 🛠️ **How to Run It**

Want to see it in action? Here’s how:  

```bash
cd cmd
go run main.go firstFilename.txt secondFilename.txt
```

Sit back and watch the magic happen! 🪄

---

## 🔍 **Easier Testing for Auditors**

For quick and easy testing, especially for auditors, we’ve added sample files: `sample1.txt`, `sample2.txt`, and `sample3.txt`. These files contain various test cases to demonstrate how the tool works. Just use them as inputs to see the transformations in action!

Make sure to write unit tests for every feature—it’s a lifesaver!

---

## 🖥️ **Tech Stack**

🦄 **Languages**: Go (Golang), Markdown  
⚙️ **Tools**: VSCode, Go testing framework, Git  
📦 **Packages**: `strconv`, `strings`

---
![Go Logo]([https://golang.org/lib/godoc/images/go-logo-blue.svg](https://www.google.com/url?sa=i&url=https%3A%2F%2Fopenupthecloud.com%2Fcloud-engineers-code%2Fgolang%2F&psig=AOvVaw0NS5GCrMyDHOmsTICXUQ-x&ust=1726964944983000&source=images&cd=vfe&opi=89978449&ved=0CBQQjRxqFwoTCNjE2PLj0ogDFQAAAAAdAAAAABAE))


