import { Link } from "@tanstack/react-router";
import { ArrowRight, Box, CheckCircle, ExternalLink, Terminal, Zap } from "lucide-react";

export default function HomePage() {
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
            <Link to="/docs" className="rounded text-sm font-medium text-neutral-400 transition-colors hover:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500">
              Documentation
            </Link>
          </div>
        </div>
      </nav>

      <main>
        <section className="pt-32 pb-20 px-6">
          <div className="mx-auto grid max-w-7xl items-center gap-16 md:grid-cols-[1.15fr_0.85fr]">
            <div className="animate-fade-in">
              <p className="tech-label mb-6">[ spring.initializr / terminal interface ]</p>
              <h1 className="mb-7 max-w-3xl font-mono text-4xl font-bold uppercase leading-[0.96] tracking-[-0.07em] text-neutral-100 sm:text-5xl lg:text-7xl">
                Boot projects.
                <span className="block text-primary-500">Skip ceremony.</span>
              </h1>
              <p className="mb-9 max-w-xl border-l-2 border-primary-600 pl-5 text-base leading-relaxed text-neutral-400 md:text-lg">
                An interactive tool that wraps the Spring Initializr API. No web forms, no ZIP-file cleanup. Configure a production-ready project without leaving the terminal.
              </p>
              <div className="flex flex-wrap items-center gap-5">
                <Link to="/docs" className="btn btn-secondary px-8 py-3.5">
                  Read the manual
                  <ArrowRight className="ml-3 h-4 w-4" />
                </Link>
                <a
                  href="https://github.com/SamSyntax/create-spring-app"
                  target="_blank"
                  rel="noopener noreferrer"
                  className="inline-flex items-center gap-2 rounded text-sm text-neutral-400 transition-colors hover:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500"
                >
                  <ExternalLink className="h-4 w-4" />
                  Source
                </a>
              </div>
            </div>

            <div className="brutal-panel animate-fade-in min-w-0 p-1" aria-label="Terminal installation preview">
              <div className="flex items-center justify-between border-b border-neutral-700 bg-neutral-900 px-4 py-3 font-mono text-[0.65rem] uppercase tracking-[0.18em] text-neutral-500">
                <span>csa://session</span>
                <span className="text-primary-500">live</span>
              </div>
              <div className="space-y-5 p-5 font-mono text-xs leading-6 sm:p-7 sm:text-sm">
                <p><span className="text-primary-500">$</span> go install github.com/SamSyntax/create-spring-app@latest</p>
                <div className="border-l border-neutral-700 pl-4 text-neutral-500">
                  <p>resolving packages........................done</p>
                  <p>building binary..........................done</p>
                </div>
                <p><span className="text-primary-500">$</span> csa</p>
                <div className="grid grid-cols-[auto_1fr] gap-x-5 gap-y-2 text-neutral-400">
                  <span className="text-neutral-600">01</span><span>PROJECT <b className="font-normal text-neutral-100">signal-api</b></span>
                  <span className="text-neutral-600">02</span><span>BUILD <b className="font-normal text-primary-400">gradle-kotlin</b></span>
                  <span className="text-neutral-600">03</span><span>JAVA <b className="font-normal text-neutral-100">25</b></span>
                </div>
                <p className="flex items-center gap-2 text-neutral-200"><span className="h-2 w-2 animate-pulse bg-primary-500" /> Ready to generate_</p>
              </div>
            </div>
          </div>
        </section>

        <section className="border-y border-neutral-800 bg-neutral-900/60 px-6 py-20">
          <div className="max-w-6xl mx-auto">
            <div className="mb-12 grid gap-4 border-b border-neutral-800 pb-8 md:grid-cols-2 md:items-end">
              <h2 className="section-title mb-4">Why Use Create Spring App?</h2>
              <p className="section-description max-w-xl md:justify-self-end">
                Built for developers who live in the terminal but want a better experience than start.spring.io
              </p>
            </div>
            <div className="grid gap-px border border-neutral-800 bg-neutral-800 md:grid-cols-12">
              <FeatureCard
                icon={<Terminal className="w-6 h-6" />}
                title="Keyboard-Driven"
                description="Navigate entirely with your keyboard. No mouse required. Built with Bubble Tea for a modern TUI experience."
                className="md:col-span-7"
              />
              <FeatureCard
                icon={<Zap className="w-6 h-6" />}
                title="Smart Defaults"
                description="Auto-infers reasonable defaults based on your selections. Spend less time configuring, more time coding."
                className="md:col-span-5"
              />
              <FeatureCard
                icon={<CheckCircle className="w-6 h-6" />}
                title="Compatibility Checked"
                description="Dependencies are validated against your selected Spring Boot version. No more runtime errors from incompatible libraries."
                className="md:col-span-12"
              />
            </div>
          </div>
        </section>

        <section className="py-20 px-6">
          <div className="max-w-4xl mx-auto">
            <div className="brutal-panel p-8 md:p-12">
              <p className="tech-label mb-3">[ 00 / install ]</p>
              <h2 className="section-title mb-8">Quick Installation</h2>
              <div className="mb-6 rounded-lg border border-neutral-800 bg-neutral-950 p-6">
                <pre className="text-sm text-neutral-300 font-mono overflow-x-auto">
                  <code>go install github.com/SamSyntax/create-spring-app@latest</code>
                </pre>
              </div>
              <div className="text-center">
                <Link to="/docs" className="btn btn-ghost text-primary-600">
                  View Full Installation Guide
                  <ArrowRight className="w-4 h-4 ml-2" />
                </Link>
              </div>
            </div>
          </div>
        </section>

        <section className="border-t border-neutral-800 bg-neutral-900/60 px-6 py-20">
          <div className="max-w-6xl mx-auto">
            <div className="mb-12">
              <p className="tech-label mb-3">[ 01—04 / execution path ]</p>
              <h2 className="section-title mb-4">How It Works</h2>
              <p className="section-description">Four simple steps to your new Spring Boot project</p>
            </div>
            <div className="grid border-l border-t border-neutral-800 sm:grid-cols-2 lg:grid-cols-4">
              <StepCard number="1" title="Run the Command" description="Execute csa in your terminal to start the interactive wizard" />
              <StepCard number="2" title="Enter Details" description="Provide project name, group, and package information" />
              <StepCard number="3" title="Select Options" description="Choose Spring Boot version, build tool, and dependencies" />
              <StepCard number="4" title="Done" description="Project is downloaded and extracted automatically" />
            </div>
          </div>
        </section>
      </main>

      <footer className="border-t border-neutral-800 bg-neutral-950 px-6 py-12 text-neutral-400">
        <div className="max-w-6xl mx-auto flex flex-col md:flex-row items-center justify-between gap-4">
          <div className="flex items-center gap-2">
            <Terminal className="w-5 h-5" />
            <span className="font-medium text-white">create-spring-app</span>
          </div>
          <div className="flex items-center gap-6 text-sm">
            <Link to="/docs" className="rounded transition-colors hover:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500">
              Documentation
            </Link>
            <a
              href="https://github.com/SamSyntax/create-spring-app"
              target="_blank"
              rel="noopener noreferrer"
              className="rounded transition-colors hover:text-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500"
            >
              GitHub
            </a>
          </div>
          <p className="text-sm">Open Source - MIT License</p>
        </div>
      </footer>
    </div>
  );
}

function FeatureCard({
  icon,
  title,
  description,
  className,
}: {
  icon: React.ReactNode;
  title: string;
  description: string;
  className: string;
}) {
  return (
    <div className={`bg-neutral-950 p-7 transition-colors hover:bg-neutral-900 ${className}`}>
      <div className="mb-8 flex h-11 w-11 items-center justify-center border border-primary-900 bg-primary-950/60 text-primary-400">
        {icon}
      </div>
      <h3 className="mb-2 text-lg font-semibold text-neutral-100">{title}</h3>
      <p className="text-neutral-400">{description}</p>
    </div>
  );
}

function StepCard({
  number,
  title,
  description,
}: {
  number: string;
  title: string;
  description: string;
}) {
  return (
    <div className="relative border-b border-r border-neutral-800 p-6">
      <div className="mb-8 flex items-center justify-between font-mono text-xs text-primary-500">
        <span>STEP/{number.padStart(2, "0")}</span>
        <Box className="h-4 w-4 text-neutral-700" />
      </div>
      <h3 className="mb-2 text-lg font-semibold text-neutral-100">{title}</h3>
      <p className="text-sm text-neutral-400">{description}</p>
    </div>
  );
}
