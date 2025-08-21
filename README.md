# 💤 LazyPM - WORK IN PROGRESS

**A lightweight, TUI-based project management tool that leverages Git for data storage and distribution.**

LazyPM helps developers manage tasks directly in the terminal, integrating project data (issues, time logs) with their Git repository.

---

## ✨ Core Features

*   **Terminal UI (TUI):** Manage projects and issues without leaving your command line.
*   **Git-Native Data:** Project data (`.lazypm/`) lives in your Git repo, allowing standard `git push`/`pull` for sync.
*   **Issue Management:** Create and list issues from the command line.
*   **Kanban Board:** Visual task board in the TUI.
*   **Time Tracking:** Log time directly against issues.
*   **GitHub Sync:** Option to sync issues with GitHub.com.

---

## 🚀 Get Started (Development)

Requires Go (1.21+).

1.  **Clone:**
    ```sh
    git clone https://github.com/Telikz/lazypm.git
    cd lazypm
    ```
2.  **Setup & Run:**
    ```sh
    go mod tidy
    go run ./cmd/lazypm
    ```
    
### Testing

```sh
go test ./tests
```

---

## 🏗 Project Structure Highlights

*   `cmd/lazypm`: Main application entry.
*   `internal/`: All core application logic (issues, TUI, commands, data handling).

---

## 🤝 Contributing

Contributions are welcome! Please ensure code is formatted.

---

## 📜 License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.