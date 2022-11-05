import axios from "axios";
import { Layout } from "components/Layout";
import { GetServerSideProps, NextPage } from "next";

const EmailVerifyPage: NextPage = (props: any) => {
  const { valid, user } = props;

  if (!valid) {
    return (
      <Layout title="Verification Failed">
        <div className="container-fluid mt-5">
          <div className="row d-flex justify-content-center">
            <div className="col-sm-6">
              <div className="mt-5">
                <div className="alert alert-danger" role="alert">
                  Verification failed
                </div>
              </div>
            </div>
          </div>
        </div>
      </Layout>
    );
  }

  return (
    <Layout title="Email Verified">
      <div className="container-fluid mt-5">
        <div className="row d-flex justify-content-center">
          <div className="col-sm-6">
            <div className="alert alert-success" role="alert">
              <h4 className="alert-heading">Email Verified!</h4>
              <p>
                Hello {user.name} your email {user.email} is verified.
              </p>
            </div>
          </div>
        </div>
      </div>
    </Layout>
  );
};

export const getServerSideProps: GetServerSideProps = async (context) => {
  const { token } = context.query;
  if (!token) {
    return {
      props: {
        valid: false,
      },
    };
  }

  try {
    let res = await axios.post(
      `${process.env.BACKEND_API_HOST}/email-verify`,
      JSON.stringify({ token: token })
    );
    return {
      props: {
        valid: true,
        user: res.data,
      },
    };
  } catch (err) {
    return {
      props: { valid: false },
    };
  }
};

export default EmailVerifyPage;
