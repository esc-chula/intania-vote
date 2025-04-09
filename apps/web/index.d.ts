/* eslint-disable @typescript-eslint/no-explicit-any */
import NextAuth from "next-auth";

declare module "*.svg" {
  const content: any;
  export const ReactComponent: any;
  export default content;
}

declare module "next-auth" {
  /**
   * Returned by `useSession`, `getSession` and received as a prop on the `SessionProvider` React Context
   */
  interface Session {
    user: {
      /** The user's postal address. */
      studentId: string;
    };
  }
  interface Profile {
    student_id: string;
  }
}
