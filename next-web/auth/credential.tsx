import BackendService from "services/backend";

const credentialSigninConfig = {
  id: "credentials-signin",
  name: "credentials",
  credentials: {
    email: { label: "Email", type: "text" },
    password: { label: "Password", type: "password" },
  },
  authorize: async (credentials: any, req: any) => {
    const { callbackUrl, redirect, ...userData } = credentials;
    try {
      const loginData = await BackendService.loginCredential(userData);
      const profileData = await BackendService.getProfile(loginData.bearer);
      return {
        ...profileData,
        accessToken: loginData.bearer,
      };
    } catch (err: any) {
      throw new Error(err.message);
    }
  },
};

const credentialSignupConfig = {
  id: "credentials-signup",
  name: "credentials",
  credentials: {
    name: { label: "Name", type: "text" },
    email: { label: "Email", type: "email" },
    phone: { label: "Phone", type: "text" },
    password: { label: "Password", type: "password" },
  },
  authorize: async (credentials: any, req: any) => {
    const { callbackUrl, redirect, ...userData } = credentials;
    try {
      const registrationData = await BackendService.signupCredential(userData);
      return registrationData;
    } catch (err: any) {
      throw new Error(err.message);
    }
  },
};

const authHandler = async (ctx: any) => {
  return ctx.user;
};

const getUserProfile = async (ctx: any) => {
  const profile = await BackendService.getProfile(ctx.user.accessToken);
  return {
    ...profile,
    accessToken: ctx.user.accessToken,
  };
};

const isCredentialProvider = (provider: string) => {
  return (
    provider === credentialSigninConfig.id ||
    provider === credentialSignupConfig.id
  );
};

const CredentialAuth = {
  signinConfig: credentialSigninConfig,
  signupConfig: credentialSignupConfig,
  handle: authHandler,
  getProfile: getUserProfile,
  canProvide: isCredentialProvider,
};

export default CredentialAuth;
