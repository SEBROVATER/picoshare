import { apiPath } from "../lib/paths.js";

export async function authenticate(passphrase) {
  return fetch(apiPath("/api/auth"), {
    method: "POST",
    mode: "same-origin",
    credentials: "include",
    cache: "no-cache",
    redirect: "error",
    body: JSON.stringify({
      sharedSecretKey: passphrase,
    }),
  }).then((response) => {
    if (!response.ok) {
      return response.text().then((error) => {
        return Promise.reject(error);
      });
    }
    return Promise.resolve();
  });
}

export function logOut() {
  return fetch(apiPath("/api/auth"), {
    method: "DELETE",
    mode: "same-origin",
    credentials: "include",
    cache: "no-cache",
    redirect: "error",
  }).then((response) => {
    if (!response.ok) {
      return response.text().then((error) => {
        return Promise.reject(error);
      });
    }
    return Promise.resolve();
  });
}
