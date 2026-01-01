# Create Spring App 

> "Because I was too lazy to find a better solution than spring generator"

Welcome to **Create Spring App**, the CLI tool for developers hate the initialization ceremony. If you've ever felt the soul-crushing despair of navigating to `start.spring.io`, clicking a bunch of boxes, downloading a zip, unzipping it, and then moving the files to where you actually wanted them... well, this tool is your therapy.

Honestly, this is just a wrapper around the Spring Initializr API that provides an interactive TUI.

This is a modern, interactive Terminal User Interface (TUI) that brings the power of the Spring Initializr directly to your command line. It's fast, it's pretty (thanks to [Charm](https://charm.sh/)), and it doesn't involve your browser's download folder.

## Features

- **Interactive TUI**: Navigate through project setup with your keyboard. No mouse required.
- **Dynamic Metadata**: Fetches the latest Spring Boot versions, Java versions, and dependencies directly from the source.
- **Dependency Management**: Search and select dependencies with filtering. It even checks compatibility with your chosen Spring Boot version so you don't break things before you start.
- **Build Tool Selection**: Maven, Gradle (Groovy), or Gradle (Kotlin)? You choose.
- **Smart Defaults**: Infers reasonable defaults based on your input to save you keystrokes.
- **Instant Extraction**: Downloads and sets up the project structure right where you are. No `.zip` debris left behind.

## Installation

### Using Go

If you have Go installed, you can grab the latest version directly:

```bash
go install github.com/SamSyntax/create-spring-app@latest
```
Then you can run it:

```bash
create-spring-app
```

### Building from Source

1.  Clone the repository:
    ```bash
    git clone https://github.com/SamSyntax/create-spring-app.git
    cd create-spring-app
    ```

2.  Build the binary:
    ```bash
    make build
    ```

3.  (Optional) Move it to your path:
    ```bash
    mv bin/csa /usr/local/bin/
    ```

## Usage

Simply run the command in your terminal:

```bash
csa
```

Follow the on-screen prompts to configure your project:

1.  **Project Name**: The artifact ID (e.g., `my-cool-api`).
2.  **Group Name**: The group ID (e.g., `com.syntax`).
3.  **Package Name**: Auto-generated but customizable.
4.  **Spring Boot Version**: Choose from the currently supported active versions.
5.  **Build Tool**: Pick your poison (Maven or Gradle).
6.  **Java Version**: Select your target LTS.
7.  **Dependencies**: Press `Space` to select dependencies (Web, JPA, Security, etc.) and `Enter` to confirm.

> **Note**: This tool communicates with the Spring Initializr API. An internet connection is required.

## Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - The fun, functional, and stateful terminal apps framework.
- [Huh?](https://github.com/charmbracelet/huh) - A lightweight form library for Bubble Tea.
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for nice terminal layouts.
- Go 1.25+

## Contributing

Found a bug? Want to add a feature? Feel free to open an issue or submit a pull request. We promise not to make you download a zip file to contribute.

1.  Fork it
2.  Create your feature branch (`git checkout -b feature/my-new-feature`)
3.  Commit your changes (`git commit -am 'Add some feature'`)
4.  Push to the branch (`git push origin feature/my-new-feature`)
5.  Create a new Pull Request
