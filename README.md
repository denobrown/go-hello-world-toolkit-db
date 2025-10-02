git clone https://github.com/your-username/go-hello-world-toolkit.git
cd go-hello-world-toolkit
2. Initialize Go Module
bash
go mod init go-hello-world
3. Run the Server
bash
go run main.go
4. Test the Endpoints
Open your browser and visit:

http://localhost:8080 - Welcome page

http://localhost:8080/api - JSON API

http://localhost:8080/health - Health check

http://localhost:8080/time - Current time

Or use curl:

bash
curl http://localhost:8080/api
curl http://localhost:8080/health
📁 Project Structure
text
go-hello-world-toolkit/
├── main.go              # Main server application
├── go.mod              # Go module file
├── README.md           # Project documentation
└── examples/
    └── advanced-server.go  # More advanced example
🚀 Available Endpoints
Endpoint	Method	Description	Response
/	GET	Welcome page	HTML
/api	GET	JSON API	JSON
/health	GET	Health status	JSON
/time	GET	Current time	Plain text
🔧 Building and Deployment
Build Executable
bash
go build -o hello-server
Run Executable
bash
./hello-server
Cross-Platform Building
bash
# Linux
GOOS=linux GOARCH=amd64 go build -o hello-server-linux

# Windows
GOOS=windows GOARCH=amd64 go build -o hello-server-windows.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o hello-server-macos
🐛 Common Issues & Solutions
Port Already in Use
bash
# Error: listen tcp :8080: bind: address already in use
# Solution: Change port or kill existing process
lsof -ti:8080 | xargs kill -9
# Or change port in main.go
Go Module Issues
bash
# Initialize module if not exists
go mod init your-project-name

# Download dependencies
go mod tidy
Permission Denied (Linux/Mac)
bash
# Make executable and run
chmod +x hello-server
./hello-server
📚 Learning Resources
Go Official Documentation

Go by Example

Effective Go

Go Standard Library

🤝 Contributing
This is a learning project. Feel free to fork and experiment with:

Adding new endpoints

Implementing middleware

Adding database connectivity

Creating frontend interfaces

📄 License
This project is created for educational purposes as part of the Moringa AI Capstone Project.

🎓 About Moringa AI Capstone
Part of the "Prompt-Powered Kickstart: Building a Beginner's Toolkit for Go" project, demonstrating how to use AI prompts to learn new technologies effectively.

Happy Coding! 🎉