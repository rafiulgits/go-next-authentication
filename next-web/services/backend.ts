import axios from "axios";
import { CredentialLoginDto, CredentialSignupDto, OAuthDto } from "dtos/auth";

const throwAxiosError = (err: any) => {
  if (!!err.response) {
    if (!!err.response.data && !!err.response.data["message"]) {
      throw new Error(err.response.data["message"]);
    }
    throw new Error(err.message);
  }
  throw new Error("failed to perform request");
};

const BackendService = {
  loginOAuth: async (data: OAuthDto) => {
    try {
      const res = await axios.post(
        `${process.env.BACKEND_API_HOST}/login/oauth`,
        JSON.stringify(data)
      );
      return res.data;
    } catch (err: any) {
      throwAxiosError(err);
    }
  },

  loginCredential: async (data: CredentialLoginDto) => {
    try {
      const res = await axios.post(
        `${process.env.BACKEND_API_HOST}/login/credential`,
        JSON.stringify(data)
      );
      return res.data;
    } catch (err: any) {
      throwAxiosError(err);
    }
  },

  signupOAuth: async (data: OAuthDto) => {
    try {
      const res = await axios.post(
        `${process.env.BACKEND_API_HOST}/signup/oauth`,
        JSON.stringify(data)
      );
      return res.data;
    } catch (err: any) {
      throwAxiosError(err);
    }
  },

  signupCredential: async (data: CredentialSignupDto) => {
    try {
      const res = await axios.post(
        `${process.env.BACKEND_API_HOST}/signup/credential`,
        JSON.stringify(data)
      );
      return res.data;
    } catch (err: any) {
      throwAxiosError(err);
    }
  },

  getProfile: async (accessToken: string) => {
    try {
      const res = await axios.get(`${process.env.BACKEND_API_HOST}/profile`, {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });
      return res.data;
    } catch (err: any) {
      throwAxiosError(err);
    }
  },
};

export default BackendService;
