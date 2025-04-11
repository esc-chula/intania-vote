export { default } from "next-auth/middleware";

// match everything except for the login page
export const config = {
  matcher: ["/((?!_next/static|_next/image|favicon.ico|assets|login).*)"],
};
