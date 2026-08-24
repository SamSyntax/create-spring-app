import { Link } from "@tanstack/react-router";
import { Terminal, BookOpen, ArrowLeft, CheckCircle, ExternalLink, Copy, Check } from "lucide-react";
import { useState } from "react";

export default function DocsPage() {
  return (
    <div className="tech-grid min-h-[100dvh] bg-neutral-950 text-neutral-100">
      <nav className="fixed top-0 left-0 right-0 z-50 border-b border-neutral-800 bg-neutral-950/85 backdrop-blur-md">
        <div className="max-w-7xl mx-auto px-6 h-16 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <span className="flex h-8 w-8 items-center justify-center border border-primary-700 bg-primary-950/60">
              <Terminal className="h-4 w-4 text-primary-400" />
            </span>
            <span className="font-mono text-sm font-semibold uppercase tracking-tight text-neutral-100">create-spring-app</span>
          </div>
          <div className="flex items-center gap-6">
            <Link to="/" className="rounded text-sm font-medium text-neutral-400 transition-colors hover:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500">
              Home
            </Link>
          </div>
        </div>
      </nav>

      <div className="pt-24 pb-16 px-6">
        <div className="mx-auto max-w-7xl">
          <Link to="/" className="mb-8 inline-flex items-center gap-2 rounded text-sm text-neutral-400 transition-colors hover:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500">
            <ArrowLeft className="w-4 h-4" />
            Back to Home
          </Link>

          <header className="brutal-panel mb-12 p-7 md:p-10">
            <p className="tech-label mb-4">[ reference / v1.0 ]</p>
            <div className="mb-4 flex items-center gap-4">
              <BookOpen className="h-8 w-8 text-primary-500" />
              <h1 className="font-mono text-4xl font-bold uppercase tracking-[-0.06em] text-neutral-100 md:text-5xl">Documentation</h1>
            </div>
            <p className="text-xl text-neutral-400">
              Everything you need to know to get started with create-spring-app
            </p>
          </header>

          <div className="grid items-start gap-10 lg:grid-cols-[13rem_1fr]">
            <aside className="border border-neutral-800 bg-neutral-950/90 p-5 lg:sticky lg:top-24">
              <p className="tech-label mb-5">Index</p>
              <nav className="flex flex-col font-mono text-xs uppercase tracking-wider text-neutral-500">
                {[
                  ["overview", "01 / Overview"],
                  ["prerequisites", "02 / Prerequisites"],
                  ["installation", "03 / Installation"],
                  ["usage", "04 / Usage"],
                  ["troubleshooting", "05 / Troubleshooting"],
                  ["source", "06 / Source"],
                ].map(([href, label]) => (
                  <a key={href} href={`#${href}`} className="border-t border-neutral-800 py-3 transition-colors hover:text-primary-400 focus-visible:outline-none focus-visible:text-primary-400">
                    {label}
                  </a>
                ))}
              </nav>
            </aside>
            <div className="min-w-0 space-y-6">
            <Section title="Overview" id="overview">
              <p className="mb-4 text-neutral-400">
                <strong>create-spring-app</strong> is an interactive CLI tool that wraps the Spring Initializr API,
                providing a modern, keyboard-driven Terminal User Interface (TUI) for generating Spring Boot projects.
                It eliminates the need to use the web-based start.spring.io interface.
              </p>
              <p className="text-neutral-400">
                Built with Go and Bubble Tea, it offers a seamless experience for developers who prefer working
                in the terminal but want a better alternative to manually downloading and configuring projects.
              </p>
            </Section>

            <Section title="Prerequisites" id="prerequisites">
              <p className="mb-6 text-neutral-400">
                Before installing create-spring-app, ensure you have the following requirements met:
              </p>
              <div className="space-y-4">
                <RequirementCard
                  title="Go 1.25 or later"
                  description="The tool is built with Go. You can check your version by running `go version` in your terminal."
                  command="go version"
                />
                <RequirementCard
                  title="Internet Connection"
                  description="Required to fetch metadata from the Spring Initializr API and download generated projects."
                />
                <RequirementCard
                  title="tar Command"
                  description="Used for extracting the downloaded project archive. Available by default on most Unix-like systems."
                  command="tar --version"
                />
              </div>
            </Section>

            <Section title="Installation" id="installation">
              <p className="mb-6 text-neutral-400">
                There are two ways to install create-spring-app. Choose the option that best fits your needs:
              </p>

              <div className="space-y-6">
                <InstallationCard
                  method="Go Install (Recommended)"
                  description="The quickest way to install the latest release using Go's built-in package manager."
                  command="go install github.com/SamSyntax/create-spring-app@latest"
                  note="Ensure your Go bin directory is in your PATH (usually ~/go/bin)"
                />
                <InstallationCard
                  method="Build from Source"
                  description="Clone the repository and build from source for the latest development version."
                  commandLines={[
                    "git clone https://github.com/SamSyntax/create-spring-app.git",
                    "cd create-spring-app",
                    "make build",
                    "mv bin/csa /usr/local/bin/",
                  ]}
                  note="Requires Git and Make to be installed"
                />
              </div>
            </Section>

            <Section title="Verifying Installation" id="verify">
              <p className="mb-4 text-neutral-400">
                After installation, verify that create-spring-app is properly installed by checking its version:
              </p>
              <CommandBlock command="csa --version" />
              <p className="mt-4 text-neutral-400">
                You should see version information displayed. If you get a "command not found" error,
                make sure the installation directory is in your PATH.
              </p>
            </Section>

            <Section title="Usage" id="usage">
              <p className="mb-4 text-neutral-400">
                To start creating a new Spring Boot project, simply run:
              </p>
              <CommandBlock command="csa" />
              <p className="mb-6 mt-4 text-neutral-400">
                This will launch an interactive wizard that guides you through:
              </p>
              <ul className="mb-6 list-inside list-disc space-y-2 text-neutral-400 marker:text-primary-600">
                <li>Entering project name (artifact ID)</li>
                <li>Entering group name (group ID)</li>
                <li>Customizing package name</li>
                <li>Selecting Spring Boot version</li>
                <li>Choosing build tool (Maven or Gradle)</li>
                <li>Selecting Java version</li>
                <li>Choosing dependencies with search and filtering</li>
              </ul>
            </Section>

            <Section title="Available Options" id="options">
              <p className="mb-4 text-neutral-400">
                The CLI is entirely interactive and does not accept command-line flags. All configuration
                is done through the TUI interface. This ensures a consistent and guided experience for all users.
              </p>
            </Section>

            <Section title="Dependencies" id="dependencies">
              <p className="mb-4 text-neutral-400">
                The tool fetches available dependencies dynamically from the Spring Initializr API, ensuring
                you always have access to the latest compatible libraries. Dependencies are validated against
                your selected Spring Boot version to prevent compatibility issues.
              </p>
            </Section>

            <Section title="Output" id="output">
              <p className="mb-4 text-neutral-400">
                After completing the wizard, the tool:
              </p>
              <ol className="mb-6 list-inside list-decimal space-y-2 text-neutral-400 marker:text-primary-600">
                <li>Constructs a request to the Spring Initializr API</li>
                <li>Downloads the generated project as a tar.gz archive</li>
                <li>Extracts the project to your current directory</li>
              </ol>
              <p className="text-neutral-400">
                The project is ready to use immediately with no additional configuration required.
              </p>
            </Section>

            <Section title="Troubleshooting" id="troubleshooting">
              <div className="space-y-4">
                <TroubleshootingItem
                  problem="Command not found"
                  solution="Add ~/go/bin to your PATH, or move the csa binary to a directory in your PATH like /usr/local/bin"
                />
                <TroubleshootingItem
                  problem="Metadata fetch fails"
                  solution="Check your internet connection. The tool requires access to https://start.spring.io"
                />
                <TroubleshootingItem
                  problem="tar: command not found"
                  solution="Install tar using your system's package manager (e.g., apt install tar on Ubuntu)"
                />
                <TroubleshootingItem
                  problem="Dependency compatibility warnings"
                  solution="Some dependencies may not be compatible with your selected Spring Boot version. Consider using a different version or removing the incompatible dependency"
                />
              </div>
            </Section>

            <Section title="Source Code" id="source">
              <p className="mb-4 text-neutral-400">
                Create Spring App is open source and available on GitHub. Contributions are welcome!
              </p>
              <a
                href="https://github.com/SamSyntax/create-spring-app"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-2 rounded font-medium text-primary-400 transition-colors hover:text-primary-300 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500"
              >
                View on GitHub
                <ExternalLink className="w-4 h-4" />
              </a>
            </Section>
            </div>
          </div>
        </div>
      </div>

      <footer className="border-t border-neutral-800 bg-neutral-950 px-6 py-8 text-neutral-400">
        <div className="max-w-4xl mx-auto flex flex-col md:flex-row items-center justify-between gap-4">
          <div className="flex items-center gap-2">
            <Terminal className="w-5 h-5" />
            <span className="font-medium text-white">create-spring-app</span>
          </div>
          <p className="text-sm">Open Source - MIT License</p>
        </div>
      </footer>
    </div>
  );
}

