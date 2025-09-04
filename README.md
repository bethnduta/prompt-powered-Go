**Getting Started with Go (Golang) – Beginner’s Toolkit**

This is a beginner-friendly toolkit for learning Go (Golang). It walks you through installing Go, running your first “Hello World” program, and understanding common errors.

Project Objective

Learn the basics of Go using AI prompts.

Create a minimal working project (Hello World).

Document setup, issues, and fixes for beginners.

Overview of Go

Go (or Golang) is a statically typed, compiled language designed by Google. It’s widely used in backend development, APIs, cloud-native systems, and DevOps tools.

Real-world example: Docker and Kubernetes are built in Go.

System Requirements

OS: Windows, Linux, or macOS

Editor: VS Code (recommended)

Tools: Go compiler (download from go.dev/dl
)

⚡ Installation & Setup
Step 1: Install Go

Download from go.dev/dl
.

Follow the installation wizard.

Step 2: Verify Installation
go version


Expected output (version may vary):

go version go1.22.0 windows/amd64

Step 3: Setup Workspace
mkdir go_projects
cd go_projects

Minimal Working Example

Create a file main.go with the following code:

package main

import "fmt"

func main() {
    fmt.Println("Hello, World from Go!")
}


Run the program:

go run main.go


Expected output:

Hello, World from Go!

AI Prompt Journal

Prompt 1: “Give me a step-by-step guide to install Go (Golang) on Windows.”

Prompt 2: “Write a Hello World program in Go and explain the code.”

Prompt 3: “What are common beginner errors in Go installation?”

The AI responses guided installation, helped understand code syntax, and resolved PATH issues.

Common Issues & Fixes
Issue	Cause	Fix
go: command not found	PATH not set	Add Go binary to environment variables
VS Code not detecting Go	Missing extension	Install “Go” extension in VS Code
Program not running	Wrong directory	Run inside folder with main.go
References

Official Go Documentation

Go by Example

VS Code Go Extension

Deliverables

Code → main.go

Documentation → Toolkit PDF/Markdown (this README can be part of it)

Run Instructions → go run main.go
