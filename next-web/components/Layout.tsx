import Head from "next/head";
import { signOut, useSession } from "next-auth/react";

interface Props {
  title?: string;
  children?: React.ReactNode;
}

export const Layout = (props: Props) => {
  const { status } = useSession();

  return (
    <>
      <Head>
        <title>{props.title}</title>
      </Head>
      <nav className="navbar navbar-expand-lg navbar-light bg-light">
        <div className="container-fluid">
          <a className="navbar-brand" href="/">
            Go Next
          </a>

          <div className="d-flex">
            {status === "authenticated" ? <AuthNav /> : <PublicNav />}
          </div>
        </div>
      </nav>
      {props.children}
    </>
  );
};

const AuthNav = () => {
  const { data } = useSession();
  return (
    <div className="navbar-nav">
      <a className="nav-link" aria-current="page" href="/profile">
        {data?.user?.name}
      </a>
      <a
        className="nav-link"
        aria-current="page"
        href="#"
        onClick={() => {
          signOut();
        }}
      >
        Logout
      </a>
    </div>
  );
};

const PublicNav = () => {
  return (
    <div className="navbar-nav">
      <a className="nav-link" aria-current="page" href="/auth/login">
        Login
      </a>
      <a className="nav-link" aria-current="page" href="/auth/signup">
        Signup
      </a>
    </div>
  );
};
