# Log Aggregator TUI

**Week 3 Challenge** - A terminal-based log file viewer with real-time monitoring, filtering, and search capabilities.

## What You're Building

A powerful TUI application that helps developers monitor and analyze log files in real-time. Think of it as `tail -f` meets `grep` with a beautiful interactive interface.

## Core Features

### Must-Have (MVP)
- **Log Display**: Read and display log files with syntax highlighting
- **Real-time Monitoring**: Watch files for changes and auto-update
- **Log Level Filtering**: Toggle visibility of INFO, WARN, ERROR, DEBUG levels
- **Text Search**: Filter logs by content with regex support
- **Scrolling**: Navigate through large log files smoothly

### Nice-to-Have (Stretch Goals)
- Multiple file support (tabs or split view)
- Export filtered results
- Timestamp range filtering
- Statistics dashboard (error count, trends)
- Custom log format parsing

## Why This Project Matters

**For Cloud Games Engineering:**
- Essential skill for debugging production game servers
- Learn observability patterns used in distributed systems
- Understand how monitoring tools work under the hood
- Practice handling high-throughput data streams

**Technical Growth:**
- File system operations and watching
- Regular expressions and text processing
- Concurrent programming (UI + file watching)
- Advanced Bubbletea patterns (complex state management)

## Learning Goals

By completing this project, you'll master:
- File I/O patterns (reading, seeking, tailing)
- Using `fsnotify` or similar for file watching
- Regex for log parsing and filtering
- Managing complex application state in Bubbletea
- Handling real-time updates without blocking UI

## Suggested Architecture

```
log-aggregator/
├── cmd/
│   └── main.go                 # Entry point
├── internal/
│   ├── log/
│   │   ├── parser.go           # Parse log lines into structured data
│   │   ├── watcher.go          # File watching with fsnotify
│   │   └── filter.go           # Filtering and search logic
│   └── ui/
│       ├── model.go            # Bubbletea model (application state)
│       ├── view.go             # Rendering logic
│       ├── update.go           # Event handling and updates
│       └── styles.go           # Colors and formatting
├── testdata/
│   └── sample.log              # Test log files
├── go.mod
└── README.md
```

## Development Phases

### Phase 1: Basic Display (Days 1-2)
- [ ] Set up Bubbletea application structure
- [ ] Read log file and display content
- [ ] Implement basic scrolling
- [ ] Parse log lines (timestamp, level, message)

### Phase 2: Filtering & Search (Days 3-4)
- [ ] Add log level filter toggles
- [ ] Implement text search functionality
- [ ] Highlight matching search terms
- [ ] Show filtered line count

### Phase 3: Real-time Updates (Days 5-6)
- [ ] Integrate `fsnotify` for file watching
- [ ] Update display when file changes
- [ ] Handle file rotation scenarios
- [ ] Test with rapidly changing logs

### Phase 4: Polish & Testing (Day 7)
- [ ] Add help screen with keyboard shortcuts
- [ ] Improve UI with colors and indicators
- [ ] Add statistics (total lines, errors, etc.)
- [ ] Test with real application logs
- [ ] Write documentation

## Key Packages

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling and layout
- `github.com/fsnotify/fsnotify` - File system notifications
- Standard library: `bufio`, `regexp`, `os`

## Log Format Example

Your parser should handle common formats like:
```
2025-01-15 14:23:45 INFO  Server started on port 8080
2025-01-15 14:23:46 DEBUG Connected to database
2025-01-15 14:23:50 WARN  High memory usage detected: 85%
2025-01-15 14:24:12 ERROR Failed to process request: connection timeout
```

## Keyboard Shortcuts (Suggested)

- `↑/↓` or `j/k` - Scroll up/down
- `g/G` - Jump to top/bottom
- `/` - Enter search mode
- `i/w/e/d` - Toggle INFO/WARN/ERROR/DEBUG
- `c` - Clear all filters
- `r` - Refresh/reload file
- `?` - Show help
- `q` - Quit

## Testing Strategy

1. **Generate test logs**: Create sample files with various log levels
2. **Stress test**: Generate large log files (10k+ lines)
3. **Real-time test**: Use a script to append logs while running
4. **Edge cases**: Empty files, malformed logs, rapid updates

## Success Criteria

You've completed this project when:
- ✅ You can monitor a log file in real-time
- ✅ Filtering by log level works smoothly
- ✅ Search highlights matching text
- ✅ UI remains responsive with large files
- ✅ Code is well-organized and documented

## Next Steps After Completion

This project prepares you for:
- **Week 8**: Metrics Collection Server (similar concepts, different scale)
- **Week 15**: Game Analytics Pipeline (processing event streams)
- Production debugging and monitoring in cloud environments

## Resources

- [Bubbletea Documentation](https://github.com/charmbracelet/bubbletea)
- [fsnotify Examples](https://github.com/fsnotify/fsnotify)
- [Go Regex Guide](https://pkg.go.dev/regexp)

---

**Remember**: Every game server generates logs. Understanding how to build tools to analyze them is a critical skill. This project gives you insight into how professional monitoring tools work!

Now go build something awesome! 🚀