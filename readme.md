# 🖥️ Linux Server Monitor (Go)

A lightweight REST API application built in Golang to monitor Linux server health in real-time.

---

## 🚀 Features

- 📈 **CPU Usage** – Real-time percentage
- 💾 **Memory Usage** – Used memory percentage
- 🧮 **Disk Usage** – Disk space used on root (`/`) partition
- ⏱️ **Uptime** – How long the system has been running
- 🌐 **REST API Endpoints** – `/stats`, `/uptime`, `/health`, `/metrics`
- 📊 **Prometheus Metrics Exporter** – Exposes real-time resource stats
- 🔁 **Auto-updating Metrics** – Background goroutine updates every 5s

---

## 📁 Project Structure

```

server-monitor/
├── main.go                 # Starts the HTTP server
├── monitor/                # Contains system metric logic
│   ├── cpu.go
│   ├── memory.go
│   ├── disk.go
│   └── uptime.go
└── api/
└── server.go           # Exposes API routes and metrics

````

---

## 📦 Requirements

- Go 1.19+
- Linux or WSL2 environment

---

## 🔧 Setup & Run

```bash
# Clone the project
git clone https://github.com/yourusername/server-monitor.git
cd server-monitor

# Initialize dependencies
go mod tidy

# Run the app
go run main.go

# From another machine or browser
curl http://<your-server-ip>:8080/stats
````

---

## 🌐 API Endpoints

| Method | Endpoint   | Description                            |
| ------ | ---------- | -------------------------------------- |
| GET    | `/stats`   | Returns CPU, RAM, Disk, Uptime as JSON |
| GET    | `/uptime`  | Returns system uptime in plain text    |
| GET    | `/health`  | Health check for service monitoring    |
| GET    | `/metrics` | Prometheus-compatible metrics endpoint |

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

```
up 40 minutes
```

---

## 🛠️ Work in Progress

* [x] Add `/health` endpoint
* [x] Add Prometheus `/metrics` endpoint
* [x] Auto-updating metric values (background goroutine)
* [ ] Logging & request audit
* [ ] Docker support
