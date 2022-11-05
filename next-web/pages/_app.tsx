import { SessionProvider, useSession } from "next-auth/react";

import "bootstrap/dist/css/bootstrap.css";

export default function App({
  Component,
  pageProps: { session, ...pageProps },
}: any) {
  return (
    <SessionProvider session={session}>
      {Component.auth ? (
        <Auth>
          <Component {...pageProps} />
        </Auth>
      ) : (
        <Component {...pageProps} />
      )}
    </SessionProvider>
  );
}

function Auth(props: any) {
  const { status } = useSession({ required: true });

  if (status === "loading") {
    return (
      <div style={{ marginTop: "45vh" }}>
        <div className="d-flex justify-content-center">
          <div className="spinner-border" role="status">
            <span className="visually-hidden">Loading...</span>
          </div>
        </div>
      </div>
    );
  }

  return props.children;
}
