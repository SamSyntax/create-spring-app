import { createRootRoute, createRoute, Outlet } from "@tanstack/react-router";
import HomePage from "./routes/index";
import DocsPage from "./routes/docs";

const rootRoute = createRootRoute({
  component: () => (
    <div className="min-h-[100dvh] bg-neutral-950 text-neutral-100">
      <Outlet />
    </div>
  ),
});

const indexRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/",
  component: HomePage,
});

const docsRoute = createRoute({
  getParentRoute: () => rootRoute,
  path: "/docs",
  component: DocsPage,
});

const routeTree = rootRoute.addChildren([indexRoute, docsRoute]);

export { routeTree };