function Section({ title, children, id }: { title: string; children: React.ReactNode; id?: string }) {
  return (
    <section id={id} className="brutal-panel scroll-mt-24 p-6 md:p-8">
      <p className="tech-label mb-3">[ module / {id ?? "section"} ]</p>
      <h2 className="mb-6 border-b border-neutral-800 pb-4 font-mono text-2xl font-semibold uppercase tracking-[-0.04em] text-neutral-100">
        {title}
      </h2>
      {children}
    </section>
  );
}

function RequirementCard({
  title,
  description,
  command,
}: {
  title: string;
  description: string;
  command?: string;
}) {
  return (
    <div className="card p-5">
      <div className="flex items-start gap-4">
        <div className="mt-0.5 flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-full border border-primary-900 bg-primary-950/60">
          <CheckCircle className="w-5 h-5 text-primary-600" />
        </div>
        <div>
          <h3 className="mb-1 font-semibold text-neutral-100">{title}</h3>
          <p className="mb-2 text-neutral-400">{description}</p>
          {command && <CommandBlock command={command} compact />}
        </div>
      </div>
    </div>
  );
}

function InstallationCard({
  method,
  description,
  command,
  commandLines,
  note,
}: {
  method: string;
  description: string;
  command?: string;
  commandLines?: string[];
  note?: string;
}) {
  return (
    <div className="card overflow-hidden">
      <div className="border-b border-neutral-800 bg-neutral-800/70 px-5 py-3">
        <h3 className="font-semibold text-neutral-100">{method}</h3>
      </div>
      <div className="p-5">
        <p className="mb-4 text-neutral-400">{description}</p>
        {(command || commandLines) && (
          <div className="mb-4 rounded-lg border border-neutral-800 bg-neutral-950 p-4">
            <pre className="text-sm text-neutral-300 font-mono overflow-x-auto">
              {command || commandLines?.join("\n")}
            </pre>
          </div>
        )}
        {note && (
          <p className="flex items-center gap-2 text-sm text-neutral-500">
            <span className="w-1.5 h-1.5 rounded-full bg-neutral-400"></span>
            {note}
          </p>
        )}
      </div>
    </div>
  );
}

