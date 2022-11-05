import { NextPage } from "next";
import { signIn } from "next-auth/react";
import Head from "next/head";
import Link from "next/link";
import { useRouter } from "next/router";
import { SignupForm } from "components/SignupForm";

const SignupPage: NextPage = (props: any) => {
  const router = useRouter();
  const { error } = router.query;

  const handleGoogleSignup = async () => {
    await signIn("google-signup", {
      callbackUrl: "/profile",
      redirect: false,
    });
  };

  const handleCredentialSignup = async (data: any) => {
    const res = await signIn("credentials-signup", {
      ...data,
      callbackUrl: "/profile",
      redirect: false,
    });

    if (res) {
      if (res.ok) {
        router.push(res.url!);
      } else {
        router.push(`/auth/signup?error=${res.error}`, undefined, {
          shallow: true,
        });
      }
    }
  };

  return (
    <div>
      <Head>
        <title>Signup</title>
      </Head>
      <div>
        <div className="container-fluid mt-3">
          <div className="row d-flex justify-content-center">
            <div className="col-md-4">
              <div className="card p-4">
                <h6 className="text-center">Registration</h6>
                {error && (
                  <div className="alert alert-danger" role="alert">
                    {error}
                  </div>
                )}
                <SignupForm onSubmit={handleCredentialSignup} />
                <p className="text-center mt-2">Or</p>
                <button
                  className="btn btn-outline-primary"
                  onClick={handleGoogleSignup}
                >
                  <img
                    src="https://cdn1.iconfinder.com/data/icons/google-s-logo/150/Google_Icons-09-32.png"
                    width="20px"
                    alt="google-icon"
                  />{" "}
                  Signup with google
                </button>
                <hr />
                <p>
                  Already have account? <Link href="/auth/login">Login</Link>{" "}
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

export default SignupPage;
