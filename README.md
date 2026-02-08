# 🏀 nbarecap

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)]()

A terminal-based NBA companion for checking scores, box scores, and play-by-play action.
Keep your ball knowledge high without leaving your dev environment!

##  Features

- 📅 **Games List**: Browse games for any specific date.
- 📊 **Box Scores**: View detailed stats and percentages for each player.
- 🔄 **Play-by-Play**: Follow the sequence of events for every period.
- 🎨 **TUI Interface**: Interactive terminal interface built with Bubble Tea.
- 🔌 **Custom NBA Stats API Client**: Includes a custom-authored Go client for the NBA Stats API (pkg/nba_api).

## 🚀 Installation

```bash
go install github.com/bobbyskywalker/nbarecap@latest
```

## 🛠 Usage

Start by viewing today's games:

```bash
nbarecap games
```

To view games for a specific date:

```bash
nbarecap games --date 2026-02-02
```

## 📦 Internal API Client

The project features a custom implementation of the NBA Stats API located in `pkg/nba_api`. It handles requests and data mapping for:
- Scoreboard V2
- Box Score Traditional V3
- Play-by-Play V3

Feel free to use it in your own projects!

## 🗂 Dependencies
The Holy Trinity:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [Cobra](https://github.com/spf13/cobra)

## 🗺 Roadmap

Future plans for `nbarecap`:
- 🚀 **Official Release**
- 🏆 **Leaders**: New `leaders` command.
- 📈 **Standings**: New `standings` command.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