function CommandBlock({ command, compact }: { command: string; compact?: boolean }) {
  const [copied, setCopied] = useState(false);

  const copyToClipboard = async () => {
    await navigator.clipboard.writeText(command);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className={`overflow-hidden rounded-lg border border-neutral-800 bg-neutral-950 ${compact ? "p-3" : "p-4"}`}>
      <div className="flex items-center justify-between gap-4">
        <pre className={`${compact ? "text-xs" : "text-sm"} text-neutral-300 font-mono overflow-x-auto`}>
          <code>{command}</code>
        </pre>
        <button
          onClick={copyToClipboard}
          className="flex-shrink-0 rounded-md p-2 transition-colors hover:bg-neutral-800 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 active:translate-y-px"
          aria-label="Copy command to clipboard"
          title="Copy to clipboard"
        >
          {copied ? (
            <Check className="w-4 h-4 text-green-500" />
          ) : (
            <Copy className="w-4 h-4 text-neutral-400" />
          )}
        </button>
      </div>
    </div>
  );
}

function TroubleshootingItem({ problem, solution }: { problem: string; solution: string }) {
  return (
    <div className="card p-5">
      <h3 className="mb-2 font-semibold text-neutral-100">Problem: {problem}</h3>
      <p className="text-neutral-400">
        <span className="font-medium text-neutral-200">Solution:</span> {solution}
      </p>
    </div>
  );
}
