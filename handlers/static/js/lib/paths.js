export function basePath() {
  // window.__BASE_PATH__ is injected by the server-side templates.
  // Normalize: remove trailing slash except for root.
  let b = (window.__BASE_PATH__ || "/").trim();
  if (b !== "/" && b.endsWith("/")) {
    b = b.replace(/\/+$/, "");
  }
  return b;
}

export function apiPath(path) {
  const b = basePath();
  if (!path.startsWith("/")) {
    path = "/" + path;
  }
  if (b === "/") {
    return path;
  }
  return b + path;
}
