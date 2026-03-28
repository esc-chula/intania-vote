import { AuthOptions, getServerSession } from "next-auth";
import GoogleProvider from "next-auth/providers/google";

import { login } from "~/server/user";

const authOptions: AuthOptions = {
  providers: [
    GoogleProvider({
      clientId: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
      authorization: {
        params: {
          hd: "student.chula.ac.th",
        },
      },
    }),
  ],
  callbacks: {
    async signIn({ user, account, profile }) {
      if (!user.email || !user.email.endsWith("@student.chula.ac.th")) {
        return false;
      }

      const studentId = user.email.split("@")[0];
      const oidcId = user.id;

      await login({
        oidcId,
        studentId,
      }).catch((error) => {
        console.error("Error in login action:", error);
      });

      return true;
    },
    async session({ session, token }) {
      if (token.studentId) {
        session.user.studentId = token.studentId as string;
      }
      if (token.oidcId) {
        session.user.oidcId = token.oidcId as string;
      }
      return session;
    },
    async jwt({ token, user, account }) {
      if (user && user.email) {
        token.studentId = user.email.split("@")[0];
        token.oidcId = user.id;
      }
      return token;
    },
  },
  pages: {
    signIn: "/login",
  },
};

const getSession = () => getServerSession(authOptions);

export { getSession, authOptions };
