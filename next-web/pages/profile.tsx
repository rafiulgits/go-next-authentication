import { Layout } from "components/Layout";
import { useSession } from "next-auth/react";

const ProfilePage = (props: any) => {
  const { data } = useSession();

  return (
    <Layout>
      <div className="container-fluid mt-3">
        <div className="row d-flex justify-content-center">
          <div className="col-sm-5">
            <ul className="list-group">
              <li className="list-group-item d-flex justify-content-between align-items-start">
                <div className="ms-2 me-auto">
                  <div className="fw-bold">Name</div>
                  {data?.user?.name}
                </div>
              </li>
              <li className="list-group-item d-flex justify-content-between align-items-start">
                <div className="ms-2 me-auto">
                  <div className="fw-bold">Email</div>
                  {data?.user?.email}
                </div>

                {/* @ts-ignore */}
                {data?.user?.isEmailVerified ? (
                  <span className="badge bg-success rounded-pill">
                    Verified
                  </span>
                ) : (
                  <span className="badge bg-danger rounded-pill">
                    Not Verified
                  </span>
                )}
              </li>
              <li className="list-group-item d-flex justify-content-between align-items-start">
                <div className="ms-2 me-auto">
                  <div className="fw-bold">Phone</div>
                  {/* @ts-ignore */}
                  {data?.user?.phone ? data?.user?.phone : "N/A"}
                </div>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </Layout>
  );
};

ProfilePage.auth = true;

export default ProfilePage;
