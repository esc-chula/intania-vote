/* eslint-disable @typescript-eslint/no-explicit-any */
// eslint-disable-next-line @typescript-eslint/no-unused-vars
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
      id: number;
      oidcId: string;
      studentId: string;
    };
  }
  interface Profile {
    student_id: string;
  }
}
