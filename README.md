# 🕹️ Go Mini Game – Player Inventory System

A simple Go project that demonstrates **structs**, **methods**, and **slice manipulation** through a mini inventory (backpack) system for a player.

---

## 🎯 Features

- 🧺 **Pick Up Items** – Add items dynamically to a player's backpack using Go slices (`append`).
- ⚔️ **Use Items** – Remove items by name, simulating consumption or usage.
- 📦 **Print Backpack** – Display current inventory contents with formatted output.
- 🧩 **Graceful Handling** – Prevents runtime errors when trying to use non-existent items.

---

## 🧠 Concepts Demonstrated

- ✅ Struct composition (`Player` and `Item`)
- ✅ Pointer receivers to modify struct fields
- ✅ Slice operations (`append`, slicing, and removal)
- ✅ Looping and conditional logic
- ✅ Encapsulation via Go methods

---

## 🧩 Example Output

```bash
Steve picked: Apple
Steve's backpack contains:
- Apple
Steve used: Apple
Steve's backpack is empty.
Steve tried to use but Apple doesn't have it.
Steve tried to use but Orange doesn't have it.
```

## 🚀 How to Run
```bash
# Clone the repository
git clone https://github.com/<your-username>/go-mini-game.git
cd go-mini-game

# Run the program
go run main.go
```