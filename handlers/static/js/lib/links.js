import { basePath } from "./paths.js";

export function makeShortLink(fileId) {
  const b = basePath();
  return `${window.location.origin}${b === "/" ? "" : b}/-${fileId}`;
}

export function makeVerboseLink(fileId, filename) {
  return makeShortLink(fileId) + "/" + encodeURIComponent(filename);
}
