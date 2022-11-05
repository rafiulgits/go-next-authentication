import { useState } from "react";

interface Props {
  onSubmit: (data: any) => void;
}

export const SignupForm = (props: Props) => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = () => {
    const data = {
      name,
      email,
      phone,
      password,
    };

    props.onSubmit(data);
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        handleSubmit();
      }}
    >
      <div className="mb-2">
        <label className="form-label">Name</label>
        <input
          className="form-control form-control-sm"
          type="text"
          placeholder="Mr. Example"
          name="name"
          required
          onChange={(e) => setName(e.target.value)}
        />
      </div>

      <div className="mb-2">
        <label className="form-label">Email</label>
        <input
          className="form-control form-control-sm"
          type="email"
          placeholder="example@mail.com"
          name="email"
          required
          onChange={(e) => setEmail(e.target.value)}
        />
      </div>

      <div className="mb-2">
        <label className="form-label">Phone</label>
        <input
          className="form-control form-control-sm"
          type="text"
          placeholder="+8801XXXXXXXXX"
          name="phone"
          required
          onChange={(e) => setPhone(e.target.value)}
        />
      </div>
      <div className="mb-2">
        <label className="form-label">Password</label>
        <input
          className="form-control form-control-sm"
          type="password"
          placeholder="Password"
          name="password"
          required
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>

      <div className="d-grid gap-2">
        <button
          className="btn btn-info"
          type="submit"
          onSubmit={(e) => e.preventDefault()}
        >
          Signup
        </button>
      </div>
    </form>
  );
};
