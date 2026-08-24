// app.config.ts
import { defineConfig } from "@tanstack/start/config";
var app_config_default = defineConfig({
  app: {
    head: {
      title: "Create Spring App - Interactive Spring Boot Project Generator",
      meta: [
        {
          name: "description",
          content: "Generate Spring Boot projects with ease. An interactive web interface for the create-spring-app CLI tool."
        }
      ],
      links: [
        { rel: "icon", href: "/favicon.ico" },
        {
          rel: "stylesheet",
          href: "https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap"
        }
      ]
    }
  },
  vite: {
    css: {
      postcss: "./postcss.config.cjs"
    }
  }
});
export {
  app_config_default as default
};
