# 🚀 Go HTTP Server Toolkit

A comprehensive demonstration project showcasing different approaches to building HTTP servers in Go, from basic implementations to advanced patterns with middleware and concurrency.

## 📋 Description

This project illustrates why structured frameworks and patterns are essential for building maintainable HTTP servers in Go. It demonstrates three levels of server implementation:

1. **Basic Server** - Shows fundamental HTTP server capabilities
2. **Advanced Server with Middleware** - Demonstrates enhanced patterns with logging, metrics, and concurrency
3. **Specialized Examples** - Practical implementations for common use cases

## 🚀 Features

- **Comparative Examples**: Side-by-side comparison of different server architectures
- **Multiple Server Types**: Basic, advanced, REST API, and file servers
- **Middleware Implementation**: Request logging, timing, and concurrency safety
- **Interactive Demos**: Real-time servers with various endpoints
- **Comprehensive Documentation**: Setup guides, usage examples, and troubleshooting

## 🛠️ Technologies Used

- **Go 1.21+**: Modern Go with latest features
- **net/http**: Standard library HTTP package
- **encoding/json**: JSON serialization and deserialization
- **sync**: Concurrency primitives for safe operations

## 📦 Installation

### Prerequisites

- Go 1.21 or higher
- Git for version control

### Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/go-http-server-toolkit.git
   cd go-http-server-toolkit
   Initialize Go module

bash
go mod init go-http-server-toolkit
Verify installation


## 📁 Project Structure

```
go-http-server-toolkit/
├── main.go                 // Primary server implementation
├── go.mod                 // Go module configuration
├── README.md              // Project documentation
├── .gitignore             // Git ignore rules
├── basic-server/          // Simple HTTP server examples
│   ├── basic-server.go
│   └── README.md
├── advanced-server/       // Enhanced server with middleware
│   ├── advanced-server.go
│   └── README.md
├── examples/              // Additional implementations
│   ├── simple-api.go
│   ├── file-server.go
│   └── README.md
└── docs/                  // Comprehensive documentation
    ├── setup.md
    ├── usage.md
    └── troubleshooting.md
    ## 🎯 Usage

### Running the Examples

#### Basic Examples
Run the main demonstration file to see all three approaches:

```bash
python app.py
```

This will execute:
- Example 1: Basic chain without memory
- Example 2: Manual memory management
- Example 3: LangChain memory integration

#### Interactive Chat Demo
For a hands-on conversation experience:

```bash
python interactive_demo.py --interactive
```

Commands during chat:
- Type your questions naturally
- Type `quit`, `exit`, or `q` to end the session
- Press `Ctrl+C` to force quit
```
## ⚙️ Configuration Options

### Model Configuration
Adjust the Groq model settings in your code:

```python
llm = ChatGroq(
    model="llama3-8b-8192",      # Model name
    temperature=0.7,              # Creativity (0.0-1.0)
    max_tokens=300,              # Response length limit
    api_key=APIKEY               # Your API key
)
```

## 🔧 Troubleshooting

### Common Issues

#### "API Key not found"
```bash
Error: API key not found
```
**Solution**: Ensure your `.env` file contains `GROQ_API_KEY=your_actual_key`

#### "Rate limit exceeded"
```bash
Error: Rate limit exceeded
```
**Solution**: Wait a moment before making another request, or check your Groq account limits.

#### "Connection timeout"
```bash
Error: Connection timeout
```
**Solution**: Check your internet connection and try again.

### Debug Mode
Enable verbose logging by setting environment variable:
```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}

```

## 🤝 Contributing

We welcome contributions! Here's how to get started:

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**
4. **Add tests if applicable**
5. **Commit your changes**
   ```bash
   git commit -m 'Add some amazing feature'
   ```
6. **Push to the branch**
   ```bash
   git push origin feature/amazing-feature
   ```
7. **Open a Pull Request**

### Contribution Guidelines

- Follow Go standard formatting (go fmt)
- Add comments for complex logic
- Include examples for new features
- Update README.md if needed
- Test your changes locally

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Moringa School - For the AI Capstone project opportunity and guidance
- Go Developers - For the excellent standard library and documentation
- The open-source community for continuous inspiration

## 📞 Support

If you encounter any issues or have questions:

1. Review the documentation for detailed guides
2. Search existing [Issues](../../issues)
3. Create a new issue if your problem isn't covered
4. Join our community discussions

## 🔗 Useful Links

- Go Official Documentation
- Go by Example
- Go Standard Library

---

**Happy coding!** 🎉