import { AuthOptions, getServerSession } from "next-auth";
import type { Profile } from "next-auth";

import { login } from "~/server/user";

// Profile and JWT Payload fields is depend on the scope that you request
interface IntaniaProfile extends Profile {
  student_id: string; // In this case, we requested student_id scope
}

const authOptions: AuthOptions = {
  providers: [
    {
      id: "intania",
      name: "Intania",
      type: "oauth",
      wellKnown: "https://auth.intania.org/.well-known/openid-configuration",
      idToken: true,

      // Configure the scope as you created in the client creator
      authorization: { params: { scope: "openid student_id" } },
      // Client ID and Client Secret is provided in the client creator
      clientId: process.env.INTANIA_CLIENT_ID,
      clientSecret: process.env.INTANIA_CLIENT_SECRET,

      profile(profile) {
        return {
          id: profile.sub,
          studentId: profile.student_id,
        };
      },
    },
  ],
  callbacks: {
    async signIn({ profile }) {
      if (!profile || !profile.sub || !profile.student_id) {
        console.error("Invalid profile data:", profile);
        return false;
      }

      await login({
        oidcId: profile.sub,
        studentId: profile.student_id,
      }).catch((error) => {
        console.error("Error in login action:", error);
      });

      return true;
    },
    async session({ session, token }) {
      session.user.oidcId = token.sub as string;
      session.user.studentId = token.studentId as string;
      return session;
    },
    async jwt({ token, profile }) {
      const intaniaProfile = profile as IntaniaProfile;
      if (profile && intaniaProfile.student_id) {
        token.studentId = intaniaProfile.student_id;
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
