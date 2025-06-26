# 🖥️ Linux Server Monitor (Go)

A lightweight REST API application built in Golang to monitor Linux server health in real-time.


## 🚀 Features

- 📈 **CPU Usage** – Real-time percentage.
- 💾 **Memory Usage** – Used memory percentage.
- 🧮 **Disk Usage** – Disk space used on root (`/`) partition.
- ⏱️ **Uptime** – How long the system has been running.
- 🌐 **REST API Endpoints** – Access system info from a browser or external tools.


## 📁 Project Structure


server-monitor/
├── main.go                # Starts the HTTP server
├── monitor/               # Contains system metric logic
│   ├── cpu.go
│   ├── memory.go
│   ├── disk.go
│   └── uptime.go
└── api/
└── server.go          # Exposes /stats and /uptime APIs


## 📦 Requirements

- Go 1.19+
- Linux or WSL2 environment

## 🔧 Setup & Run

```bash
# Clone the project
git clone 
cd server-monitor

# Initialize dependencies
go mod tidy

# Run the app
go run main.go


## 🌐 API Endpoints

| Endpoint  | Description                            |
| --------- | -------------------------------------- |
| `/stats`  | Returns CPU, RAM, Disk, Uptime as JSON |
| `/uptime` | Returns uptime as plain text           |

---

## 📸 Sample Output

### `GET /stats`

```json
{
  "uptime": "up 40 minutes",
  "cpu_usage": 1.34,
  "memory_usage": 5.12,
  "disk_usage": 22.4
}
```

### `GET /uptime`

up 40 minutes

## 🛠️ Work in Progress

* [ ] Add `/health` endpoint
* [ ] Add Prometheus `/metrics` endpoint
* [ ] Logging & request audit
* [ ] Docker support
