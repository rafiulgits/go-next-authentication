import BackendService from "services/backend";

const googleSignupConfig = {
  id: "google-signup",
  name: "google",
  clientId: process.env.GOOGLE_AUTH_CLIENT_ID!,
  clientSecret: process.env.GOOGLE_AUTH_CLIENT_SECRET!,
};

const googleSigninConfig = {
  id: "google-signin",
  name: "google",
  clientId: process.env.GOOGLE_AUTH_CLIENT_ID!,
  clientSecret: process.env.GOOGLE_AUTH_CLIENT_SECRET!,
};

const authHandler = async (ctx: any) => {
  const authType = ctx.account.provider;
  if (authType === googleSigninConfig.id) {
    try {
      const data = {
        provider: "google",
        accessToken: ctx.account.access_token,
      };
      await BackendService.loginOAuth(data);
      return true;
    } catch (err: any) {
      return `/auth/login?error=${err.message}`;
    }
  }

  if (authType === googleSignupConfig.id) {
    try {
      const data = {
        provider: "google",
        accessToken: ctx.account.access_token,
      };
      await BackendService.signupOAuth(data);
      return true;
    } catch (err: any) {
      return `/auth/signup?error=${err.message}`;
    }
  }

  throw new Error("provider not found");
};

const getUserProfile = async (ctx: any) => {
  const data = {
    provider: "google",
    accessToken: ctx.account.access_token,
  };

  const access = await BackendService.loginOAuth(data);
  const profile = await BackendService.getProfile(access.bearer);

  return {
    ...profile,
    accessToken: access.bearer,
  };
};

const isGoogleProvider = (provider: string) => {
  return (
    provider === googleSignupConfig.id || provider === googleSigninConfig.id
  );
};

const GoogleAuth = {
  signinConfig: googleSigninConfig,
  signupConfig: googleSignupConfig,
  handle: authHandler,
  getProfile: getUserProfile,
  canProvide: isGoogleProvider,
};

export default GoogleAuth;
