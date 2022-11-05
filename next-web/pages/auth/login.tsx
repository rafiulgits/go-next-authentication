import { NextPage } from "next";
import { signIn } from "next-auth/react";
import Head from "next/head";
import Link from "next/link";
import { useRouter } from "next/router";
import { LoginForm } from "components/LoginForm";

const Login: NextPage = (props: any) => {
  const router = useRouter();
  const { error, callbackUrl } = router.query;

  const getLoginSuccessRedirectUrl = () => {
    if (callbackUrl && typeof callbackUrl === "string") {
      return callbackUrl;
    }
    return "/profile";
  };

  const handleGoogleLogin = async () => {
    await signIn("google-signin", {
      callbackUrl: getLoginSuccessRedirectUrl(),
      redirect: false,
    });
  };

  const handleCredentialLogin = async (data: any) => {
    const res = await signIn("credentials-signin", {
      ...data,
      callbackUrl: getLoginSuccessRedirectUrl(),
      redirect: false,
    });

    if (res) {
      if (res.ok) {
        router.push(res.url!);
      } else {
        router.push(`/auth/login?error=${res.error}`, undefined, {
          shallow: true,
        });
      }
    }
  };

  return (
    <div>
      <Head>
        <title>Login</title>
      </Head>
      <div>
        <div className="container-fluid mt-4">
          <div className="row d-flex justify-content-center">
            <div className="col-md-4">
              <div className="card p-5">
                <h6 className="text-center">Login</h6>
                {error && (
                  <div className="alert alert-danger" role="alert">
                    {error}
                  </div>
                )}
                <LoginForm onSubmit={handleCredentialLogin} />
                <p className="text-center mt-2">Or</p>
                <button
                  className="btn btn-outline-primary"
                  onClick={handleGoogleLogin}
                >
                  <img
                    src="https://cdn1.iconfinder.com/data/icons/google-s-logo/150/Google_Icons-09-32.png"
                    width="20px"
                    alt="google-icon"
                  />{" "}
                  Login with google
                </button>

                <hr />
                <p>
                  No account? Please <Link href="/auth/signup">Signup</Link>{" "}
                  here
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Login;
